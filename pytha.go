// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
// Pythagoras' tree.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"os"
)

func Pythagoras(A, B grsys.Vec, n int) {
	if n <= 0 {
		return
	}
	C := B.Add(grsys.Vec{X: A.Y - B.Y, Y: B.X - A.X})
	D := A.Add(C.Sub(B))
	E := D.Add(C.Sub(A).Mul(0.5))
	A.Move()
	B.Draw()
	C.Draw()
	D.Draw()
	A.Draw()
	Pythagoras(D, E, n-1)
	Pythagoras(E, C, n-1)
}

func main() {
	fmt.Print("Enter xA, yA, xB, yB and recursion depth n: ")
	var A, B grsys.Vec
	var n int
	if _, err := fmt.Scanln(&A.X, &A.Y, &B.X, &B.Y, &n); err != nil {
		grsys.Error(err)
	}
	// Optional argument for output to HP-GL file
	if len(os.Args) > 1 {
		grsys.InitGr(os.Args[1])
	} else {
		grsys.InitGr("")
	}
	Pythagoras(A, B, n)
	grsys.EndGr()
}
