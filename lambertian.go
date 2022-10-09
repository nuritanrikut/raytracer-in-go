package main

type Lambertian struct {
	Albedo Vec3
}

func (m Lambertian) Diffuse() Vec3 {
	return m.Albedo
}

func (m Lambertian) Scatter(rng *RandomNumberGenerator, r_in Ray, rec HitRecord) (bool, *Vec3, *Ray) {
	scatter_direction := rec.Normal.Add(rng.RandomUnitVec3())
	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}

	scattered := Ray{rec.P, scatter_direction}
	attenuation := m.Albedo

	return true, &attenuation, &scattered
}
