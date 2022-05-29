package main

import "math"

const MaterialTypeDiffuse = 0
const MaterialTypeMetal = 1
const MaterialTypeDielectric = 2

const TextureTypeSolid = 0
const TextureTypeCheckered = 1

type Material struct {
	materialType    uint8
	textureType     uint8
	albedo          Vec3
	albedo2         Vec3
	refractionIndex float64
}

func (material *Material) Scatter(ray Ray, hitRecord HitRecord) ScatterRecord {
	switch material.materialType {
	case MaterialTypeDiffuse:
		return scatterDiffuse(hitRecord)
	case MaterialTypeMetal:
		return scatterMetal(ray, hitRecord)
	case MaterialTypeDielectric:
		return scatterDielectric(ray, hitRecord)
	}

	return ScatterRecord{}
}

func scatterDiffuse(hitRecord HitRecord) ScatterRecord {
	scatterDirection := hitRecord.normal.Add(RandomUnitVector())

	if scatterDirection.IsNearZero() {
		scatterDirection = hitRecord.normal
	}

	return ScatterRecord{
		doesScatter: true,
		attenuation: hitRecord.material.getColor(hitRecord.point),
		scatteredRay: Ray{
			Origin:    hitRecord.point,
			Direction: scatterDirection,
		},
	}
}

func scatterMetal(ray Ray, hitRecord HitRecord) ScatterRecord {
	scatterDirection := ray.Direction.Normalized().Reflect(hitRecord.normal)

	return ScatterRecord{
		doesScatter: hitRecord.normal.Dot(scatterDirection) > 0.0,
		attenuation: hitRecord.material.getColor(hitRecord.point),
		scatteredRay: Ray{
			Origin:    hitRecord.point,
			Direction: scatterDirection,
		},
	}
}

func scatterDielectric(ray Ray, hitRecord HitRecord) ScatterRecord {
	refractionRatio := hitRecord.material.refractionIndex
	if hitRecord.isFrontFace {
		refractionRatio = 1.0 / refractionRatio
	}

	scatterDirection := ray.Direction.Normalized().Refract(hitRecord.normal, refractionRatio)

	return ScatterRecord{
		doesScatter: true,
		attenuation: Vec3{1.0, 1.0, 1.0},
		scatteredRay: Ray{
			Origin:    hitRecord.point,
			Direction: scatterDirection,
		},
	}
}

func (material *Material) getColor(point Vec3) Vec3 {
	switch material.textureType {
	case TextureTypeSolid:
		return material.albedo
	case TextureTypeCheckered:
		size := 6.0
		sin := math.Sin(size*point.X) * math.Sin(size*point.Y) * math.Sin(size*point.Z)
		if sin < 0.0 {
			return material.albedo
		}
		return material.albedo2
	}

	return material.albedo
}
