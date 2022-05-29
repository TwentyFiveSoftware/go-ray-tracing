package main

import (
	"math"
	"math/rand"
)

type Scene struct {
	spheres []Sphere
}

func GenerateScene() Scene {
	var spheres []Sphere

	// GROUND
	spheres = append(spheres, Sphere{
		center: Vec3{0, -1000.0, 1.0},
		radius: 1000.0,
		material: Material{
			materialType: MaterialTypeDiffuse,
			textureType:  TextureTypeCheckered,
			albedo:       Vec3{0.05, 0.05, 0.05},
			albedo2:      Vec3{0.95, 0.95, 0.95},
		},
	})

	for x := -11; x < 11; x++ {
		for z := -11; z < 11; z++ {
			materialRandom := rand.Float64()
			var material Material

			if materialRandom < 0.8 {
				material = Material{materialType: MaterialTypeDiffuse, textureType: TextureTypeSolid, albedo: GetRandomColor()}
			} else if materialRandom < 0.95 {
				material = Material{materialType: MaterialTypeMetal, textureType: TextureTypeSolid, albedo: GetRandomColor()}
			} else {
				material = Material{materialType: MaterialTypeDielectric, refractionIndex: 1.5}
			}

			spheres = append(spheres, Sphere{
				center:   Vec3{float64(x) + 0.9*rand.Float64(), 0.2, float64(z) + 0.9*rand.Float64()},
				radius:   0.2,
				material: material,
			})
		}
	}

	// BIG SPHERES
	spheres = append(spheres, Sphere{
		center: Vec3{0.0, 1.0, 0.0},
		radius: 1.0,
		material: Material{
			materialType:    MaterialTypeDielectric,
			refractionIndex: 1.5,
		},
	})

	spheres = append(spheres, Sphere{
		center: Vec3{-4.0, 1.0, 0.0},
		radius: 1.0,
		material: Material{
			materialType: MaterialTypeDiffuse,
			textureType:  TextureTypeSolid,
			albedo:       Vec3{0.6, 0.3, 0.1},
		},
	})

	spheres = append(spheres, Sphere{
		center: Vec3{4.0, 1.0, 0.0},
		radius: 1.0,
		material: Material{
			materialType: MaterialTypeMetal,
			textureType:  TextureTypeSolid,
			albedo:       Vec3{0.7, 0.6, 0.5},
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
