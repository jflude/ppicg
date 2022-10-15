// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

var vtx [5]grsys.Vec

func drawStar(center grsys.Vec, r float64, n int) {
	if n == 0 {
		return
	}
	var v [5]grsys.Vec
	for i := range v {
		v[i] = center.Add(vtx[i].Mul(r))
	}
	v[0].Move()
	v[2].Draw()
	v[4].Draw()
	v[1].Draw()
	v[3].Draw()
	v[0].Draw()
	for i := range v {
		v[i] = center.Add(vtx[i].Mul(2 * r))
	}
	n--
	r /= 3
	for i := range v {
		drawStar(v[i], r, n)
	}
}

func main() {
	fmt.Print("Enter the maximum depth of recursion: ")
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	for i := range vtx {
		theta := 2 * math.Pi * float64(i) / 5
		vtx[i].X = math.Cos(theta)
		vtx[i].Y = math.Sin(theta)
	}
	c := grsys.Vec{X: grsys.XCenter, Y: grsys.YCenter}
	drawStar(c, grsys.RMax/4, n)
	grsys.EndGr()
}
