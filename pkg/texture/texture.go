package texture

import "go-ray-tracing/pkg/vec3"

type Texture interface {
	GetColorAt(point vec3.Vec3) vec3.Vec3
}
