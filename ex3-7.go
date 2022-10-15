// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

const EPS = 1e-6

func main() {
	fmt.Print("Enter the coordinates of the triangle vertices: ")
	var a, b, c grsys.Vec
	_, err := fmt.Scanln(&a.X, &a.Y, &b.X, &b.Y, &c.X, &c.Y)
	if err != nil {
		grsys.Error(err)
	}
	ba := b.Sub(a)
	bc := b.Sub(c)
	orient := ba.Y*bc.X - ba.X*bc.Y
	if math.Abs(orient) < EPS {
		grsys.Error(errors.New("vertices are collinear"))
	}
	if orient < 0 {
		a, c = c, a
		ba = b.Sub(a)
		bc = b.Sub(c)
	}

	fmt.Print("Enter the radius of the fillet: ")
	var r float64
	if _, err := fmt.Scanln(&r); err != nil {
		grsys.Error(err)
	}

	ac := a.Sub(c)
	ab2 := ba.X*ba.X + ba.Y*ba.Y
	bc2 := bc.X*bc.X + bc.Y*bc.Y
	ca2 := ac.X*ac.X + ac.Y*ac.Y
	abLen := math.Sqrt(ab2)
	bcLen := math.Sqrt(bc2)
	cosB := (ab2 + bc2 - ca2) / (2 * abLen * bcLen)
	bAngle2 := math.Acos(cosB) / 2

	bfLen := -r / math.Tan(bAngle2)
	bf := ba.Mul(bfLen / abLen)
	bg := bc.Mul(bfLen / bcLen) // bgLen == bfLen
	f := b.Add(bf)
	g := b.Add(bg)
	center := f.Add(grsys.Vec{X: -bf.Y, Y: bf.X}.Mul(r / bfLen))

	p := f.Sub(center)
	q := g.Sub(center)
	fTheta := math.Atan2(p.Y, p.X)
	gTheta := math.Atan2(q.Y, q.X)

	grsys.InitGr("")
	b.Sub(grsys.Vec{X: -0.05, Y: -0.05}).Move()
	b.Sub(grsys.Vec{X: 0.05, Y: 0.05}).Draw()
	b.Sub(grsys.Vec{X: -0.05, Y: 0.05}).Move()
	b.Sub(grsys.Vec{X: 0.05, Y: -0.05}).Draw()
	a.Move()
	if gTheta < fTheta {
		gTheta += 2 * math.Pi
	}
	for th := fTheta; th < gTheta; th += 0.05 {
		p := grsys.Vec{X: r * math.Cos(th), Y: r * math.Sin(th)}
		center.Add(p).Draw()
	}
	c.Draw()
	grsys.EndGr()
}
