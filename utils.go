package main

import "math"

func Clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
