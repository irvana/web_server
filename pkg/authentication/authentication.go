package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-redis/redis"

	jwt "github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type (
	// JWTClaims Wrapper for JWT claims
	JWTClaims struct {
		Csrf   string `json:"csrf"`
		UserID string `json:"user_id"`
		jwt.RegisteredClaims
	}

	Authentication struct {
		SignKey   *rsa.PrivateKey
		VerifyKey *rsa.PublicKey
		cfg       Config
		redisCli  *redis.Client
	}

	Config struct {
		SignMethod   string        `mapstructure:"signMethod"`
		Param        string        `mapstructure:"param"`
		Secret       string        `mapstructure:"secret"`
		PrivKeyPath  string        `mapstructure:"privPath"`
		PubKeyPath   string        `mapstructure:"pubPath"`
		AuthExp      time.Duration `mapstructure:"authExp"`
		RefreshExp   time.Duration `mapstructure:"refExp"`
		RefPrefix    string        `mapstructure:"refPref"`
		AuthPrefix   string        `mapstructure:"authPref"`
		FxIssuer     string        `mapstructure:"fxIssuer"`
		IsPersistent bool          `mapstructure:"isPersistent"`
	}

	AuthDetails struct {
		CSRF         string
		AuthToken    string
		RefreshToken string
	}
)

var Jwt *Authentication

func InitAuthModule(cfg Config, redisCli *redis.Client) (*Authentication, error) {
	priv, err := ioutil.ReadFile(cfg.PrivKeyPath)
	if err != nil {
		return nil, err
	}

	SignKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, err
	}

	pub, err := ioutil.ReadFile(cfg.PubKeyPath)
	if err != nil {
		return nil, err
	}

	VerifyKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, err
	}

	Jwt = &Authentication{SignKey: SignKey, VerifyKey: VerifyKey, cfg: cfg, redisCli: redisCli}
	return Jwt, nil
}

// generateAuthToken generate JWT for respective configuration
func (j *Authentication) generateAuthToken(uuid, csrfSecret string) (authTokenStr string, err error) {
	claims := JWTClaims{
		Csrf:   csrfSecret,
		UserID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.FxIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AuthExp)),
		}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	authTokenStr, err = token.SignedString(j.SignKey)

	ttl := j.cfg.AuthExp
	if j.cfg.IsPersistent {
		ttl = 0
	}
	j.redisCli.Set(fmt.Sprintf(j.cfg.AuthPrefix, uuid), authTokenStr, ttl)
	return
}

// generateRefreshToken generate JWT for respective configuration
func (j *Authentication) generateRefreshToken(uuid, csrfSecret string) (refreshTokenStr string, err error) {
	refTokenKey := fmt.Sprintf(j.cfg.RefPrefix, uuid)
	refTokenJTI := j.redisCli.Get(refTokenKey).Val()

	if refTokenJTI == "" {
		refTokenJTI, _ = j.generateRandomString()
		j.redisCli.Set(refTokenKey, refTokenJTI, 0)
	}

	claims := JWTClaims{
		Csrf:   csrfSecret,
		UserID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        refTokenJTI,
			Issuer:    j.cfg.FxIssuer,
			Subject:   uuid,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.RefreshExp)),
		}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	refreshTokenStr, err = token.SignedString(j.SignKey)
	return
}

// generateRandomString will return encoded string
func (j *Authentication) generateRandomString() (csrfStr string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return
	}
	csrfStr = base64.URLEncoding.EncodeToString(b)
	return
}

func (j *Authentication) updateAuthTokenString(refreshTokenString string, oldAuthTokenString string) (newAuthTokenString, csrfSecret string, err error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})
	refreshTokenClaims, ok := refreshToken.Claims.(*JWTClaims)
	if !ok {
		err = fmt.Errorf("%w Error reading jwt claims", err)
		return
	}

	exist := j.redisCli.Get(fmt.Sprintf(j.cfg.RefPrefix, refreshTokenClaims.UserID)).Val()

	// check if the refresh token has been revoked
	if exist != "" {
		// the refresh token has not been revoked
		// has it expired?
		if refreshToken.Valid {
			// nope, the refresh token has not expired
			// issue a new auth token
			authToken, _ := jwt.ParseWithClaims(oldAuthTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return j.VerifyKey, nil
			})

			oldAuthTokenClaims, ok := authToken.Claims.(*JWTClaims)
			if !ok {
				err = fmt.Errorf("%w Error reading jwt claims", err)
				return
			}

			// our policy is to regenerate the csrf secret for each new auth token
			csrfSecret, err = j.generateRandomString()
			if err != nil {
				return
			}

			newAuthTokenString, err = j.generateAuthToken(oldAuthTokenClaims.UserID, csrfSecret)
			return
		}
		err = errors.New("Unauthorized")
		// the refresh token has expired!
		log.WithError(err).Error("Refresh token has expired!")
		// Revoke the token in our db and require the user to login again
		j.redisCli.Del(fmt.Sprintf(j.cfg.RefPrefix, refreshTokenClaims.UserID))

		return
	}
	// the refresh token has been revoked!
	err = errors.New("Unauthorized")
	log.WithError(err).Error("Refresh token has been revoked!")
	return
}

func (j *Authentication) updateRefreshTokenExp(oldRefreshTokenString string) (newRefreshTokenString string, err error) {
	refreshToken, err := jwt.ParseWithClaims(oldRefreshTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})

	oldRefreshTokenClaims, ok := refreshToken.Claims.(*JWTClaims)
	if !ok {
		return
	}

	refreshClaims := JWTClaims{
		Csrf:   oldRefreshTokenClaims.Csrf,
		UserID: oldRefreshTokenClaims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        oldRefreshTokenClaims.RegisteredClaims.ID, // jti
			Subject:   oldRefreshTokenClaims.RegisteredClaims.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 7)),
		},
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refreshClaims)

	// generate the refresh token string
	newRefreshTokenString, err = refreshJwt.SignedString(j.SignKey)
	return
}

