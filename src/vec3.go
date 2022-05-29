package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Neg() Vec3 {
	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) Mul(other Vec3) Vec3 {
	return Vec3{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

func (v Vec3) MulScalar(scalar float64) Vec3 {
	return Vec3{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func (v Vec3) Div(other Vec3) Vec3 {
	return Vec3{
		X: v.X / other.X,
		Y: v.Y / other.Y,
		Z: v.Z / other.Z,
	}
}

func (v Vec3) DivScalar(scalar float64) Vec3 {
	return Vec3{
		X: v.X / scalar,
		Y: v.Y / scalar,
		Z: v.Z / scalar,
	}
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Normalized() Vec3 {
	return v.DivScalar(v.Length())
}

func (v Vec3) Sqrt() Vec3 {
	return Vec3{
		X: math.Sqrt(v.X),
		Y: math.Sqrt(v.Y),
		Z: math.Sqrt(v.Z),
	}
}

func (v Vec3) Clamp(min float64, max float64) Vec3 {
	return Vec3{
		X: math.Min(max, math.Max(min, v.X)),
		Y: math.Min(max, math.Max(min, v.Y)),
		Z: math.Min(max, math.Max(min, v.Z)),
	}
}

func (v Vec3) IsNearZero() bool {
	const EPSILON = 1e-8
	return math.Abs(v.X) < EPSILON && math.Abs(v.Y) < EPSILON && math.Abs(v.Z) < EPSILON
}

func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Sub(normal.MulScalar(2.0 * v.Dot(normal)))
}

func (v Vec3) Refract(normal Vec3, refractionRatio float64) Vec3 {
	cosTheta := math.Min(normal.Dot(v.Neg()), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	r0 := (1.0 - refractionRatio) / (1.0 + refractionRatio)
	reflectance := r0*r0 + (1.0-r0*r0)*math.Pow(1.0-cosTheta, 5)

	if refractionRatio*sinTheta > 1.0 || reflectance > rand.Float64() {
		return v.Reflect(normal)
	}

	rOutPerpendicular := v.Add(normal.MulScalar(cosTheta)).MulScalar(refractionRatio)
	rOutParallel := normal.MulScalar(-math.Sqrt(1.0 - rOutPerpendicular.Dot(rOutPerpendicular)))
	return rOutPerpendicular.Add(rOutParallel)
}

func RandomUnitVector() Vec3 {
	for {
		vector := Vec3{
			X: rand.Float64()*2.0 - 1.0,
			Y: rand.Float64()*2.0 - 1.0,
			Z: rand.Float64()*2.0 - 1.0,
		}

		if vector.Dot(vector) < 1.0 {
			return vector.Normalized()
		}
	}
}
