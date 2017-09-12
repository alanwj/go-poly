// The poly package provides types and functions for manipulating polynomials.
package poly

import (
	"bytes"
	"fmt"
	"math"
)

// Poly represents a polynomial of arbitrary degree.
// A zero valued Poly is equivalent to 0.0.
type Poly struct {
	coeff []float64
}

// Returns the coefficient array for a Poly.
// This level of indirection makes implementing the zero valued Poly easier.
func (p Poly) co() []float64 {
	if len(p.coeff) == 0 {
		return []float64{0}
	}
	return p.coeff
}

// Returns a normalized polynomial with the given coefficients.
// All leading terms with degree greater than 0 and coefficients that are zero
// within the precision of a float64 are removed.
func normalized(c []float64) Poly {
	i := len(c) - 1
	for i > 0 && c[i] == 0.0 {
		i--
	}
	return Poly{c[0 : i+1]}
}

// Creates a new Poly.
// The ith parameter represents the coefficient of x^i.
// Example:
//   p := poly.New(1.5, 2.3, 3.7)
//
//   This represents 1.5 + 2.3*x + 3.7*x^2
func New(c ...float64) Poly {
	if len(c) == 0 {
		return Poly{[]float64{0.0}}
	}
	a := make([]float64, len(c))
	copy(a, c)
	return normalized(a)
}

// Returns the highest degree of the polynomial's highest order term.
func (p Poly) Deg() int {
	return len(p.co()) - 1
}

// Returns the coefficient of the ith order term.
func (p Poly) Coeff(i int) float64 {
	if i < 0 || i > p.Deg() {
		return 0.0
	}
	return p.co()[i]
}

// Evaluates a polynomial at the given point x.
func (p Poly) Eval(x float64) float64 {
	var n float64
	for i, c := range p.co() {
		n += c * math.Pow(x, float64(i))
	}
	return n
}

// Adds a polynomial to another polynomial.
// Returns p+q.
func (p Poly) Add(q Poly) Poly {
	pco := p.co()
	plen := len(pco)
	qco := q.co()
	qlen := len(qco)
	if plen < qlen {
		return q.Add(p)
	}

	c := make([]float64, plen)

	pco = p.co()
	for i, qc := range qco {
		c[i] = pco[i] + qc
	}

	for i := qlen; i < plen; i++ {
		c[i] = pco[i]
	}

	return normalized(c)
}

// Subtracts a polynomial from another polynomial.
// Returns p-q.
func (p Poly) Sub(q Poly) Poly {
	qco := q.co()
	qlen := len(qco)
	c := make([]float64, qlen)
	for i, qc := range qco {
		c[i] = -qc
	}
	return p.Add(Poly{c})
}

// Multiplies a polynomial by another polynomial.
// Returns p*q.
func (p Poly) Mul(q Poly) Poly {
	pco := p.co()
	plen := len(pco)
	qco := q.co()
	qlen := len(qco)
	c := make([]float64, plen+qlen-1)
	for i, pc := range pco {
		for j, qc := range qco {
			c[i+j] += pc * qc
		}
	}
	return normalized(c)
}

// use Euclidean division algorithm to find remainder (the mod)
func (p Poly) Mod(q Poly) Poly {
  r := p
  d := q.Deg()
  c := q.Coeff(q.Deg())
  if p.Deg() >= d {
    sT := make([]float64, r.Deg()-d + 1)
    sT[len(sT)-1] = r.Coeff(r.Deg())/c
    s := New(sT...)
    r = r.Sub(s.Mul(q))
  }
  return r
}

// Computes the derivative of a polynomial.
func (p Poly) Der() Poly {
	pco := p.co()
	plen := len(pco)
	c := make([]float64, plen-1)
	for i, pc := range pco {
		if i > 0 {
			c[i-1] = pc * float64(i)
		}
	}
	return normalized(c)
}

// Computes the definite integral of a polynomial.
// The provided constant k will be used as the 0th order term of the result.
func (p Poly) Int(k float64) Poly {
	pco := p.co()
	plen := len(pco)
	c := make([]float64, plen+1)
	c[0] = k
	for i, pc := range pco {
		c[i+1] = pc / float64(i+1)
	}
	return normalized(c)
}

// Returns a printable string representing the polynomial value.
func (p Poly) String() string {
	var buffer bytes.Buffer

	pco := p.co()
	plen := len(pco)

	first := true
	for i := plen; i > 0; i-- {
		e := i - 1
		absc := math.Abs(pco[e])
		if absc < 0.0001 && !(first && e == 0) {
			continue
		}

		c := pco[e]
		if !first {
			if c < 0 {
				buffer.WriteString(" - ")
			} else {
				buffer.WriteString(" + ")
			}
			c = absc
		}
		if absc != 1.0 || e == 0 {
			buffer.WriteString(fmt.Sprintf("%.3f", c))
		} else if c == -1.0 && first {
			buffer.WriteString("-")
		}
		if e != 0 {
			buffer.WriteString("x")
			if e != 1 {
				buffer.WriteString(fmt.Sprintf("^%d", e))
			}
		}
		first = false
	}
	return buffer.String()
}
