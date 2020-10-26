package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

var xmin, ymin, xmax, ymax, epsilon float64
var bisects int

func main() {
	alpha := 72 * math.Pi / 180
	phi0 := 0.0
	grsys.InitGr("")
	epsilon = 1 / grsys.Density

	d := 0.1 * grsys.RMax
	xmin = grsys.XMin + d
	xmax = grsys.XMax - d
	ymin = grsys.YMin + d
	ymax = grsys.YMax - d

	grsys.Move(xmin, ymin)
	grsys.Draw(xmax, ymin)
	grsys.Draw(xmax, ymax)
	grsys.Draw(xmin, ymax)
	grsys.Draw(xmin, ymin)

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
			clipDraw(x1, y1, x2, y2)
		}
	}
	grsys.EndGr()
	fmt.Println("Bisections:", bisects)
}

func clipCode(x, y float64) int {
	var c int
	if x < xmin {
		c |= 8
	} else if x > xmax {
		c |= 4
	}
	if y < ymin {
		c |= 2
	} else if y > ymax {
		c |= 1
	}
	return c
}

func clipDraw(x1, y1, x2, y2 float64) {
	if math.Abs(x1-x2) < epsilon && math.Abs(y1-y2) < epsilon {
		return
	}
	c1 := clipCode(x1, y1)
	c2 := clipCode(x2, y2)
	if (c1 & c2) != 0 {
		return
	}
	if (c1 | c2) == 0 {
		grsys.Move(x1, y1)
		grsys.Draw(x2, y2)
		return
	}
	bisects++
	xm := (x1 + x2) / 2
	ym := (y1 + y2) / 2
	clipDraw(x1, y1, xm, ym)
	clipDraw(xm, ym, x2, y2)
}
