package null

type _nullInt struct {
	Void  bool
	Value int
}

type Int _nullInt

func NewInt(value *int) Int {
	return Int{
		Void:  value == nil,
		Value: *value,
	}
}
