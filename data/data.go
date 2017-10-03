package data

import (
	"github.com/GanymedeNil/shorturl/base62"
	"github.com/GanymedeNil/shorturl/con"
	"github.com/garyburd/redigo/redis"
	"log"
)

func SetURL(URL string) string {
	c := con.Redis()
	tinyURL := base62.Encode(countAdd())
	_, err := redis.Int64(c.Do("HSET", "URLLIST", tinyURL, URL))
	if err != nil {
		log.Fatal(err)
	}
	return tinyURL

}

func GetURL(tinyURL string) string {
	c := con.Redis()
	data, err := redis.String(c.Do("HGET", "URLLIST", tinyURL))
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func countAdd() int {
	c := con.Redis()
	count, err := redis.Int(c.Do("INCR", "URLCOUNT"))
	if err != nil {
		log.Panic(err)
	}
	return count
}
