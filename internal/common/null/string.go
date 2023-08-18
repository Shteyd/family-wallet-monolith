package null

type _nullString struct {
	Void  bool
	Value string
}

type String _nullString

func NewString(value *string) String {
	return String{
		Void:  value == nil,
		Value: *value,
	}
}

func (model String) IsZero() bool {
	return model.Value == ""
}
