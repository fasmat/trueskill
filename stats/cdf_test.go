package stats

import (
	"math"
	"testing"
)

func TestNormalCDF(t *testing.T) {
	in := []float64{
		math.Inf(1),
		0,
		math.Inf(-1),
		InverseCDF(0.1),
		InverseCDF(0.9),
		InverseCDF(0.86),
	}

	out := []float64{
		1,
		0.5,
		0,
		0.1,
		0.9,
		0.86,
	}

	if len(in) != len(out) {
		t.Error("Input values have a different length from their outputs")
	}

	for i := range in {
		o := NormalCDF(in[i])

		if math.Abs(o-out[i]) > 1e-12 {
			t.Errorf("NormalCDF(%f) == %f, want %f", in[i], o, out[i])
		}
	}
}

func TestInverseCDF(t *testing.T) {
	in := []float64{
		1,
		0.5,
		0,
		NormalCDF(1),
		NormalCDF(-1),
		NormalCDF(4),
	}

	out := []float64{
		math.Inf(1),
		0,
		math.Inf(-1),
		1,
		-1,
		4,
	}

	if len(in) != len(out) {
		t.Error("Input values have a different length from their outputs")
	}

	for i := range in {
		o := InverseCDF(in[i])

		if math.Abs(o-out[i]) > 1e-12 {
			t.Errorf("InverseCDF(%f) == %f, want %f", in[i], o, out[i])
		}
	}
}
