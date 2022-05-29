package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
)

const Width = 800
const Height = 450
const MaxRayTraceDepth = 50
const SamplesPerPixel = 1

func main() {
	camera := NewCamera(Vec3{12.0, 2.0, -3.0}, Vec3{0.0, 0.0, 0.0}, 25.0, 10.0)
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
			img.SetRGBA(x, y, ColorToRGB(pixelColor))
		}
	}

	saveImage(img)
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

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
