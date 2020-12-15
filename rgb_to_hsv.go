package clr

import "math"

func RGBToHSV(r, g, b float64) (h, s, v float64) {
	r, g, b = r/255, g/255, b/255

	v = math.Max(math.Max(r, g), b)
	c := v - math.Min(math.Min(r, g), b)
	if c != 0 {
		if v == r {
			h = (g - b) / c
		} else if v == g {
			h = 2 + (b-r)/c
		} else {
			h = 4 + (r-g)/c
		}
	}

	if h < 0 {
		h = h + 6
	}

	if v != 0 {
		s = c / v
	}

	return h * 60, s * 100, v * 100
}
