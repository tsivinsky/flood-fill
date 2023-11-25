package main

import "image/color"

type modifyColor = func(c color.Color) color.Color

func complementaryColor(c color.Color) color.Color {
	r, g, b, a := c.RGBA()

	cc := color.RGBA{
		R: uint8(255 - r),
		G: uint8(255 - g),
		B: uint8(255 - b),
		A: uint8(a),
	}

	return cc
}

func bwColor(c color.Color) color.Color {
	return color.GrayModel.Convert(c)
}

func noopColor(c color.Color) color.Color {
	return c
}
