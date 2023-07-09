package wamp

import (
	"errors"
	"fmt"
	"time"
	"web_server/pkg/authentication"

	"github.com/gammazero/nexus/v3/router/auth"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type JWTAuthenticator struct {
	keyStore auth.KeyStore
	timeout  time.Duration
	auth     *authentication.Authentication
}

func InitAuthenticator(keystore auth.KeyStore, timeout time.Duration, auth *authentication.Authentication) *JWTAuthenticator {
	return &JWTAuthenticator{keyStore: keystore, timeout: timeout, auth: auth}
}

func (t *JWTAuthenticator) Authenticate(sid wamp.ID, details wamp.Dict, client wamp.Peer) (*wamp.Welcome, error) {
	// The HELLO.Details.authid|string is the authentication ID (e.g. username)
	// the client wishes to authenticate as. For Ticket-based authentication,
	// this MUST be provided.
	authID, _ := wamp.AsString(details["authid"])
	if authID == "" {
		return nil, errors.New("missing authid")
	}

	authrole, err := t.keyStore.AuthRole(authID)
	if err != nil {
		authrole = ""
	}

	ks, ok := t.keyStore.(auth.BypassKeyStore)
	if ok {
		if ks.AlreadyAuth(authID, details) {
			// Create welcome details containing auth info.
			welcome := &wamp.Welcome{
				Details: wamp.Dict{
					"authid":       authID,
					"authrole":     authrole,
					"authmethod":   t.AuthMethod(),
					"authprovider": t.keyStore.Provider(),
				},
			}
			if err = ks.OnWelcome(authID, welcome, details); err != nil {
				return nil, err
			}
			return welcome, nil
		}
	}

	// Challenge Extra map is empty since the ticket challenge only asks for a
	// ticket (using authmethod) and provides no additional challenge info.
	err = client.Send(&wamp.Challenge{
		AuthMethod: t.AuthMethod(),
		Extra:      wamp.Dict{},
	})
	if err != nil {
		return nil, err
	}

	// Read AUTHENTICATE response from client.
	msg, err := wamp.RecvTimeout(client, t.timeout)
	if err != nil {
		return nil, err
	}
	authRsp, ok := msg.(*wamp.Authenticate)
	if !ok {
		return nil, fmt.Errorf("unexpected %v message received from client %v",
			msg.MessageType(), client)
	}

	// The client will send an AUTHENTICATE message containing a ticket.  The
	// server will then check if the ticket provided is permissible (for the
	// authid given).
	err = t.verifyTicket(authRsp.Signature)
	if err != nil {
		return nil, err
	}

	// Create welcome details containing auth info.
	welcome := &wamp.Welcome{
		Details: wamp.Dict{
			"authid":       authID,
			"authmethod":   t.AuthMethod(),
			"authrole":     authrole,
			"authprovider": t.keyStore.Provider(),
		},
	}

	if ks != nil {
		// Tell the keystore that the client was authenticated, and provide the
		// transport details if available.
		if err = ks.OnWelcome(authID, welcome, details); err != nil {
			return nil, err
		}
	}
	return welcome, nil
}

func (t *JWTAuthenticator) AuthMethod() string {
	return "jwt_ticket"
}

func (t *JWTAuthenticator) verifyTicket(tokenStr string) error {
	authToken, err := jwt.ParseWithClaims(tokenStr, &authentication.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.auth.Secret, nil
	})

	if authToken != nil && authToken.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		log.Error("Auth token is not valid")
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			log.Error("Auth token is expired")
			return errors.New("Token expired")
		}
		log.WithError(err)
		return errors.New("Error in auth token")
	}

	log.WithError(err)
	return errors.New("Error in auth token")
}
