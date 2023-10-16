package camera

import (
	"math"

	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/vec3"
)

type Camera struct {
	lookFrom            vec3.Vec3
	upperLeftCorner     vec3.Vec3
	horizontalDirection vec3.Vec3
	verticalDirection   vec3.Vec3
	up                  vec3.Vec3
	right               vec3.Vec3
}

func NewCamera(lookFrom vec3.Vec3, lookAt vec3.Vec3, fov float64, focusDistance float64, width int, height int) Camera {
	aspectRatio := float64(width) / float64(height)
	viewportHeight := math.Tan(fov/360.0*math.Pi) * 2.0
	viewportWidth := viewportHeight * aspectRatio

	forward := lookAt.Sub(lookFrom).Normalized()
	right := vec3.Vec3{Y: 1.0}.Cross(forward).Normalized()
	up := forward.Cross(right).Normalized()

	horizontalDirection := right.MulScalar(viewportWidth * focusDistance)
	verticalDirection := up.MulScalar(viewportHeight * focusDistance)

	upperLeftCorner := lookFrom.
		Sub(horizontalDirection.MulScalar(0.5)).
		Add(verticalDirection.MulScalar(0.5)).
		Add(forward.MulScalar(focusDistance))

	return Camera{
		lookFrom,
		upperLeftCorner,
		horizontalDirection,
		verticalDirection,
		up,
		right,
	}
}

func (camera *Camera) GetRay(u float64, v float64) ray.Ray {
	target := camera.upperLeftCorner.Add(camera.horizontalDirection.MulScalar(u)).Sub(camera.verticalDirection.MulScalar(v))

	return ray.Ray{
		Origin:    camera.lookFrom,
		Direction: target.Sub(camera.lookFrom),
	}
}
