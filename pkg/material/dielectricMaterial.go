package material

import (
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/scatterRecord"
	"go-ray-tracing/pkg/vec3"
)

type DielectricMaterial struct {
	RefractionIndex float64
}

func (m *DielectricMaterial) Scatter(incomingRay ray.Ray, point vec3.Vec3, normal vec3.Vec3, isFrontFace bool) *scatterRecord.ScatterRecord {
	refractionRatio := m.RefractionIndex
	if isFrontFace {
		refractionRatio = 1.0 / refractionRatio
	}

	scatterDirection := incomingRay.Direction.Normalized().Refract(normal, refractionRatio)

	return &scatterRecord.ScatterRecord{
		Attenuation: vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0},
		ScatteredRay: ray.Ray{
			Origin:    point,
			Direction: scatterDirection,
		},
	}
}
