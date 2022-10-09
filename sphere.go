package main

import "math"

type Sphere struct {
	Center   Vec3
	Radius   float64
	Material AbsMaterial
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	half_b := Dot(oc, r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false, nil
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	t := (-half_b - sqrtd) / a
	if t < tMin || t > tMax {
		t = (-half_b + sqrtd) / a
		if t < tMin || t > tMax {
			return false, nil
		}
	}

	var rec *HitRecord = new(HitRecord)

	rec.T = t
	rec.P = r.At(t)
	outward_normal := rec.P.Sub(s.Center).Div(s.Radius)
	rec.SetFaceNormal(r, outward_normal)
	rec.Material = s.Material

	return true, rec
}
