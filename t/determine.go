package t

import "encoding/json"

type iDetermine interface {
	IsNumeric() bool
	IsInteger() bool
	IsFloat() bool
	IsJsonSlice() bool
	IsJsonMap() bool
	IsJson() bool
}

// IsInteger 是否为整数
func (t Type) IsInteger() bool {
	return t.String() == New(t.Int64()).String()
}

// IsNumeric 是否为数字,包含整数和小数
func (t Type) IsNumeric() bool {
	return t.String() == New(t.Float64()).String()
}

// IsFloat 是否为float
func (t Type) IsFloat() bool {
	switch t.val.(type) {
	case float64, float32:
		return true
	default:
		if t.IsNumeric() && !t.IsInteger() {
			return true
		}
	}
	return false
}

// IsJsonSlice ...
func (t Type) IsJsonSlice() bool {
	var js []interface{}
	return json.Unmarshal(t.Bytes(), &js) == nil
}

// IsJsonMap ...
func (t Type) IsJsonMap() bool {
	var js map[string]interface{}
	return json.Unmarshal(t.Bytes(), &js) == nil
}

// IsJson ...
func (t Type) IsJson() bool {
	return t.IsJsonSlice() || t.IsJsonMap()
}
