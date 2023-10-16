package material

import (
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/scatterRecord"
	"go-ray-tracing/pkg/texture"
	"go-ray-tracing/pkg/vec3"
)

type DiffuseMaterial struct {
	Texture texture.Texture
}

func (m *DiffuseMaterial) Scatter(_ ray.Ray, point vec3.Vec3, normal vec3.Vec3, _ bool) scatterRecord.ScatterRecord {
	scatterDirection := normal.Add(vec3.RandomUnitVector())

	if scatterDirection.IsNearZero() {
		scatterDirection = normal
	}

	return scatterRecord.ScatterRecord{
		DoesScatter: true,
		Attenuation: m.Texture.GetColorAt(point),
		ScatteredRay: ray.Ray{
			Origin:    point,
			Direction: scatterDirection,
		},
	}
}
