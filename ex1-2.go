package main

import (
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	grsys.InitGr("")
	r := grsys.RMax
	delta := r / 20
	for i := 0; i < 20; i++ {
		var lastx, lasty float64
		for j := 0; j < 60; j++ {
			theta := 2 * math.Pi * float64(j) / 60
			x := grsys.XCenter + r*math.Cos(theta)
			y := grsys.YCenter + r*math.Sin(theta)
			if j == 0 {
				grsys.Move(x, y)
				lastx = x
				lasty = y
			} else {
				grsys.Draw(x, y)
			}
		}
		grsys.Draw(lastx, lasty)
		r -= delta
	}
	grsys.EndGr()
}
