package null

type _nullBool struct {
	Void  bool
	Value bool
}

type Bool _nullBool

func NewBool(value *bool) Bool {
	return Bool{
		Void:  value == nil,
		Value: *value,
	}
}
