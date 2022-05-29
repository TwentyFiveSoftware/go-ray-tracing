package main

func CalculateRayColor(ray Ray) Vec3 {
	unitDirection := ray.Direction.Normalized()
	t := 0.5 * (unitDirection.Y + 1.0)
	return Vec3{X: 1.0, Y: 1.0, Z: 1.0}.MulScalar(1.0 - t).Add(Vec3{X: 0.5, Y: 0.7, Z: 1.0}.MulScalar(t))
}
