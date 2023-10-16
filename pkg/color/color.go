package color

import (
	"math"
	"math/rand"

	"go-ray-tracing/pkg/vec3"
)

func GetRandomColor() vec3.Vec3 {
	return hsvToRGB(rand.Float64()*360.0, 0.75, 0.45)
}

func hsvToRGB(h float64, s float64, v float64) vec3.Vec3 {
	c := s * v
	x := c * (1.0 - math.Abs(math.Mod(h/60.0, 2.0)-1.0))
	m := v - c

	var r float64
	var g float64
	var b float64

	if h >= 0.0 && h < 60.0 {
		r = c
		g = x
		b = 0.0
	} else if h >= 60.0 && h < 120.0 {
		r = x
		g = c
		b = 0.0
	} else if h >= 120.0 && h < 180.0 {
		r = 0.0
		g = c
		b = x
	} else if h >= 180.0 && h < 240.0 {
		r = 0.0
		g = x
		b = c
	} else if h >= 240.0 && h < 300.0 {
		r = x
		g = 0.0
		b = c
	} else {
		r = c
		g = 0.0
		b = x
	}

	return vec3.Vec3{X: r + m, Y: g + m, Z: b + m}
}
