package main

import "fmt"

type Result struct {
	U, R, B, G, W, C, Non int
}

type Results struct {
	U, R, B, G, W, C, Non float64
}

// Add adds values of r2 to r
func (r *Result) Add(r2 Result) {
	r.U += r2.U
	r.R += r2.R
	r.B += r2.B
	r.G += r2.G
	r.W += r2.W
	r.C += r2.C
	r.Non += r2.Non
}

// Returns the average of the Result
func (r Result) Average(n int) Results {
	return Results{
		U:   float64(r.U) / float64(n),
		R:   float64(r.R) / float64(n),
		B:   float64(r.B) / float64(n),
		G:   float64(r.G) / float64(n),
		W:   float64(r.W) / float64(n),
		C:   float64(r.C) / float64(n),
		Non: float64(r.Non) / float64(n),
	}
}

func (r Results) String() string {
	return fmt.Sprintf("U: %.2f\nR: %.2f\nB: %.2f\nG: %.2f\nW: %.2f\nC: %.2f\nNon: %.2f\n", r.U, r.R, r.B, r.G, r.W, r.C, r.Non)
}
