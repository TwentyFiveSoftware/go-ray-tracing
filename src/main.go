package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const Width = 800
const Height = 450
const MaxRayTraceDepth = 50
const SamplesPerPixel = 100

func main() {
	camera := NewCamera(Vec3{12.0, 2.0, -3.0}, Vec3{0.0, 0.0, 0.0}, 25.0, 0.0, 10.0)
	scene := GenerateScene()

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: Width, Y: Height}})

	for y := 0; y < Height; y++ {
		fmt.Printf("%d / %d (%.2f%%)\n", y+1, Height, float64(y+1)*100.0/Height)

		for x := 0; x < Width; x++ {
			pixelColor := Vec3{}

			for sample := 0; sample < SamplesPerPixel; sample++ {
				u := (float64(x) + rand.Float64()) / (Width - 1)
				v := (float64(y) + rand.Float64()) / (Height - 1)

				ray := camera.GetRay(u, v)
				pixelColor = pixelColor.Add(CalculateRayColor(scene, ray, MaxRayTraceDepth))
			}

			pixelColor = pixelColor.DivScalar(SamplesPerPixel)
			img.SetRGBA(x, y, colorToRGB(pixelColor))
		}
	}

	saveImage(img)
}

func colorToRGB(pixelColor Vec3) color.RGBA {
	pixelColor = pixelColor.Sqrt()
	pixelColor = pixelColor.Clamp(0.0, 1.0)
	pixelColor = pixelColor.MulScalar(0xFF)

	return color.RGBA{
		R: uint8(pixelColor.X),
		G: uint8(pixelColor.Y),
		B: uint8(pixelColor.Z),
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
