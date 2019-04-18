module github.com/guowenshuai/redMutex

go 1.12

require (
	github.com/go-redsync/redsync v1.2.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/sirupsen/logrus v1.4.1
	redMutexError v0.0.0
)

replace redMutexError => ./error
