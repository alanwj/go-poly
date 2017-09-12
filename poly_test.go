package poly

import (
	"math"
	"testing"
)

func comparePoly(p, q Poly) bool {
	if p.Deg() != q.Deg() {
		return false
	}
	for i := 0; i <= p.Deg(); i++ {
		if math.Abs(p.Coeff(i)-q.Coeff(i)) > 0.00001 {
			return false
		}
	}
	return true
}

// Tests that the degree of various polynomials is reported as expected.
func TestDeg(t *testing.T) {
	cases := []struct {
		p    Poly
		want int
	}{
		{New(), 0},
		{New(1), 0},
		{New(1, 2), 1},
		{New(1, 2, 3), 2},
		{New(0, 0, 0), 0},
	}
	for i, c := range cases {
		if got := c.p.Deg(); got != c.want {
			t.Errorf("case %d: Deg() on %q == %d, want %d", i, c.p, got, c.want)
		}
	}
}

// Tests the degree of a zero valued Poly.
func TestDegZero(t *testing.T) {
	var p Poly
	want := 0
	if got := p.Deg(); got != want {
		t.Errorf("Deg() == %d, want %d", got, want)
	}
}

// Tests that the coefficients of various terms are correctly reported.
func TestCoeff(t *testing.T) {
	c := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	p := New(c...)
	for i, want := range c {
		if got := p.Coeff(i); got != want {
			t.Errorf("Coeff(%d) == %f, want %f", i, got, want)
		}
	}
}

// Tests the coefficient of a zero valued Poly.
func TestCoeffZero(t *testing.T) {
	var p Poly
	want := float64(0)
	if got := p.Coeff(0); got != want {
		t.Errorf("Coeff(0) == %f, want %f", got, want)
	}
}

// Tests that out of range coefficients are zero.
func TestCoeffOutOfRange(t *testing.T) {
	cases := []struct {
		i    int
		want float64
	}{
		{-2, 0.0},
		{-1, 0.0},
		{0, 1.0},
		{1, 2.0},
		{2, 3.0},
		{3, 0.0},
		{4, 0.0},
	}
	p := New(1, 2, 3)
	for i, c := range cases {
		if got := p.Coeff(c.i); got != c.want {
			t.Errorf("case %d: Coeff(%d) == %f, want %f", i, c.i, got, c.want)
		}
	}
}

// Tests that function evaluation produces correct results.
func TestEval(t *testing.T) {
	cases := []struct {
		p    Poly
		x    float64
		want float64
	}{
		{Poly{}, 0.0, 0.0},
		{New(), 0.0, 0.0},
		{New(0, 1, 2), 0.0, 0.0},
		{New(1, 2, 3), 0.0, 1.0},
		{New(-1, 2, -3), 2.5, -14.75},
	}
	for i, c := range cases {
		if got := c.p.Eval(c.x); math.Abs(c.want-got) > 0.00001 {
			t.Errorf("case %d: Eval(%f) on %q == %f, want %f", i, c.x, c.p, got, c.want)
		}
	}
}

// Tests that polynomials add correctly.
func TestAdd(t *testing.T) {
	cases := []struct {
		p    Poly
		q    Poly
		want Poly
	}{
		{Poly{}, Poly{}, Poly{}},
		{New(), New(), New()},
		{New(1, 2), Poly{}, New(1, 2)},
		{Poly{}, New(1, 2), New(1, 2)},
		{New(1, 2), New(3, 4), New(4, 6)},
		{New(1, 2, 3), New(3, 4), New(4, 6, 3)},
		{New(1, 2), New(3, 4, 5), New(4, 6, 5)},
		{New(1, 2, 3), New(-1, 2, -3), New(0, 4)},
	}
	for i, c := range cases {
		if got := c.p.Add(c.q); !comparePoly(got, c.want) {
			t.Errorf("case %d: Add(%q) on %q == %q, want %q", i, c.q, c.p, got, c.want)
		}
	}
}

// Tests that polynomials subtract correctly.
func TestSub(t *testing.T) {
	cases := []struct {
		p    Poly
		q    Poly
		want Poly
	}{
		{Poly{}, Poly{}, Poly{}},
		{New(), New(), New()},
		{New(1, 2), Poly{}, New(1, 2)},
		{Poly{}, New(1, 2), New(-1, -2)},
		{New(1, 2), New(3, 4), New(-2, -2)},
		{New(1, 2, 3), New(3, 4), New(-2, -2, 3)},
		{New(1, 2), New(3, 4, 5), New(-2, -2, -5)},
		{New(1, 4, 3), New(1, 2, 3), New(0, 2)},
	}
	for i, c := range cases {
		if got := c.p.Sub(c.q); !comparePoly(got, c.want) {
			t.Errorf("case %d: Sub(%q) on %q == %q, want %q", i, c.q, c.p, got, c.want)
		}
	}
}

