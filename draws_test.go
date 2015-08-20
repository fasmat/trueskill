package trueskill

import (
	"math"
	"testing"
)

const (
	errTolerance float64 = 1e-06
)

func TestDrawMargin(t *testing.T) {
	beta := 25.0 / 6

	pDraw := []float64{
		0.1,
		0.25,
		0.33}
	epsilon := []float64{
		0.74046637542690541,
		1.87760059883033,
		2.5111010132487492,
	}

	for i, p := range pDraw {
		e := GetDrawMargin(p, beta, 2)
		if math.Abs(e-epsilon[i]) > errTolerance {
			t.Error("Expected draw margin for", p, "=", epsilon[i], "got", e)
		}
	}
}
