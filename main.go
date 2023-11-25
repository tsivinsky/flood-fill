package main

import (
	"flag"
	"log"
	"os"

	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
)

func isSameColor(c1 color.Color, c2 color.Color) bool {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	return r1 == r2 && g1 == g2 && b1 == b2
}

func getDominantColors(img image.Image, n int) map[color.Color]int {
	colors := make(map[color.Color]int)

	bounds := img.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			p := img.At(x, y)
			if _, ok := colors[p]; !ok {
				colors[p] = 1
			} else {
				colors[p] += 1
			}
		}
	}

	dominantColors := make(map[color.Color]int)

	if n == 0 {
		return colors
	}

	for i := 0; i < n; i++ {
		maxN := 0
		var domC color.Color
		for c, n := range colors {
			if n > maxN {
				maxN = n
				domC = c
			}
		}

		dominantColors[domC] = maxN
		delete(colors, domC)
	}

	return dominantColors
}

func getDominantColor(img image.Image) map[color.Color]int {
	return getDominantColors(img, 1)
}

func replaceColors(img image.Image, dominantColors map[color.Color]int, modifyColor modifyColor) *image.RGBA {
	bounds := img.Bounds()
	dstImg := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			p := img.At(x, y)

			if _, ok := dominantColors[p]; ok {
				dstImg.Set(x, y, modifyColor(p))
			} else {
				dstImg.Set(x, y, p)
			}
		}
	}

	return dstImg
}

func readImageFile(fp string) (image.Image, error) {
	f, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveImageFile(img image.Image, fp string) error {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

var (
	numberOfColors = flag.Int("n", 3, "number of dominant colors")
	operation      = flag.String("op", "", "modification to do on colors")
)

func operationToFunc(op string) modifyColor {
	switch op {
	case "c":
		return complementaryColor
	case "bw":
		return bwColor
	default:
		return noopColor
	}
}

func main() {
	flag.Parse()

	fp := flag.Arg(0)
	if fp == "" {
		log.Fatal("no file path provided")
	}

	img, err := readImageFile(fp)
	if err != nil {
		panic(err)
	}

	mod := operationToFunc(*operation)

	dominantColors := getDominantColors(img, *numberOfColors)
	dstImg := replaceColors(img, dominantColors, mod)

	err = saveImageFile(dstImg, "./image.png")
	if err != nil {
		panic(err)
	}
}
