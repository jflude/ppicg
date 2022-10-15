// Demonstration of the Cohen-Sutherland line-clipping algorithm.
package main

import (
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	alpha := 72 * math.Pi / 180
	phi0 := 0.0
	grsys.InitGr("")

	d := 0.1 * grsys.RMax
	xmin := grsys.XMin + d
	xmax := grsys.XMax - d
	ymin := grsys.YMin + d
	ymax := grsys.YMax - d

	// The window is now drawn:
	grsys.Move(xmin, ymin)
	grsys.Draw(xmax, ymin)
	grsys.Draw(xmax, ymax)
	grsys.Draw(xmin, ymax)
	grsys.Draw(xmin, ymin)
	grsys.SetClipBoundaries(xmin, xmax, ymin, ymax)

	// As far as permitted by the boundaries of the window,
	// 20 concentric regular pentagons are drawn:
	for j := 1; j <= 20; j++ {
		r := float64(j) * d
		x2 := grsys.XCenter + r*math.Cos(phi0)
		y2 := grsys.YCenter + r*math.Sin(phi0)
		for i := 1; i <= 5; i++ {
			phi := phi0 + float64(i)*alpha
			x1 := x2
			y1 := y2
			x2 = grsys.XCenter + r*math.Cos(phi)
			y2 = grsys.YCenter + r*math.Sin(phi)
			grsys.ClipDraw(x1, y1, x2, y2)
		}
	}
	grsys.EndGr()
}
