package main

import "math"

type HittableList struct {
	Objects []Hittable
}

func (l *HittableList) Add(object Hittable) {
	l.Objects = append(l.Objects, object)
}

func (l *HittableList) Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	var result *HitRecord
	hit_anything := false
	closest_so_far := math.MaxFloat64

	for _, object := range l.Objects {
		recursed_hit, recursed_rec := object.Hit(r, tMin, closest_so_far)
		if recursed_hit {
			hit_anything = true
			closest_so_far = recursed_rec.T
			result = recursed_rec
		}
	}

	if hit_anything {
		return true, result
	}

	return false, nil
}
