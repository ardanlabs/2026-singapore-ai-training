// Example 01 — Vectors
package main

import (
	"fmt"
)

type data struct {
	Name      string
	Authority float64 // These fields are called features.
	Animal    float64
	Human     float64
	Rich      float64
	Gender    float64
}

// Vector can convert the specified data into a vector.
func (d data) Vector() []float64 {
	return []float64{
		d.Authority,
		d.Animal,
		d.Human,
		d.Rich,
		d.Gender,
	}
}

// String pretty prints an embedding to a vector representation.
func (d data) String() string {
	return fmt.Sprintf("%f", d.Vector())
}
