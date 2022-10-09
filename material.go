package main

type AbsMaterial interface {
	Diffuse() Vec3
	Scatter(rng *RandomNumberGenerator, r_in Ray, rec HitRecord) (bool, *Vec3, *Ray)
}
