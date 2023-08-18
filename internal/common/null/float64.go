package null

type _nullFloat64 struct {
	Void  bool
	Value float64
}

type Float64 _nullFloat64

func NewFloat64(value *float64) Float64 {
	return Float64{
		Void:  value == nil,
		Value: *value,
	}
}
