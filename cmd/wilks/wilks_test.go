package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		bw     float64
		total  float64
		gender string
		out    bool
	}{
		{0, 10, "m", false},
		{10, 0, "m", false},
		{0, 0, "m", false},
		{10, 10, "", false},
		{10, 10, "test", false},
		{10, 10, "m", true},
		{10, 10, "f", true},
	}

	for idx, tt := range tests {
		if validate(tt.bw, tt.total, tt.gender) != tt.out {
			t.Fatalf("[%v] Unexpected validate, expected=%v, got=%v", idx, tt.out, validate(tt.bw, tt.total, tt.gender))
		}
	}
}
