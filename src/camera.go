package main

import "math"

type Camera struct {
	lookFrom            Vec3
	upperLeftCorner     Vec3
	horizontalDirection Vec3
	verticalDirection   Vec3
	up                  Vec3
	right               Vec3
}

func NewCamera(lookFrom Vec3, lookAt Vec3, fov float64, focusDistance float64) Camera {
	const aspectRatio = float64(Width) / float64(Height)

	viewportHeight := math.Tan(fov/360.0*math.Pi) * 2.0
	viewportWidth := viewportHeight * aspectRatio

	forward := lookAt.Sub(lookFrom).Normalized()
	right := Vec3{0.0, 1.0, 0.0}.Cross(forward).Normalized()
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

func (camera *Camera) GetRay(u float64, v float64) Ray {
	target := camera.upperLeftCorner.Add(camera.horizontalDirection.MulScalar(u)).Sub(camera.verticalDirection.MulScalar(v))

	return Ray{
		Origin:    camera.lookFrom,
		Direction: target.Sub(camera.lookFrom),
	}
}
