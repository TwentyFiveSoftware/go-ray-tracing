package material

import (
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/scatterRecord"
	"go-ray-tracing/pkg/vec3"
)

type Material interface {
	Scatter(incomingRay ray.Ray, point vec3.Vec3, normal vec3.Vec3, isFrontFace bool) *scatterRecord.ScatterRecord
}
