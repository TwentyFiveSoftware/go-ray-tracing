package main

import "math"

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
