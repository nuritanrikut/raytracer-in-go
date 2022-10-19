package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%+.4f, %+.4f, %+.4f)", v.X, v.Y, v.Z)
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (lhs Vec3) Add(rhs Vec3) Vec3 {
	return Vec3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

func (lhs *Vec3) AddAssign(rhs Vec3) *Vec3 {
	lhs.X += rhs.X
	lhs.Y += rhs.Y
	lhs.Z += rhs.Z
	return lhs
}

func (lhs Vec3) Sub(rhs Vec3) Vec3 {
	return Vec3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func (lhs Vec3) Neg() Vec3 {
	return Vec3{-lhs.X, -lhs.Y, -lhs.Z}
}

func (lhs Vec3) Mul(rhs Vec3) Vec3 {
	return Vec3{lhs.X * rhs.X, lhs.Y * rhs.Y, lhs.Z * rhs.Z}
}

func (lhs Vec3) Times(rhs float64) Vec3 {
	return Vec3{lhs.X * rhs, lhs.Y * rhs, lhs.Z * rhs}
}

func (lhs Vec3) Div(rhs float64) Vec3 {
	return Vec3{lhs.X / rhs, lhs.Y / rhs, lhs.Z / rhs}
}

func (v Vec3) UnitVector() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func Dot(a Vec3, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Sub(n.Times(Dot(v, n) * 2.0))
}

func Refract(uv Vec3, n Vec3, etai_over_etat float64) Vec3 {
	var cos_theta = math.Min(Dot(uv.Neg(), n), 1.0)
	var r_out_perp = (uv.Add(n.Times(cos_theta))).Times(etai_over_etat)
	var r_out_parallel = n.Times(-math.Sqrt(math.Abs(-(1.0 - r_out_perp.LengthSquared()))))
	return r_out_perp.Add(r_out_parallel)
}
