// Package wilks provides the ability to calculate a Wilks Score for the purpose or ranking
// powerlifters across gender and weight class.
//
// See: https://en.wikipedia.org/wiki/Wilks_Coefficient
package wilks

import (
	"math"
)

const poundsPerKilo = 2.20462

type coefficientVariables struct {
	dividend float64

	a float64
	b float64
	c float64
	d float64
	e float64
	f float64
}

var (
	maleVars   = coefficientVariables{500, -216.0475144, 16.2606339, -0.002388645, -0.00113732, 7.01863E-06, -1.291E-08}
	femaleVars = coefficientVariables{500, 594.31747775582, -27.23842536447, 0.82112226871, -0.00930733913, 4.731582E-05, -9.054E-08}
)

// PoundsToKilos converts a value from pounds to kilograms.
func PoundsToKilos(lbs float64) float64 {
	return lbs / poundsPerKilo
}

// Score returns the wilks score for a given gender, bodyweight, and total.
func Score(isFemale bool, bw, total float64) float64 {
	return coefficient(isFemale, bw) * total
}

// coefficient returns the coefficient to use for a particular gender and bodyweight.
func coefficient(isFemale bool, bw float64) float64 {
	c := maleVars
	if isFemale {
		c = femaleVars
	}

	// See: https://en.wikipedia.org/wiki/Wilks_Coefficient
	// coeff = 500 / (a + bx + cx^2 + dx^3 + ex^4 + fx^5)
	return c.dividend / (c.a + (c.b * bw) + (c.c * math.Pow(bw, 2)) + (c.d * math.Pow(bw, 3)) + (c.e * math.Pow(bw, 4)) + (c.f * math.Pow(bw, 5)))
}
