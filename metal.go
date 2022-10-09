package main

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Diffuse() Vec3 {
	return m.Albedo
}

func (m Metal) Scatter(rng *RandomNumberGenerator, r_in Ray, rec HitRecord) (bool, *Vec3, *Ray) {
	reflected := Reflect(r_in.Direction.UnitVector(), rec.Normal)
	scattered := Ray{rec.P, reflected.Add(rng.RandomVec3InUnitSphere().Times(m.Fuzz))}
	attenuation := m.Albedo

	return Dot(scattered.Direction, rec.Normal) > 0, &attenuation, &scattered
}
