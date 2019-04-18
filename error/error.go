package error

import "errors"

var redMutexPrefix = "redMutex: "

type RedMutexError struct {
	prefix string
	error
	details []error
}

func newError(msg string) RedMutexError {
	return RedMutexError{
		prefix: redMutexPrefix,
		error:  errors.New(redMutexPrefix + msg),
	}
}

func (r RedMutexError) AppendErrors(errs ...error) RedMutexError {
	r.details = append(r.details, errs...)
	return r
}

func (r RedMutexError) Error() string {
	msg := r.error.Error()
	for _, e := range r.details {
		msg += "; " + e.Error()
	}
	return msg
}

var ErrGetLock = newError("can not get redisLock")
var ErrInitRedisDia = newError("init redis dial error")
var ErrGetPool = newError("get redis pool error")
var ErrGetPools = newError("get redis pools error")
