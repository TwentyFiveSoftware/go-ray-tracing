package main

func CalculateRayColor(scene Scene, ray Ray, depth int32) Vec3 {
	if depth <= 0 {
		return Vec3{}
	}

	if hitRecord := scene.RayHitScene(ray); hitRecord.hit {
		return hitRecord.normal.Add(Vec3{1, 1, 1}).MulScalar(0.5)
	}

	unitDirection := ray.Direction.Normalized()
	t := 0.5 * (unitDirection.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(t))
}
