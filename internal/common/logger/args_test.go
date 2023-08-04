package logger_test

import (
	"monolith/internal/common/logger"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseArgs(t *testing.T) {
	type args struct {
		arguments logger.Args
	}
	testCases := map[string]struct {
		args args
		want []any
	}{
		"one key-value": {
			args: args{
				arguments: map[string]any{"example": "example"},
			},
			want: []any{"example", "example"},
		},
		"many key-value": {
			args: args{
				arguments: map[string]any{
					"example1": "example1",
					"example2": "example2",
					"example3": "example3",
				},
			},
			want: []any{"example1", "example1", "example2", "example2", "example3", "example3"},
		},
		"nil key-value": {
			args: args{arguments: nil},
			want: nil,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.args.arguments.ParseArgs()
			assert.Equal(t, tc.want, got)
		})
	}
}
