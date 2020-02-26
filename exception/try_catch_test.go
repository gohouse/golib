package exception

import (
	"testing"
)

func TestTry(t *testing.T) {
	TryCatch(func() {
		t.Log("do something buggy")
		panic("asdfasfadsf")
	},
	//func(err e.Error) {
	//	t.Log(err.ErrorWithStack())
	//},
	)
	t.Log("done")
}
