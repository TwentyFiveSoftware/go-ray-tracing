package texture

import (
	"math"

	"go-ray-tracing/pkg/vec3"
)

type CheckeredTexture struct {
	Albedo1 vec3.Vec3
	Albedo2 vec3.Vec3
}

func (t *CheckeredTexture) GetColorAt(point vec3.Vec3) vec3.Vec3 {
	const size = 6.0
	sin := math.Sin(size*point.X) * math.Sin(size*point.Y) * math.Sin(size*point.Z)

	if sin < 0.0 {
		return t.Albedo1
	}
	return t.Albedo2
}
