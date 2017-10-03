package con

import (
	"github.com/GanymedeNil/shorturl/config"
	"github.com/garyburd/redigo/redis"
	"log"
)

var redisConn *redis.Conn

func Redis() redis.Conn {

	if redisConn != nil {
		return *redisConn
	}
	address := config.Get("redis.default")
	c, err := redis.Dial("tcp", address.(string))
	if err != nil {
		log.Panic("data conn error")
	}
	redisConn = &c
	return *redisConn

}
