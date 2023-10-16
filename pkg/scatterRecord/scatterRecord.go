package scatterRecord

import (
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/vec3"
)

type ScatterRecord struct {
	Attenuation  vec3.Vec3
	ScatteredRay ray.Ray
}
