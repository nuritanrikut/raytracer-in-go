package main

import (
	"math"
)

type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
	U               Vec3
	V               Vec3
	LensRadius      float64
}

func Create(lookfrom Vec3,
	lookat Vec3,
	vup Vec3,
	vfov float64,
	aspect_ratio float64,
	aperture float64,
	focus_dist float64,
) Camera {
	theta := DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Sub(lookat).UnitVector()
	u := Cross(vup, w).UnitVector()
	v := Cross(w, u)

	origin := lookfrom
	horizontal := u.Times(focus_dist * viewport_width)
	vertical := v.Times(focus_dist * viewport_height)
	lower_left_corner := origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2)).Sub(w.Times(focus_dist))

	lens_radius := aperture / 2.0

	return Camera{
		origin,
		lower_left_corner,
		horizontal,
		vertical,
		u,
		v,
		lens_radius,
	}
}

func (c Camera) GetRay(rng *RandomNumberGenerator, s float64, t float64) Ray {
	rd := rng.RandomVec3InUnitDisk().Times(c.LensRadius)
	offset := c.U.Times(rd.X).Add(c.V.Times(rd.Y))
	return Ray{
		c.Origin.Add(offset),
		c.LowerLeftCorner.Add(c.Horizontal.Times(s)).Add(c.Vertical.Times(t)).Sub(c.Origin).Sub(offset),
	}
}
