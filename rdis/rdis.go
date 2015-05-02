package rdis

import (
	"time"

	"github.com/drinkin/di/env"
	"github.com/drinkin/shop/src/lg"
	"github.com/garyburd/redigo/redis"
)

//PoolFromEnv with `REDIS_URL` and optional `REDIS_PASS`
func PoolFromEnv() *redis.Pool {
	return NewPool(env.MustGet("REDIS_URL"), env.Get("REDIS_PASS"))
}

//NewPool creates a new redigo pool
func NewPool(server, password string) *redis.Pool {
	lg.Debug("Redis Connect", lg.F{"url": server})

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close() // Don't care if fails
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
