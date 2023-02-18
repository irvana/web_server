package wamp

import (
	"fmt"

	"github.com/go-redis/redis"
)

type sKeyStore struct {
	ticket       string
	redisStorage *redis.Client
}

func NewKeyStore(storage *redis.Client) *sKeyStore {
	return &sKeyStore{redisStorage: storage, ticket: "ticket"}
}

func (s *sKeyStore) AuthKey(authid string, authmethod string) ([]byte, error) {
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
