package wilks

import (
	"testing"
)

func TestPoundsToKilos(t *testing.T) {
	tests := []struct {
		in  float64
		out float64
	}{
		{0, 0},
		{poundsPerKilo, 1},
		{100, 45.35929094356398},
	}

	for idx, tt := range tests {
		if PoundsToKilos(tt.in) != tt.out {
			t.Fatalf("[%v] Unexpected kilos, expected=%v, got=%v", idx, tt.out, PoundsToKilos(tt.in))
		}
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		female bool
		bw     float64
		total  float64
		out    float64
	}{
		{true, 52, 200, 249.3273978860303},
		{true, 84, 550, 490.45711693782926},
		{false, 83, 500, 333.7497133016794},
		{false, 74, 600, 431.58815523816156},
	}

	for idx, tt := range tests {
		if Score(tt.female, tt.bw, tt.total) != tt.out {
			t.Fatalf("[%v] Unexpected Score, expected=%v, got=%v", idx, tt.out, Score(tt.female, tt.bw, tt.total))
		}
	}
}

func TestCoefficient(t *testing.T) {
	tests := []struct {
		female bool
		bw     float64
		out    float64
	}{
		{true, 52, 1.2466369894301514},
		{true, 80, 0.9150087716058078},
		{false, 83, 0.6674994266033587},
		{false, 74, 0.7193135920636026},
	}

	for idx, tt := range tests {
		if coefficient(tt.female, tt.bw) != tt.out {
			t.Fatalf("[%v] Unexpected coefficient, expected=%v, got=%v", idx, tt.out, coefficient(tt.female, tt.bw))
		}
	}
}
