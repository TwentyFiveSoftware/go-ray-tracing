package scene

import (
	"math"
	"math/rand"

	"go-ray-tracing/pkg/color"
	"go-ray-tracing/pkg/hitRecord"
	"go-ray-tracing/pkg/material"
	"go-ray-tracing/pkg/ray"
	"go-ray-tracing/pkg/sphere"
	"go-ray-tracing/pkg/texture"
	"go-ray-tracing/pkg/vec3"
)

type Scene struct {
	spheres []sphere.Sphere
}

func GenerateScene() Scene {
	var spheres []sphere.Sphere

	for x := -11; x < 11; x++ {
		for z := -11; z < 11; z++ {
			var sphereMaterial material.Material

			materialRandom := rand.Float64()
			if materialRandom < 0.8 {
				sphereMaterial = &material.DiffuseMaterial{Texture: &texture.SolidTexture{Albedo: color.GetRandomColor()}}
			} else if materialRandom < 0.95 {
				sphereMaterial = &material.MetalMaterial{Texture: &texture.SolidTexture{Albedo: color.GetRandomColor()}}
			} else {
				sphereMaterial = &material.DielectricMaterial{RefractionIndex: 1.5}
			}

			spheres = append(spheres, sphere.Sphere{
				Center:   vec3.Vec3{X: float64(x) + 0.9*rand.Float64(), Y: 0.2, Z: float64(z) + 0.9*rand.Float64()},
				Radius:   0.2,
				Material: sphereMaterial,
			})
		}
	}

	// GROUND
	spheres = append(spheres, sphere.Sphere{
		Center: vec3.Vec3{Y: -1000.0, Z: 1.0},
		Radius: 1000.0,
		Material: &material.DiffuseMaterial{
			Texture: &texture.CheckeredTexture{
				Albedo1: vec3.Vec3{X: 0.05, Y: 0.05, Z: 0.05},
				Albedo2: vec3.Vec3{X: 0.95, Y: 0.95, Z: 0.95},
			},
		},
	})

	// BIG SPHERES
	spheres = append(spheres, sphere.Sphere{
		Center:   vec3.Vec3{Y: 1.0},
		Radius:   1.0,
		Material: &material.DielectricMaterial{RefractionIndex: 1.5},
	})

	spheres = append(spheres, sphere.Sphere{
		Center:   vec3.Vec3{X: -4.0, Y: 1.0},
		Radius:   1.0,
		Material: &material.DiffuseMaterial{Texture: &texture.SolidTexture{Albedo: vec3.Vec3{X: 0.6, Y: 0.3, Z: 0.1}}},
	})

	spheres = append(spheres, sphere.Sphere{
		Center:   vec3.Vec3{X: 4.0, Y: 1.0},
		Radius:   1.0,
		Material: &material.MetalMaterial{Texture: &texture.SolidTexture{Albedo: vec3.Vec3{X: 0.7, Y: 0.6, Z: 0.5}}},
	})

	return Scene{spheres}
}

func (scene *Scene) RayHitScene(ray ray.Ray) hitRecord.HitRecord {
	currentRecord := hitRecord.HitRecord{T: math.Inf(1)}

	for _, sphereInScene := range scene.spheres {
		if hitRecord := sphereInScene.RayHitSphere(ray, 0.001, currentRecord.T); hitRecord.DoesHit {
			currentRecord = hitRecord
		}
	}

	return currentRecord
}