// Tests that polynomials multiply correctly.
func TestMul(t *testing.T) {
	cases := []struct {
		p    Poly
		q    Poly
		want Poly
	}{
		{Poly{}, Poly{}, Poly{}},
		{New(), New(), New()},
		{New(1, 2), Poly{}, Poly{}},
		{Poly{}, New(1, 2), Poly{}},
		{New(2, 1), New(-2, 1), New(-4, 0, 1)},
		{New(1, 2), New(3, 4), New(3, 10, 8)},
		{New(1, 2, 3), New(3, 4), New(3, 10, 17, 12)},
		{New(3, 4), New(1, 2, 3), New(3, 10, 17, 12)},
	}
	for i, c := range cases {
		if got := c.p.Mul(c.q); !comparePoly(got, c.want) {
			t.Errorf("case %d: Mul(%q) on %q == %q, want %q", i, c.q, c.p, got, c.want)
		}
	}
}

func TestMod(t *testing.T) {
	cases := []struct {
		p    Poly
		q    Poly
		want Poly
	}{
		{Poly{}, Poly{}, Poly{}},
		{New(), New(), New()},
		{Poly{}, New(1, 2), Poly{}},
		{New(2, 1), New(-2, 1), New(4)},
		{New(3, 4), New(1, 2), New(1)},
		{New(1, 2, 3), New(3, 4), New(1, -0.25)},
		{New(3, 4), New(1, 2, 3), New(3, 4)},
	}
	for i, c := range cases {
		if got := c.p.Mod(c.q); !comparePoly(got, c.want) {
			t.Errorf("case %d: Mod(%q) on %q == %q, want %q", i, c.q, c.p, got, c.want)
		}
	}
}

// Tests that derivatives are computed correctly.
func TestDer(t *testing.T) {
	cases := []struct {
		p    Poly
		want Poly
	}{
		{Poly{}, Poly{}},
		{New(), New()},
		{New(1), New(0)},
		{New(1, 2), New(2)},
		{New(1, 2, 3), New(2, 6)},
	}
	for i, c := range cases {
		if got := c.p.Der(); !comparePoly(got, c.want) {
			t.Errorf("case %d: Der() on %q == %q, want %q", i, c.p, got, c.want)
		}
	}
}

// Tests that integrals are computed correctly.
func TestInt(t *testing.T) {
	cases := []struct {
		p    Poly
		k    float64
		want Poly
	}{
		{Poly{}, 0, Poly{}},
		{New(), 0, New()},
		{Poly{}, 3, New(3)},
		{New(1), 4, New(4, 1)},
		{New(1, 4), 5, New(5, 1, 2)},
	}
	for i, c := range cases {
		if got := c.p.Int(c.k); !comparePoly(got, c.want) {
			t.Errorf("case %d: Int(%.3f) on %q == %q, want %q", i, c.k, c.p, got, c.want)
		}
	}
}

// Tests that the string representation is correct.
func TestString(t *testing.T) {
	cases := []struct {
		p    Poly
		want string
	}{
		{Poly{}, "0.000"},
		{New(), "0.000"},
		{New(1.234), "1.234"},
		{New(-1.234), "-1.234"},
		{New(0, 1), "x"},
		{New(0, -1), "-x"},
		{New(0, 2), "2.000x"},
		{New(0, -2), "-2.000x"},
		{New(0, 0, 1), "x^2"},
		{New(0, 0, -1), "-x^2"},
		{New(0, 0, 2), "2.000x^2"},
		{New(0, 0, -2), "-2.000x^2"},
		{New(0, 1, 1), "x^2 + x"},
		{New(0, 2, 2), "2.000x^2 + 2.000x"},
		{New(0, -1, -1), "-x^2 - x"},
		{New(0, -2, -2), "-2.000x^2 - 2.000x"},
		{New(-3, -1, 2, 0, 4), "4.000x^4 + 2.000x^2 - x - 3.000"},
	}
	for i, c := range cases {
		if got := c.p.String(); got != c.want {
			t.Errorf("case %d: String() on %q == %q, want %q", i, c.p, got, c.want)
		}
	}
}
