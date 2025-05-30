package stats

import (
	"math"
	"testing"
)

const (
	ErrTolerance = 1e-6
)

func TestNormGaussAt(t *testing.T) {
	t.Parallel()

	in := []float64{
		0,
		1,
		-1,
		3.7,
		-3.7,
	}

	// These values were generated using dnorm on http://www.r-fiddle.org/
	out := []float64{
		0.3989422804,
		0.24197072452,
		0.24197072452,
		0.00042478027055,
		0.00042478027055,
	}

	for i, v := range in {
		result := NormGaussAt(v)
		if math.Abs(result-out[i]) > ErrTolerance {
			t.Error("Expected", out[i], "got", result, "for", v)
		}
	}
}

func TestGaussAt(t *testing.T) {
	t.Parallel()

	in := []float64{
		20,
		25,
		15,
		38.5,
		1.5,
	}

	// These values were generated using dnorm on http://www.r-fiddle.org/
	out := []float64{
		0.07978845608,
		0.048394144904,
		0.048394144904,
		8.495605411e-5,
		8.495605411e-5,
	}

	for i, v := range in {
		result := GaussAt(v, 20, 5)
		if math.Abs(result-out[i]) > ErrTolerance {
			t.Error("Expected", out[i], "got", result, "for", v)
		}
	}
}

func TestNormalCDF(t *testing.T) {
	t.Parallel()

	in := []float64{
		math.Inf(1),
		0,
		math.Inf(-1),
		1,
		-1,
		2,
		-2,
		InverseCDF(0.86),
	}

	// These results were calculated using pnorm(x) on www.r-fiddle.org
	out := []float64{
		1,
		0.5,
		0,
		0.841344746069,
		0.158655253931,
		0.977249868052,
		0.022750131948,
		0.86,
	}

	if len(in) != len(out) {
		t.Error("Input values have a different length from their outputs")
	}

	for i := range in {
		o := NormalCDF(in[i])

		if math.Abs(o-out[i]) > ErrTolerance {
			t.Errorf("NormalCDF(%f) == %f, want %f", in[i], o, out[i])
		}
	}
}

func TestInverseCDF(t *testing.T) {
	t.Parallel()

	in := []float64{
		1,
		0.5,
		0,
		0.2,
		0.3,
		0.9,
		-1,
		2,
		NormalCDF(4),
	}

	// These results were calculated using qnorm(x) on www.r-fiddle.org
	out := []float64{
		math.Inf(1),
		0,
		math.Inf(-1),
		-0.84162123357,
		-0.52440051271,
		1.2815515655,
		math.NaN(),
		math.NaN(),
		4,
	}

	if len(in) != len(out) {
		t.Error("Input values have a different length from their outputs")
	}

	for i := range in {
		o := InverseCDF(in[i])

		if math.Abs(o-out[i]) > ErrTolerance {
			t.Errorf("InverseCDF(%f) == %f, want %f", in[i], o, out[i])
		}
	}
}
