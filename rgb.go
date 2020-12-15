package clr

import (
	"fmt"
)

type RGB struct {
	Red, Green, Blue float64
}

func NewRGB(red, green, blue float64) RGB {
	if red < 0 {
		red = 0
	}
	if red > 255 {
		red = 255
	}

	if green < 0 {
		green = 0
	}
	if green > 255 {
		green = 255
	}

	if blue < 0 {
		blue = 0
	}
	if blue > 255 {
		blue = 255
	}

	return RGB{red, green, blue}
}

func NewRGBFromHex(hex string) RGB {
	rgb := RGB{}

	if hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 3 && len(hex) != 6 {
		return rgb
	}

	var format string
	var factor int
	if len(hex) == 3 {
		format = "%1x%1x%1x"
		factor = 16
	} else if len(hex) == 6 {
		format = "%02x%02x%02x"
		factor = 0
	} else {
		return rgb
	}

	var red, green, blue int
	n, err := fmt.Sscanf(hex, format, &red, &green, &blue)
	if n != 3 || err != nil {
		return rgb
	}

	red += red * factor
	green += green * factor
	blue += blue * factor

	rgb.Red, rgb.Green, rgb.Blue = float64(red), float64(green), float64(blue)

	return rgb
}

func (rgb RGB) ToHex() string {
	r, g, b := uint8(rgb.Red), uint8(rgb.Green), uint8(rgb.Blue)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (rgb RGB) ToHSV() HSV {
	h, s, v := RGBToHSV(rgb.Red, rgb.Green, rgb.Blue)
	return HSV{h, s, v}
}
