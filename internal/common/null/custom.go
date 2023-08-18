package null

type _nullCustom[T any] struct {
	Void  bool
	Value T
}

type Custom[T any] _nullCustom[T]

func NewCustom[T any](value *T) Custom[T] {
	return Custom[T]{
		Void:  value == nil,
		Value: *value,
	}
}
