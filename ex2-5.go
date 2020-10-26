package main

import (
	"github.com/jflude/ppicg/grsys"
	"math/rand"
	"time"
)

const (
	PointCount = 30
	Intervals  = 30
)

func main() {
	grsys.InitGr("")
	grsys.InitWindow()
	rand.Seed(time.Now().UnixNano())
	var x, y []float64
	for i := 0; i < PointCount; i++ {
		xx := grsys.XCenter + grsys.RMax*(2*rand.Float64()-1)
		yy := grsys.YCenter + grsys.RMax*(2*rand.Float64()-1)
		x = append(x, xx)
		y = append(y, yy)
		grsys.UpdateWindowBoundaries(xx, yy)
	}

	grsys.ViewportBoundaries(grsys.XMin, grsys.XMax,
		grsys.YMin, grsys.YMax, 0.9)
	eps := 0.04
	for i := range x {
		X := grsys.XViewport(x[i])
		Y := grsys.YViewport(y[i])
		grsys.Move(X-eps, Y-eps)
		grsys.Draw(X+eps, Y+eps)
		grsys.Move(X+eps, Y-eps)
		grsys.Draw(X-eps, Y+eps)
	}

	first := true
	for i := 1; i < len(x)-2; i++ {
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

		for j := 0; j <= Intervals; j++ {
			t := float64(j) / Intervals
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
