package main

type HitRecord struct {
	P         Vec3
	Normal    Vec3
	Material  AbsMaterial
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r Ray, outward_normal Vec3) {
	h.FrontFace = Dot(r.Direction, outward_normal) < 0
	if h.FrontFace {
		h.Normal = outward_normal
	} else {
		h.Normal = outward_normal.Neg()
	}
}
