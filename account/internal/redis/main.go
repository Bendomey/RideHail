package redis

import (
	"fmt"

	"github.com/Bendomey/RideHail/account/pkg/utils"
	"github.com/go-redis/redis/v8"
)

var host, port, password string

func init() {
	host = utils.MustGet("REDIS_HOST")
	port = utils.MustGet("REDIS_PORT")
	password = ""
}

// Factory connects app to reid
func Factory() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	return rdb

}
