package calculation

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
		err        error
	}{
		{"2+2*2", "6", nil},
		{"10/2-1", "4", nil},
		{"5-1+3*2", "10", nil},
		{"invalid", "", ErrInvalidExpression},
		{"10/0", "", errors.New("division by zero")},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			result, err := Calc(tt.expression)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			} else if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
