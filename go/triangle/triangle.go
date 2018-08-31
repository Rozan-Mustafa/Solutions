package triangle

import (
	"math"
	"sort"
)

// Kind is an enumerated type for triangle kinds
type Kind string

const (
	// NaT indicates it is not a triangule
	NaT = Kind("NaT")
	// Equ represents an equilateral triangule
	Equ = Kind("Equ")
	// Iso represents an isosceles triangule
	Iso = Kind("Iso")
	// Sca represents a scalene triangule
	Sca = Kind("Sca")
	// Deg represents a degenerate triangule
	Deg = Kind("Deg")
)

/*KindFromSides determines the type of triangle formed by sides a, b and c
the triangle is invaid if the longest side is equal or longer than the shorter sides
or any of the sides are invalid (0 or less, NaN or Inf+)*/
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(sides))
	switch {
	case math.IsNaN(sides[0]) || sides[0] <= 0 || sides[2] == math.Inf(1) ||
		sides[1]+sides[0] < sides[2]:
		return NaT
	case sides[0] == sides[1] && sides[1] == sides[2]:
		return Equ
	case sides[0] == sides[1] || sides[1] == sides[2]:
		return Iso
	}
	return Sca
}
