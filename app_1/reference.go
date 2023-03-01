// package app

// import (
// 	"sync"
// 	"web_server/pkg/models"
// 	"web_server/pkg/redis"

// 	log "github.com/sirupsen/logrus"
// )

// type Reference struct {
// 	Config *models.Reference
// 	mtx    sync.RWMutex
// 	cli    *redis.Client
// }

// var ref *Reference

// func InitReference(redisCli *redis.Client) (*Reference, error) {
// 	ref = &Reference{mtx: sync.RWMutex{}, cli: redisCli}
// 	log.Info("Collecting reference configuration")
// 	ref.readFromRedis()
// 	return ref, nil
// }
// func (r *Reference) readFromRedis() (*models.Reference, error) {

// 	return nil, nil
// }
