package ray

import "go-ray-tracing/pkg/vec3"

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
}

func (ray *Ray) At(t float64) vec3.Vec3 {
	return ray.Origin.Add(ray.Direction.MulScalar(t))
}
