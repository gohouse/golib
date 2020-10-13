package dc

import (
	"testing"
)

func TestDecimalToAny(t *testing.T) {
	t.Log(DecimalTo62(22))
	t.Log(ToDecimal("X"))
}
