package helper

import (
	"fmt"
	"github.com/gohouse/t"
	"log"
	"sync"
)

// Max 获取最大的数
func Max(args ...interface{}) int {
	if len(args) == 0 {
		return 0
	}
	var max = t.New(args[0]).Int()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := t.New(arg).Int()
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// Min 获取最小的数
func Min(args ...interface{}) int {
	if len(args) == 0 {
		return 0
	}
	var min = t.New(args[0]).Int()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := t.New(arg).Int()
		if tmp < min {
			min = tmp
		}
	}
	return min
}

var defaultRecover = func(err error) {
	log.Println("panic err catch:", err)
}

func SetDefaultRecover(dr func(err error)) {
	defaultRecover = dr
}
func WithRecover(h func(), errFunc ...func(err error)) {
	defer func() {
		if err := recover(); err != nil {
			if len(errFunc) > 0 {
				errFunc[0](fmt.Errorf("%v", err))
			} else {
				defaultRecover(fmt.Errorf("%v", err))
			}
		}
	}()
	h()
}

var mu *sync.Mutex
var mur *sync.RWMutex

func WithLockContext(ctx func()) {
	mu.Lock()
	defer mu.Unlock()
	ctx()
}
func WithRLockContext(ctx func()) {
	mur.RLock()
	defer mur.RUnlock()
	ctx()
}
