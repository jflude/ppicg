// Curve fitting using B splines
package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"io"
	"os"
)

const (
	// At most MAX points and N intervals between any two neighbor points.
	MAX = 1000
	N   = 30
)

func main() {
	f, err := os.Open("curvefit.txt")
	if err != nil {
		grsys.Error(err)
	}
	var n int
	if _, err := fmt.Fscanln(f, &n); err != nil {
		grsys.Error(err)
	}
	if n < 3 || n+1 >= MAX {
		grsys.Error(errors.New("First number read incorrectly"))
	}

	grsys.InitWindow()
	var x, y []float64
	for i := 0; i <= n; i++ {
		var xx, yy float64
		if _, err := fmt.Fscanln(f, &xx, &yy); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			grsys.Error(err)
		}
		x = append(x, xx)
		y = append(y, yy)
		grsys.UpdateWindowBoundaries(xx, yy)
	}

	grsys.InitGr("")
	grsys.ViewportBoundaries(grsys.XMin, grsys.XMax,
		grsys.YMin, grsys.YMax, 0.9)
	// Mark the given points:
	eps := 0.04
	for i := 0; i <= n; i++ {
		X := grsys.XViewport(x[i])
		Y := grsys.YViewport(y[i])
		grsys.Move(X-eps, Y-eps)
		grsys.Draw(X+eps, Y+eps)
		grsys.Move(X+eps, Y-eps)
		grsys.Draw(X-eps, Y+eps)
	}
	first := true
	for i := 1; i < n-1; i++ {
		xA := x[i-1]
		xB := x[i]
		xC := x[i+1]
		xD := x[i+2]

		yA := y[i-1]
		yB := y[i]
		yC := y[i+1]
		yD := y[i+2]

		a3 := (-xA + 3*(xB-xC) + xD) / 6
		a2 := (xA - 2*xB + xC) / 2
		a1 := (xC - xA) / 2
		a0 := (xA + 4*xB + xC) / 6

		b3 := (-yA + 3*(yB-yC) + yD) / 6
		b2 := (yA - 2*yB + yC) / 2
		b1 := (yC - yA) / 2
		b0 := (yA + 4*yB + yC) / 6

		for j := 0; j <= N; j++ {
			t := float64(j) / N
			X := grsys.XViewport(((a3*t+a2)*t+a1)*t + a0)
			Y := grsys.YViewport(((b3*t+b2)*t+b1)*t + b0)
			if first {
				first = false
				grsys.Move(X, Y)
			} else {
				grsys.Draw(X, Y)
			}
		}
	}
	grsys.EndGr()
}
