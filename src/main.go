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
	camera := NewCamera(Vec3{X: 12.0, Y: 2.0, Z: -3.0}, Vec3{X: 0.0, Y: 0.0, Z: 0.0}, 25.0, 0.0, 10.0)

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: WIDTH, Y: HEIGHT}})

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			u := float64(x) / (WIDTH - 1)
			v := float64(y) / (HEIGHT - 1)

			ray := camera.GetRay(u, v)
			pixelColor := CalculateRayColor(ray)
			img.SetRGBA(x, y, colorToRGB(pixelColor))
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
