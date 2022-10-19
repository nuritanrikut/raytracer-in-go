package main

import (
	"math/rand"
)

type RandomNumberGenerator struct {
	seed  rand.Source
	gen   rand.Rand
	state uint64
	div   uint64
	mod   uint64
}

func CreateRNG() *RandomNumberGenerator {
	p := new(RandomNumberGenerator)
	p.seed = rand.NewSource(1)
	p.gen = *rand.New(p.seed)
	p.state = 675248
	p.div = 1000
	p.mod = 1000000
	return p
}

func (rng *RandomNumberGenerator) Next() uint64 {
	rng.state = rng.state * rng.state / rng.div % rng.mod
	return rng.state
}

func (rng *RandomNumberGenerator) RandomDouble() float64 {
	// use this version if you want actual PRNG
	// return rng.gen.Float64()

	// use this version if you want repeatable sequence of numbers
	t1 := rng.Next()
	t2 := rng.Next()
	return float64(t1*rng.mod+t2) / float64(rng.mod) / float64(rng.mod)
}

func (rng *RandomNumberGenerator) RandomDoubleRange(min float64, max float64) float64 {
	return min + (max-min)*rng.RandomDouble()
}

func (rng *RandomNumberGenerator) RandomVec3() Vec3 {
	return Vec3{
		rng.RandomDouble(),
		rng.RandomDouble(),
		rng.RandomDouble(),
	}
}

func (rng *RandomNumberGenerator) RandomVec3Range(min float64, max float64) Vec3 {
	return Vec3{
		rng.RandomDoubleRange(min, max),
		rng.RandomDoubleRange(min, max),
		rng.RandomDoubleRange(min, max),
	}
}

func (rng *RandomNumberGenerator) RandomVec3InUnitSphere() Vec3 {
	for {
		var p = rng.RandomVec3Range(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func (rng *RandomNumberGenerator) RandomUnitVec3() Vec3 {
	return rng.RandomVec3InUnitSphere().UnitVector()
}

func (rng *RandomNumberGenerator) RandomVec3InUnitDisk() Vec3 {
	for {
		var p = Vec3{
			rng.RandomDoubleRange(-1, 1),
			rng.RandomDoubleRange(-1, 1),
			0,
		}
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}
