package renderer

import (
	"image/color"
	"math/rand"

	"go-ray-tracing/pkg/camera"
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/scene"
	"go-ray-tracing/pkg/vec3"
)

func RenderRow(y int, camera camera.Camera, scene scene.Scene, width int, height int, samplesPerPixel int, maxRayTraceDepth int) []color.RGBA {
	var row []color.RGBA

	for x := 0; x < width; x++ {
		pixelColor := vec3.Vec3{}

		for sample := 0; sample < samplesPerPixel; sample++ {
			u := (float64(x) + rand.Float64()) / (float64(width) - 1)
			v := (float64(y) + rand.Float64()) / (float64(height) - 1)

			pixelColor = pixelColor.Add(calculateRayColor(scene, camera.GetRay(u, v), maxRayTraceDepth))
		}

		pixelColor = pixelColor.DivScalar(float64(samplesPerPixel))
		row = append(row, colorToRGB(pixelColor))
	}

	return row
}

func calculateRayColor(scene scene.Scene, ray ray.Ray, depth int) vec3.Vec3 {
	if depth <= 0 {
		return vec3.Vec3{}
	}

	hitRecord := scene.RayHitScene(ray)
	if hitRecord != nil {
		scatterRecord := hitRecord.Material.Scatter(ray, hitRecord.Point, hitRecord.Normal, hitRecord.IsFrontFace)
		if scatterRecord == nil {
			return vec3.Vec3{}
		}

		return scatterRecord.Attenuation.Mul(calculateRayColor(scene, scatterRecord.ScatteredRay, depth-1))
	}

	t := 0.5 * (ray.Direction.Normalized().Y + 1.0)
	return vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}.MulScalar(1.0 - t).Add(vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}.MulScalar(t))
}

func colorToRGB(pixelColor vec3.Vec3) color.RGBA {
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
