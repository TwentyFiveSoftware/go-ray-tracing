package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const WIDTH = 800
const HEIGHT = 450

func main() {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: WIDTH, Y: HEIGHT}})

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			u := float64(x) / (WIDTH - 1)
			v := float64(y) / (HEIGHT - 1)

			pixelColor := colorToRGB(Vec3{X: u, Y: v, Z: 0.25})
			img.SetRGBA(x, y, pixelColor)
		}
	}

	saveImage(img)
}

func colorToRGB(pixelColor Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(pixelColor.X * 0xFF),
		G: uint8(pixelColor.Y * 0xFF),
		B: uint8(pixelColor.Z * 0xFF),
		A: 0xFF,
	}
}

func saveImage(image image.Image) {
	f, err := os.Create("render.png")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = png.Encode(f, image)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
}
