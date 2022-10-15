// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func drawTriangles(n int, a, b, c grsys.Vec) {
	if n == 0 {
		return
	}
	a.Move()
	b.Draw()
	c.Draw()
	a.Draw()
	ab2 := a.Add(b.Sub(a).Mul(0.5))
	ac2 := a.Add(c.Sub(a).Mul(0.5))
	bc2 := b.Add(c.Sub(b).Mul(0.5))
	n--
	drawTriangles(n, a, ab2, ac2)
	drawTriangles(n, ab2, b, bc2)
	drawTriangles(n, ac2, bc2, c)
}

func main() {
	fmt.Printf("Enter the vertices of the triangle: ")
	var a, b, c grsys.Vec
	_, err := fmt.Scanln(&a.X, &a.Y, &b.X, &b.Y, &c.X, &c.Y)
	if err != nil {
		grsys.Error(err)
	}
	fmt.Print("Enter the maximum depth of recursion: ")
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	drawTriangles(n, a, b, c)
	grsys.EndGr()
}