func (j *Authentication) updateRefreshTokenCsrf(oldRefreshTokenString string, newCsrfString string) (newRefreshTokenString string, err error) {
	refreshToken, err := jwt.ParseWithClaims(oldRefreshTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})

	oldRefreshTokenClaims, ok := refreshToken.Claims.(*JWTClaims)
	if !ok {
		return
	}

	refreshClaims := JWTClaims{
		UserID: oldRefreshTokenClaims.UserID,
		Csrf:   newCsrfString,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        oldRefreshTokenClaims.RegisteredClaims.ID, // jti
			Subject:   oldRefreshTokenClaims.RegisteredClaims.Subject,
			ExpiresAt: oldRefreshTokenClaims.RegisteredClaims.ExpiresAt,
		},
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refreshClaims)

	// generate the refresh token string
	newRefreshTokenString, err = refreshJwt.SignedString(j.SignKey)
	return
}

// GenerateNewTokens generate new tokens for user
func (j *Authentication) GenerateNewTokens(uuid string) (*AuthDetails, error) {
	csrf, err := j.generateRandomString()
	if err != nil {
		log.WithError(err).Error("Error Generate CSRF")
		return nil, err
	}
	accessToken, err := j.generateAuthToken(uuid, csrf)
	if err != nil {
		log.WithError(err).Error("Error Generate Access Token")
		return nil, err
	}

	refToken, err := j.generateRefreshToken(uuid, csrf)
	if err != nil {
		log.WithError(err).Error("Error Generate Access Token")
		return nil, err
	}

	return &AuthDetails{
		AuthToken:    accessToken,
		CSRF:         csrf,
		RefreshToken: refToken,
	}, nil

}

// CheckAndRefreshTokens will be called by auth handler
func (j *Authentication) CheckAndRefreshTokens(oldAuthTokenString string, oldRefreshTokenString string, oldCsrfSecret string) (newAuthTokenString, newRefreshTokenString, newCsrfSecret string, err error) {
	if oldCsrfSecret == "" {
		err = errors.New("Unauthorized")
		log.WithError(err).Error("No CSRF token!")
		return
	}
	// now, check that it matches what's in the auth token claims
	authToken, err := jwt.ParseWithClaims(oldAuthTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})
	authTokenClaims, ok := authToken.Claims.(*JWTClaims)
	if !ok {
		return
	}
	if oldCsrfSecret != authTokenClaims.Csrf {
		err = errors.New("Unauthorized")
		log.WithError(err).Error("CSRF token doesn't match jwt!")
		return
	}

	if authToken.Valid {
		log.Info("Auth token is valid")
		newCsrfSecret = authTokenClaims.Csrf
		newRefreshTokenString, err = j.updateRefreshTokenExp(oldRefreshTokenString)
		newAuthTokenString = oldAuthTokenString
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		log.Info("Auth token is not valid")
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			log.Error("Auth token is expired")
			// auth token is expired
			newAuthTokenString, newCsrfSecret, err = j.updateAuthTokenString(oldRefreshTokenString, oldAuthTokenString)
			if err != nil {
				return
			}

			// update the exp of refresh token string
			newRefreshTokenString, err = j.updateRefreshTokenExp(oldRefreshTokenString)
			if err != nil {
				return
			}

			// update the csrf string of the refresh token
			newRefreshTokenString, err = j.updateRefreshTokenCsrf(newRefreshTokenString, newCsrfSecret)
			return
		}
		err = fmt.Errorf("%w Error in auth token", err)
		log.WithError(err)
		return
	}
	err = fmt.Errorf("%w Unauthorized", err)
	log.WithError(err)
	return
}

// RevokeRefreshToken revoking refresh token
func (j *Authentication) RevokeRefreshToken(refreshTokenString string) error {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})
	if err != nil {
		return fmt.Errorf("%w Could not parse refresh token with claims", err)
	}

	refreshTokenClaims, ok := refreshToken.Claims.(*JWTClaims)
	if !ok {
		return fmt.Errorf("%w Could not read refresh token claims", err)
	}

	i := j.redisCli.Del(fmt.Sprintf(j.cfg.RefPrefix, refreshTokenClaims.RegisteredClaims.ID))
	if i.Err() != nil {
		return i.Err()
	}

	return nil
}

// CheckSimpleAuthToken will checking auth token only without affecting the refresh token
func (j *Authentication) CheckSimpleAuthToken(oldAuthTokenString string, oldCsrfSecret string) (newAuthTokenString, newCsrfSecret string, err error) {
	if oldCsrfSecret == "" || oldAuthTokenString == "" {
		err = errors.New("Unauthorized")
		log.WithError(err).Error("No CSRF or auth token!")
		return
	}
	// now, check that it matches what's in the auth token claims
	authToken, err := jwt.ParseWithClaims(oldAuthTokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.VerifyKey, nil
	})
	authTokenClaims, ok := authToken.Claims.(*JWTClaims)
	if !ok {
		return
	}
	if oldCsrfSecret != authTokenClaims.Csrf {
		log.WithError(err).Error("CSRF token doesn't match jwt!")
		err = errors.New("Unauthorized")
		return
	}

	if authToken.Valid {
		log.Info("Auth token is valid")
		newCsrfSecret = authTokenClaims.Csrf

		newAuthTokenString = oldAuthTokenString
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		log.Error("Auth token is not valid")
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			log.Error("Auth token is expired")
			return
		}
	}
	err = fmt.Errorf("%w Error in auth token", err)
	log.WithError(err)
	return
}
