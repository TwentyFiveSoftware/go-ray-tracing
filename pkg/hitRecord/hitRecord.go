package hitRecord

import (
	"go-ray-tracing/pkg/material"
	"go-ray-tracing/pkg/vec3"
)

type HitRecord struct {
	DoesHit     bool
	T           float64
	Point       vec3.Vec3
	Normal      vec3.Vec3
	IsFrontFace bool
	Material    material.Material
}
