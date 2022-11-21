package utils

import "math"

// to calculate natural log
func f(x, a float64) float64 {

	return math.Exp(x) - a
}

// Ln to calculate natual log. ref :https://gist.github.com/thinkphp/ae5024dbd0ea6b83b6308a028ea22323
func Ln(n float64) float64 {
	var lo, hi, m float64

	if n <= 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	EPS := 0.00001
	lo = 0
	hi = n

	for math.Abs(lo-hi) >= EPS {
		m = float64((lo + hi) / 2.0)

		if f(m, n) < 0 {
			lo = m
		} else {
			hi = m
		}
	}

	return float64((lo + hi) / 2.0)
}
