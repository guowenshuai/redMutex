# redMutex
redis distributed lock 


-----
## get pools

`NewPools` get redis pools by `ServerConfig`

```
pools := NewPools(ServerConfig{"tcp", "10.10.11.141:6379", []redis.DialOption{redis.DialPassword("password")}},
		ServerConfig{"tcp", "10.10.11.137:6379", []redis.DialOption{redis.DialPassword("password")}})

```

## get `Mutex`

```
	mutex := redsync.New(pools).NewMutex(name, redsync.SetExpiry(time.Duration(1)*time.Second),
		redsync.SetRetryDelay(time.Duration(100)*time.Millisecond),
		redsync.SetTries(10),
	)
```
