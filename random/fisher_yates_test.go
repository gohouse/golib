package random

import (
	"testing"
	"time"
)

func TestFisherYates(t *testing.T) {
	var arr = []interface{}{"a","b","c","d","e","f","g","h"}
	for i:=0;i<10;i++{
		go func() {
			res := FisherYates(arr)
			t.Log(res)
		}()
	}
	time.Sleep(time.Millisecond * 100)
}
