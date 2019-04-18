package redMutex

import (
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	error2 "redMutexError"
)

// ServerConfig is custom options, define
// redis options.
type ServerConfig struct {
	Network string
	Address string
	Options []redis.DialOption
}

func newPool(server ServerConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(server.Network, server.Address, server.Options...)
			if err != nil {
				logrus.Error(error2.ErrGetPool.AppendErrors(err))
				return nil, error2.ErrGetPool.AppendErrors(err)
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// NewPools return redis pools
func NewPools(servers ...ServerConfig) []redsync.Pool {
	pools := make([]redsync.Pool, 0)
	for _, server := range servers {
		pool := newPool(server)
		pools = append(pools, pool)
	}
	return pools
}
