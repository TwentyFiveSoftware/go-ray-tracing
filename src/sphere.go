package main

import "math"

type Sphere struct {
	center Vec3
	radius float64
}

func (sphere *Sphere) RayHitSphere(ray Ray, tMin float64, tMax float64) HitRecord {
	oc := ray.Origin.Sub(sphere.center)
	a := ray.Direction.Dot(ray.Direction)
	halfB := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - sphere.radius*sphere.radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0.0 {
		return HitRecord{}
	}

	sqrtD := math.Sqrt(discriminant)
	t := (-halfB - sqrtD) / a
	if t < tMin || t > tMax {
		t = (-halfB + sqrtD) / a

		if t < tMin || t > tMax {
			return HitRecord{}
		}
	}

	point := ray.At(t)
	normal := point.Sub(sphere.center).DivScalar(sphere.radius)
	isFrontFace := ray.Direction.Dot(normal) < 0.0

	if !isFrontFace {
		normal = normal.Neg()
	}

	return HitRecord{true, t, point, normal, isFrontFace}
}
