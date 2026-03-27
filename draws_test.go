package trueskill_test

import (
	"math"
	"testing"

	"github.com/fasmat/trueskill"
)

const (
	ErrTolerance float64 = 1e-06
)

func TestDrawMarginFromDrawProbability(t *testing.T) {
	t.Parallel()

	beta := 25.0 / 6

	pDraw := []float64{
		0.10,
		0.25,
		0.33,
	}
	epsilon := []float64{
		0.74046637542690541,
		1.87760059883033,
		2.5111010132487492,
	}

	for i, p := range pDraw {
		e := trueskill.GetDrawMargin(p, beta, 2)
		if math.Abs(e-epsilon[i]) > ErrTolerance {
			t.Error("Expected draw margin for", p, "=", epsilon[i], "got", e)
		}
	}
}

func TestGetDrawProbabilityFromDrawMargin(t *testing.T) {
	t.Parallel()

	beta := 25.0 / 6

	epsilon := []float64{
		0.74046637542690541,
		1.87760059883033,
		2.5111010132487492,
	}
	pDraw := []float64{
		0.10,
		0.25,
		0.33,
	}

	for i, e := range epsilon {
		p := trueskill.GetDrawProbability(e, beta, 2)
		if math.Abs(p-pDraw[i]) > ErrTolerance {
			t.Error("Expected draw probability for", e, "=", pDraw[i], "got", p)
		}
	}
}
