package main

import "math"

type Dielectric struct {
	Ir float64
}

func reflectance(cosine float64, ref_idx float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1.0 - ref_idx) / (1.0 + ref_idx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5.0)
}

func (m Dielectric) Diffuse() Vec3 {
	return Vec3{1, 1, 1}
}

func (m Dielectric) Scatter(rng *RandomNumberGenerator, r_in Ray, rec HitRecord) (bool, *Vec3, *Ray) {
	attenuation := Vec3{1, 1, 1}
	var refraction_ratio float64
	if rec.FrontFace {
		refraction_ratio = 1 / m.Ir
	} else {
		refraction_ratio = m.Ir
	}

	unit_direction := r_in.Direction.UnitVector()
	d := Dot(unit_direction.Neg(), rec.Normal)
	cos_theta := math.Min(d, 1.0)
	sin_theta := math.Sqrt(1 - cos_theta*cos_theta)

	cannot_refract := refraction_ratio*sin_theta > 1.0
	t := rng.RandomDouble()
	refl := reflectance(cos_theta, refraction_ratio)
	should_reflect := refl > t

	var direction Vec3
	if cannot_refract || should_reflect {
		direction = Reflect(unit_direction, rec.Normal)
	} else {
		direction = Refract(unit_direction, rec.Normal, refraction_ratio)
	}

	scattered := Ray{rec.P, direction}

	return true, &attenuation, &scattered
}
