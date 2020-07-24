package exception

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/t"
)

var (
	defaultCatcher = func(err e.Error) {
		fmt.Println("panic err catch:", err.ErrorWithStack())
	}
	defaultErrorStack = 5
)

func SetDefaultCatcher(dr func(err e.Error)) {
	defaultCatcher = dr
}

func SetDefaultErrorStack(errorStack int) {
	defaultErrorStack = errorStack
}

func TryCatch(try func(), catch ...func(err e.Error)) {
	defer func() {
		if err := recover(); err != nil {
			var err2 = e.New(t.New(err).String(), defaultErrorStack)
			if len(catch) > 0 {
				catch[0](err2)
			} else {
				defaultCatcher(err2)
			}
		}
	}()
	try()
}
