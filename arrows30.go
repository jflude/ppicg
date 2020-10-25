// This program draws 30 arrows, flying counter-clockwise about the center
// of the screen.
package main

import (
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	var p [4]grsys.Vec = [4]grsys.Vec{{0, -7}, {0, 7}, {-2, 0}, {2, 0}}
	phi := math.Pi / 15
	cosphi := math.Cos(phi)
	sinphi := math.Sin(phi)
	grsys.InitGr("")

	center := grsys.Vec{grsys.XCenter, grsys.YCenter}
	start := center.Add(grsys.Vec{0.9 * grsys.RMax, 0})
	for j := 0; j < 4; j++ {
		p[j] = p[j].Mul(0.01 * grsys.RMax).Add(start)
	}

	for i := 0; i < 30; i++ {
		for j := 0; j < 4; j++ {
			p[j] = grsys.Rotate(p[j], center, cosphi, sinphi)
		}
		p[0].Move()
		p[1].Draw()
		p[2].Draw()
		p[3].Draw()
		p[1].Draw()
	}

	grsys.EndGr()
}
