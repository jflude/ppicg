// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import "github.com/jflude/ppicg/grsys"

func drawEllipse(xC, yC, rHorz, rVert int) {
	a2 := rHorz * rHorz
	b2 := rVert * rVert

	grsys.PutPix(xC, yC+rVert) // Top
	grsys.PutPix(xC, yC-rVert) // Bottom
	x := 0
	y := rVert
	u := b2
	v := a2 * (2*rVert - 1)
	E := 0
	for u < v {
		x++
		E += u
		u += 2 * b2
		if v < 2*E {
			y--
			E -= v
			v -= 2 * a2
		}
		grsys.PutPix(xC+x, yC+y) // Octant 2
		grsys.PutPix(xC-x, yC+y) // Octant 3
		grsys.PutPix(xC+x, yC-y) // Octant 7
		grsys.PutPix(xC-x, yC-y) // Octant 6
	}

	grsys.PutPix(xC+rHorz, yC) // Right
	grsys.PutPix(xC-rHorz, yC) // Left
	x = rHorz
	y = 0
	u = a2
	v = b2 * (2*rHorz - 1)
	E = 0
	for u < v {
		y++
		E += u
		u += 2 * a2
		if v < 2*E {
			x--
			E -= v
			v -= 2 * b2
		}
		grsys.PutPix(xC+x, yC+y) // Octant 1
		grsys.PutPix(xC-x, yC+y) // Octant 4
		grsys.PutPix(xC+x, yC-y) // Octant 8
		grsys.PutPix(xC-x, yC-y) // Octant 5
	}
}

func main() {
	grsys.InitGr("")
	drawEllipse(grsys.ImageWidth/2, grsys.ImageHeight/2,
		grsys.ImageWidth/2, grsys.ImageHeight/2)
	grsys.EndGr()
}
