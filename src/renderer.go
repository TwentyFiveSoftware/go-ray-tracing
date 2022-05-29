package main

import "image/color"

func CalculateRayColor(scene Scene, ray Ray, depth int32) Vec3 {
	if depth <= 0 {
		return Vec3{}
	}

	if hitRecord := scene.RayHitScene(ray); hitRecord.hit {
		if scatterRecord := hitRecord.material.Scatter(ray, hitRecord); scatterRecord.doesScatter {
			return scatterRecord.attenuation.Mul(CalculateRayColor(scene, scatterRecord.scatteredRay, depth-1))
		}

		return Vec3{}
	}

	t := 0.5 * (ray.Direction.Normalized().Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(t))
}

func ColorToRGB(pixelColor Vec3) color.RGBA {
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
