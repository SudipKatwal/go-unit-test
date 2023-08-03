package floatingnumber

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	length := 2.03
	width := 3.5

	exp := 11.06
	p := Perimeter(length, width)
	if exp != p {
		t.Errorf("Expected %f, got %f", exp, p)
	}

	if exp != p {
		t.Errorf("Expected %.18f, got %.18f", exp, p)
	}

	if !withinTolerance(exp, p, 1e-12) {
		t.Errorf("Expected %.18f, got %.18f", exp, p)
	}
}

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}
