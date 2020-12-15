package clr_test

import (
	"testing"

	"github.com/AlexeySemigradsky/clr"
)

func TestNewRGBFromHex(t *testing.T) {
	rgb1 := clr.NewRGBFromHex("#fff")
	if rgb1.Red != 255.0 ||
		rgb1.Green != 255.0 ||
		rgb1.Blue != 255.0 {
		t.Error("#fff != 255, 255, 255")
	}

	rgb2 := clr.NewRGBFromHex("#ffffff")
	if rgb2.Red != 255.0 ||
		rgb2.Green != 255.0 ||
		rgb2.Blue != 255.0 {
		t.Error("#ffffff != 255, 255, 255")
	}

	rgb1 = clr.NewRGBFromHex("#000")
	if rgb1.Red != 0.0 ||
		rgb1.Green != 0.0 ||
		rgb1.Blue != 0.0 {
		t.Error("#000 != 0, 0, 0")
	}

	rgb2 = clr.NewRGBFromHex("#00Ff00")
	if rgb2.Red != 0.0 ||
		rgb2.Green != 255.0 ||
		rgb2.Blue != 0.0 {
		t.Error("#ffffff != 0, 255, 0")
	}

	rgb1 = clr.NewRGBFromHex("#fff")
	rgb2 = clr.NewRGBFromHex("#ffffff")
	if rgb1.Red != rgb2.Red ||
		rgb1.Green != rgb2.Green ||
		rgb1.Blue != rgb2.Blue {
		t.Error("#fff != #fffFFF")
	}

	rgb1 = clr.NewRGBFromHex("#fff")
	rgb2 = clr.NewRGBFromHex("#FFF")
	if rgb1.Red != rgb2.Red ||
		rgb1.Green != rgb2.Green ||
		rgb1.Blue != rgb2.Blue {
		t.Error("#fff != #FFF")
	}

	rgb1 = clr.NewRGBFromHex("#fff")
	rgb2 = clr.NewRGBFromHex("#fffFFF")
	if rgb1.Red != rgb2.Red ||
		rgb1.Green != rgb2.Green ||
		rgb1.Blue != rgb2.Blue {
		t.Error("#fff != #FFF")
	}

	hex1 := "#ffffff"
	rgb1 = clr.NewRGBFromHex(hex1)
	hex2 := rgb1.ToHex()
	rgb2 = clr.NewRGBFromHex(hex2)
	if rgb1.Red != rgb2.Red ||
		rgb1.Green != rgb2.Green ||
		rgb1.Blue != rgb2.Blue {
		t.Errorf("%v != %v", rgb1, rgb2)
	}
}
