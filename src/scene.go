package main

import "math"

type Scene struct {
	spheres []Sphere
}

func GenerateScene() Scene {
	var spheres []Sphere

	spheres = append(spheres, Sphere{
		center: Vec3{0, 0, -1},
		radius: 0.5,
		material: Material{
			materialType:    MaterialTypeDielectric,
			refractionIndex: 1.5,
		},
	})

	return Scene{spheres}
}

func (scene *Scene) RayHitScene(ray Ray) HitRecord {
	currentRecord := HitRecord{t: math.Inf(1)}

	for _, sphere := range scene.spheres {
		if hitRecord := sphere.RayHitSphere(ray, 0.001, currentRecord.t); hitRecord.hit {
			currentRecord = hitRecord
		}
	}

	return currentRecord
}
