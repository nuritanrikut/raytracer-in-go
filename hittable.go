package main

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord)
}
