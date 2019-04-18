package redMutex

import (
	"testing"
	"time"
)

func TestRedMutex(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i := 0
	mutex, err := RedMutex("test")
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		// t.Run(strconv.Itoa(tt), func(t *testing.T) {
		// 	if err := mutex.Lock(); err != nil {
		// 		t.Error("get distributed lock err: ", err.Error())
		// 		mutex.Extend()
		// 		return
		// 	}
		// 	i++
		// 	t.Log("i: ", i)
		// 	mutex.Unlock()
		// })
		if err := mutex.Lock(); err != nil {
			t.Error("get distributed lock err: ", err.Error())
			continue
		}
		i++
		time.Sleep(1)
		t.Log("tt: ", tt)
		t.Log("i: ", i)
		t.Log("unlock: ", mutex.Unlock())
	}
}
