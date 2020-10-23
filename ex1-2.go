package main

import (
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	grsys.Initgr("")
	r := grsys.RMax
	delta := r / 20
	for i := 0; i < 20; i++ {
		var lastx, lasty float32
		for j := 0; j < 60; j++ {
			theta := 2 * math.Pi * float64(j) / 60
			x := grsys.XCenter + r*float32(math.Cos(theta))
			y := grsys.YCenter + r*float32(math.Sin(theta))
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
	grsys.Endgr()
}
