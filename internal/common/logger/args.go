package logger

type Args map[string]any

func (model Args) ParseArgs() []any {
	if model == nil {
		return nil
	}

	var keyCount int
	for range model {
		keyCount++
	}

	argsSlice := make([]any, 0, keyCount*2)
	for key, value := range model {
		argsSlice = append(argsSlice, key, value)
	}

	return argsSlice
}
