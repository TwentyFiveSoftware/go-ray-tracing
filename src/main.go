package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const Width = 800
const Height = 450
const MaxRayTraceDepth = 50

func main() {
	camera := NewCamera(Vec3{12.0, 2.0, -3.0}, Vec3{0.0, 0.0, 0.0}, 25.0, 0.0, 10.0)
	scene := GenerateScene()

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: Width, Y: Height}})

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			u := float64(x) / (Width - 1)
			v := float64(y) / (Height - 1)

			ray := camera.GetRay(u, v)
			pixelColor := CalculateRayColor(scene, ray, MaxRayTraceDepth)
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
