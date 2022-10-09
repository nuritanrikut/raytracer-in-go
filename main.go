package main

import (
	"fmt"
	"math"
	"os"
)

func simple_scene() HittableList {
	world := HittableList{}

	ground_material := Lambertian{Albedo: Vec3{0.5, 0.5, 0.5}}
	world.Add(&Sphere{Center: Vec3{0, -1000, 0}, Radius: 1000, Material: ground_material})

	material1 := Dielectric{Ir: 1.5}
	world.Add(&Sphere{Center: Vec3{0, 1, 0}, Radius: 1, Material: material1})

	material2 := Lambertian{Albedo: Vec3{0.4, 0.2, 0.1}}
	world.Add(&Sphere{Center: Vec3{-4, 1, 0}, Radius: 1, Material: material2})

	material3 := Metal{Albedo: Vec3{0.7, 0.6, 0.5}, Fuzz: 0.0}
	world.Add(&Sphere{Center: Vec3{4, 1, 0}, Radius: 1, Material: material3})

	return world
}

func random_scene(rng *RandomNumberGenerator) HittableList {
	world := HittableList{}

	ground_material := Lambertian{Albedo: Vec3{0.5, 0.5, 0.5}}
	world.Add(&Sphere{Center: Vec3{0, -1000, 0}, Radius: 1000, Material: ground_material})

	world_center := Vec3{4, 0.2, 0}
	radius := 0.2

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			center := Vec3{
				float64(a) + 0.9*rng.RandomDouble(),
				0.2,
				float64(b) + 0.9*rng.RandomDouble(),
			}

			if center.Sub(world_center).Length() > 0.9 {
				var material AbsMaterial
				choose_mat := rng.RandomDouble()
				if choose_mat < 0.8 {
					// diffuse
					albedo := rng.RandomVec3()
					material = Lambertian{Albedo: albedo}
				} else if choose_mat < 0.95 {
					// metal
					albedo := rng.RandomVec3Range(0.5, 1)
					fuzz := rng.RandomDoubleRange(0, 0.5)
					material = Metal{Albedo: albedo, Fuzz: fuzz}
				} else {
					// glass
					material = Dielectric{Ir: 1.5}
				}

				world.Add(&Sphere{center, radius, material})
			}
		}
	}

	material1 := Dielectric{Ir: 1.5}
	world.Add(&Sphere{Center: Vec3{0, 1, 0}, Radius: 1, Material: material1})

	material2 := Lambertian{Albedo: Vec3{0.4, 0.2, 0.1}}
	world.Add(&Sphere{Center: Vec3{-4, 1, 0}, Radius: 1, Material: material2})

	material3 := Metal{Albedo: Vec3{0.7, 0.6, 0.5}, Fuzz: 0.0}
	world.Add(&Sphere{Center: Vec3{4, 1, 0}, Radius: 1, Material: material3})

	return world
}

func ray_color(i int, j int, r Ray, world HittableList, depth int, rng *RandomNumberGenerator) Vec3 {
	if depth <= 0 {
		return Vec3{}
	}
	const tMin = 0.001
	const tMax = math.MaxFloat64 // technically not infinity but good enough

	hit_something, rec := world.Hit(r, tMin, tMax)
	if hit_something {
		to_scatter, attenuation, scattered := rec.Material.Scatter(rng, r, *rec)
		if to_scatter {
			recursed_color := ray_color(i, j, *scattered, world, depth-1, rng)
			// fmt.Fprintf(os.Stderr,
			// 	"> Scatter %v %v dir=%v color= %v, %v, %v\n",
			// 	i,
			// 	j,
			// 	r.Direction,
			// 	attenuation,
			// 	recursed_color,
			// 	attenuation.Mul(recursed_color),
			// )
			return attenuation.Mul(recursed_color)
		}

		// fmt.Fprintf(os.Stderr,
		// 	"> Diffuse %v %v = %v\n",
		// 	i,
		// 	j,
		// 	rec.Material.Diffuse(),
		// )
		return rec.Material.Diffuse()
	}

	unit_direction := r.Direction.UnitVector()
	t := 0.5 * (unit_direction.Y + 1.0)
	white := Vec3{1, 1, 1}
	blue := Vec3{0.5, 0.7, 1}
	sky := white.Times(1.0 - t).Add(blue.Times(t))
	// fmt.Fprintf(os.Stderr,
	// 	"> Sky %v %v = %v\n",
	// 	i,
	// 	j,
	// 	sky,
	// )
	return sky
}

func write_color(pixel_color Vec3, samples_per_pixel int) {
	scale := 1.0 / float64(samples_per_pixel)
	r := math.Sqrt(scale * pixel_color.X)
	g := math.Sqrt(scale * pixel_color.Y)
	b := math.Sqrt(scale * pixel_color.Z)

	ri := int(256.0 * Clamp(r, 0.0, 0.999))
	gi := int(256.0 * Clamp(g, 0.0, 0.999))
	bi := int(256.0 * Clamp(b, 0.0, 0.999))

	fmt.Printf("%v %v %v\n", ri, gi, bi)
}

func main() {
	aspect_ratio := 16.0 / 10.0
	image_width := 1920
	image_height := int(float64(image_width) / aspect_ratio)
	samples_per_pixel_x := 16
	samples_per_pixel_y := 16
	max_depth := 50

	fmt.Fprintf(os.Stderr,
		"Rendering %vx%v image with %vx%v samples per pixel\n",
		image_width, image_height, samples_per_pixel_x, samples_per_pixel_y,
	)

	rng := CreateRNG()
	rng.RandomDouble()

	var world HittableList
	world_name := "simple"
	if len(os.Args) > 1 {
		world_name = os.Args[1]
	}

	if world_name == "simple" {
		fmt.Fprintf(os.Stderr, "Loading simple scene\n")
		world = simple_scene()
	} else {
		fmt.Fprintf(os.Stderr, "Loading random scene\n")
		world = random_scene(rng)
	}

	lookfrom := Vec3{13, 2, 3}
	lookat := Vec3{0, 0, 0}
	vup := Vec3{0, 1, 0}
	vfov := 20.0 // vertical field-of-view in degrees
	focus_dist := 10.0
	aperture := 0.1
	sample_count := samples_per_pixel_x * samples_per_pixel_y

	var cam Camera = Create(lookfrom, lookat, vup, vfov, aspect_ratio, aperture, focus_dist)

	fmt.Printf(
		"P3\n%v %v\n255\n", image_width, image_height,
	)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %v\n", j)
		for i := 0; i < image_width; i++ {
			pixel_color := Vec3{0, 0, 0}
			for s := 0; s < sample_count; s++ {
				y := float64(s/samples_per_pixel_y)/float64(samples_per_pixel_y) - 0.5
				x := float64(s%samples_per_pixel_y)/float64(samples_per_pixel_x) - 0.5
				u := (float64(i) + x) / float64(image_width-1)
				v := (float64(j) + y) / float64(image_height-1)
				r := cam.GetRay(rng, u, v)
				pixel_color.AddAssign(ray_color(i, j, r, world, max_depth, rng))
			}

			write_color(pixel_color, sample_count)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
