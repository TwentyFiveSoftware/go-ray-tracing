package sphere

import (
	"math"

	"go-ray-tracing/pkg/hitRecord"
	"go-ray-tracing/pkg/material"
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/vec3"
)

type Sphere struct {
	Center   vec3.Vec3
	Radius   float64
	Material material.Material
}

func (sphere *Sphere) RayHitSphere(ray ray.Ray, tMin float64, tMax float64) *hitRecord.HitRecord {
	oc := ray.Origin.Sub(sphere.Center)
	a := ray.Direction.Dot(ray.Direction)
	halfB := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0.0 {
		return nil
	}

	sqrtD := math.Sqrt(discriminant)
	t := (-halfB - sqrtD) / a
	if t < tMin || t > tMax {
		t = (-halfB + sqrtD) / a

		if t < tMin || t > tMax {
			return nil
		}
	}

	point := ray.At(t)
	normal := point.Sub(sphere.Center).DivScalar(sphere.Radius)

	isFrontFace := ray.Direction.Dot(normal) < 0.0
	if !isFrontFace {
		normal = normal.Neg()
	}

	return &hitRecord.HitRecord{
		T:           t,
		Point:       point,
		Normal:      normal,
		IsFrontFace: isFrontFace,
		Material:    sphere.Material,
	}
}
