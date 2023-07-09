package wamp

import (
	"fmt"

	"github.com/gammazero/nexus/v3/router/auth"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type sKeyStore struct {
	ticket       string
	redisStorage *redis.Client
}

func NewKeyStore(storage *redis.Client) auth.KeyStore {
	return &sKeyStore{redisStorage: storage, ticket: "jwt_ticket"}
}

func (s *sKeyStore) AuthKey(authid string, authmethod string) ([]byte, error) {
	logrus.WithField("authid", authid).WithField("authmethod ", authmethod).Info("authkey")
	switch authmethod {
	case "ticket":
		val, err := s.redisStorage.Get(fmt.Sprintf("REFRESH-%s", authid)).Result()
		if err != nil {
			return nil, err
		}
		return []byte(val), nil
	}
	return nil, nil
}

func (s *sKeyStore) PasswordInfo(authid string) (salt string, keylen int, iterations int) {
	return "salt", 32, 1000
}

func (s *sKeyStore) AuthRole(authid string) (string, error) {
	return "user", nil
}

func (s *sKeyStore) Provider() string {
	return s.ticket
}
