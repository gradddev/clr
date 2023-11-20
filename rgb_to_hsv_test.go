package clr_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/gradddev/clr"
)

func TestRGBToHSV(t *testing.T) {
	for r1 := 0.0; r1 <= 255; r1 += 1 {
		for g1 := 0.0; g1 <= 255; g1 += 1 {
			for b1 := 0.0; b1 <= 255; b1 += 1 {
				h, s, v := clr.RGBToHSV(r1, g1, b1)
				r2, g2, b2 := clr.HSVToRGB(h, s, v)

				r2 = math.Round(r2)
				g2 = math.Round(g2)
				b2 = math.Round(b2)

				message := fmt.Sprintf(
					"RGBToHSV(%v, %v, %v) => HSVToRGB(%v, %v, %v) => %v, %v, %v",
					r1, g1, b1, h, s, v, r2, g2, b2,
				)
				if r1 != r2 || g1 != g2 || b1 != b2 {
					t.Error(message)
				}
			}
		}
	}
}
