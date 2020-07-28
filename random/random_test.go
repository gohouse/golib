package random

import "testing"

func TestRandBetween(t *testing.T) {
	for i:=0;i<20;i++{
		//t.Log(RandBetween(0, 11))
		//t.Log(RandomBetween(6, 12, TypeNUMBERIC|TypeCAPITAL|TypeLOWERCASE))
		t.Log(RandomVariable(8))
	}
}
