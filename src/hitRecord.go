package main

type HitRecord struct {
	hit         bool
	t           float64
	point       Vec3
	normal      Vec3
	isFrontFace bool
}
