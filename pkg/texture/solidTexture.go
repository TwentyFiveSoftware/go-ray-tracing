package texture

import "go-ray-tracing/pkg/vec3"

type SolidTexture struct {
	Albedo vec3.Vec3
}

func (t *SolidTexture) GetColorAt(vec3.Vec3) vec3.Vec3 {
	return t.Albedo
}
