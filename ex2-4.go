// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

var xA, yA, xB, yB, xC, yC float64 // clipping triangle
var epsilon float64
var bisects int

func main() {
	alpha := 72 * math.Pi / 180
	phi0 := 0.0
	grsys.InitGr("")
	epsilon = 1 / grsys.Density

	d := 0.1 * grsys.RMax
	xA = grsys.XCenter
	yA = grsys.YMax - d
	xB = grsys.XMin + d
	yB = grsys.YMin + d
	xC = grsys.XMax - d
	yC = grsys.YMin + d

	grsys.Move(xA, yA)
	grsys.Draw(xB, yB)
	grsys.Draw(xC, yC)
	grsys.Draw(xA, yA)

	for j := 1; j <= 20; j++ {
		r := float64(j) * d
		x2 := grsys.XCenter + r*math.Cos(phi0)
		y2 := grsys.YCenter + r*math.Sin(phi0)
		for i := 1; i <= 5; i++ {
			phi := phi0 + float64(i)*alpha
			x1 := x2
			y1 := y2
			x2 = grsys.XCenter + r*math.Cos(phi)
			y2 = grsys.YCenter + r*math.Sin(phi)
			clipDraw(x1, y1, x2, y2)
		}
	}
	grsys.EndGr()
	fmt.Println("Bisections:", bisects)
}

func clipCode(x, y float64) int {
	// Compute the barymetric coordinates of (x, y) with respect to the
	// clipping triangle's vertices.  If an ordinate for a vertex is -ve
	// then (x, y) lies outside the triangle, opposite the vertex.
	det := (yB-yC)*(xA-xC) + (xC-xB)*(yA-yC)
	lambdaA := ((yB-yC)*(x-xC) + (xC-xB)*(y-yC)) / det
	lambdaB := ((yC-yA)*(x-xC) + (xA-xC)*(y-yC)) / det
	lambdaC := 1 - lambdaA - lambdaB

	var c int
	if lambdaA < 0 {
		c |= 4
	}
	if lambdaB < 0 {
		c |= 2
	}
	if lambdaC < 0 {
		c |= 1
	}
	return c
}

func clipDraw(x1, y1, x2, y2 float64) {
	// Identical implementation to the exercise with a rectangular clipping
	// region; only the clipCode needs to change to account for difference
	// in shapes.
	if math.Abs(x1-x2) < epsilon && math.Abs(y1-y2) < epsilon {
		return
	}
	c1 := clipCode(x1, y1)
	c2 := clipCode(x2, y2)
	if (c1 & c2) != 0 {
		return
	}
	if (c1 | c2) == 0 {
		grsys.Move(x1, y1)
		grsys.Draw(x2, y2)
		return
	}
	bisects++
	xm := (x1 + x2) / 2
	ym := (y1 + y2) / 2
	clipDraw(x1, y1, xm, ym)
	clipDraw(xm, ym, x2, y2)
}
