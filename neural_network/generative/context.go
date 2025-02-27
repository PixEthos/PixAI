// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// context.go
package generative

import (
	"log"

	matrix "pixai/data/matrix"
)

// primary context
func (g *Generative) PrimaryContext(input string) (float64, float64) {
	gpe := context.GPEActivator(input)
	stop := context.STOPWORDActivator(input)

	return gpe, stop
}

// primary activator
func (g *Generative) GRU_primary(input matrix.Matrix32, value string) (float64, string) {
	// primary
	gpe, stop := g.PrimaryContext(value)
	GRU, gru_pri := neurons.Gru_processed(gpe, stop, input, value)
	primary := layer.GRU_sigmoid(gru_pri, "float64", value)
	val := g.GRU_decode(GRU, value)
	log.Println("GRU: ", val)
	log.Println("ASCII: ", GRU)

	return primary, val
}

// primary caller
func (g *Generative) Primary(input matrix.Matrix32, value string) {
	_, char := g.GRU_primary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}

// secondary
func (g *Generative) SecondaryContext(input string) float64 {
	nouns := context.NOUNActivator(input)

	return nouns
}

func (g *Generative) GRU_secondary(input matrix.Matrix32, value string) (float64, string) {
	// secondary
	nouns := g.SecondaryContext(value)
	GRU_2, gru_sec := neurons.Gru_processed_secondary(nouns, input, value)
	secondary := layer.GRU_sigmoid(gru_sec, "float64", value)
	val1 := g.GRU_decode(GRU_2, value)
	log.Println("GRU_2: ", val1)
	log.Println("ASCII: ", GRU_2)

	return secondary, val1
}

func (g *Generative) Secondary(input matrix.Matrix32, value string) {
	_, char := g.GRU_secondary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}

// trinary
func (g *Generative) TrinaryContext(input string) float64 {
	verbs := context.VERBActivator(input)

	return verbs
}

func (g *Generative) GRU_trinary(input matrix.Matrix32, value string) (float64, string) {
	// trinary
	verbs := g.TrinaryContext(value)
	GRU_3, gru_tri := neurons.Gru_processed_trinary(verbs, input, value)
	trinary := layer.GRU_sigmoid(gru_tri, "float64", value)
	val2 := g.GRU_decode(GRU_3, value)
	log.Println("GRU_3: ", val2)
	log.Println("ASCII: ", GRU_3)

	return trinary, val2
}

func (g *Generative) Trinary(input matrix.Matrix32, value string) {
	_, char := g.GRU_trinary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}
