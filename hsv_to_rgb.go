package clr

import "math"

func HSVToRGB(h, s, v float64) (r, g, b float64) {
	s, v = s/100, v/100
	f := func(n float64) float64 {
		k := math.Mod(n+h/60, 6)
		return v - v*s*math.Max(math.Min(math.Min(k, 4-k), 1), 0)
	}
	return f(5) * 255, f(3) * 255, f(1) * 255
}
