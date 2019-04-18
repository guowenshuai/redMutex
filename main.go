package redMutex

import (
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

// RedMutex use redis distributed lock
func RedMutex(name string) (*redsync.Mutex, error) {
	// get redis pools
	pools := NewPools(ServerConfig{"tcp", "10.10.11.141:6379", []redis.DialOption{redis.DialPassword("password")}},
		ServerConfig{"tcp", "10.10.11.137:6379", []redis.DialOption{redis.DialPassword("password")}})

	logrus.Info("get redis pool success")

	mutex := redsync.New(pools).NewMutex(name, redsync.SetExpiry(time.Duration(1)*time.Second),
		redsync.SetRetryDelay(time.Duration(100)*time.Millisecond),
		redsync.SetTries(10),
	)
	return mutex, nil
}
