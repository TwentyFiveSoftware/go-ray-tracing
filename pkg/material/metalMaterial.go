package material

import (
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/scatterRecord"
	"go-ray-tracing/pkg/texture"
	"go-ray-tracing/pkg/vec3"
)

type MetalMaterial struct {
	Texture texture.Texture
}

func (m *MetalMaterial) Scatter(incomingRay ray.Ray, point vec3.Vec3, normal vec3.Vec3, _ bool) *scatterRecord.ScatterRecord {
	scatterDirection := incomingRay.Direction.Normalized().Reflect(normal)

	if normal.Dot(scatterDirection) <= 0.0 {
		return nil
	}

	return &scatterRecord.ScatterRecord{
		Attenuation: m.Texture.GetColorAt(point),
		ScatteredRay: ray.Ray{
			Origin:    point,
			Direction: scatterDirection,
		},
	}
}
