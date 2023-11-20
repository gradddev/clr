package clr_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/gradddev/clr"
)

func TestHSVToRGB(t *testing.T) {
	for h1 := 0.0; h1 <= 360; h1 += 1 {
		for s1 := 0.0; s1 <= 100; s1 += 1 {
			for v1 := 0.0; v1 <= 100; v1 += 1 {
				r, g, b := clr.HSVToRGB(h1, s1, v1)
				h2, s2, v2 := clr.RGBToHSV(r, g, b)

				h2 = math.Round(h2)
				s2 = math.Round(s2)
				v2 = math.Round(v2)

				message := fmt.Sprintf(
					"HSVToRGB(%v, %v, %v) => RGBToHSV(%v, %v, %v) => %v, %v, %v",
					h1, s1, v1, r, g, b, h2, s2, v2,
				)
				if v1 == 0 {
					if h2 != 0.0 || s2 != 0.0 || v2 != 0.0 {
						t.Error(message)
					}
				} else {
					if s1 != 0 && h1 != 360 && h1 != h2 {
						t.Error(message)
					}
					if s1 != s2 || v1 != v2 {
						t.Error(message)
					}
				}
			}
		}
	}
}
