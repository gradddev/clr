package clr

type HSV struct {
	Hue, Saturation, Value float64
}

func NewHSV(hue, saturation, value float64) HSV {
	if hue < 0 || hue == 360 || value == 0 {
		hue = 0
	} else if hue > 360 {
		hue = 360
	}

	if saturation < 0 || value == 0 {
		saturation = 0
	} else if saturation > 100 {
		saturation = 100
	}

	if value < 0 {
		value = 0
	} else if value > 100 {
		value = 100
	}

	return HSV{hue, saturation, value}
}

func (hsv HSV) ToRGB() RGB {
	r, g, b := HSVToRGB(hsv.Hue, hsv.Saturation, hsv.Value)
	return RGB{r, g, b}
}
