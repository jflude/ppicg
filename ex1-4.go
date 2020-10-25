package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	fmt.Print("Enter the radius of the hexagons: ")
	var r float64
	if _, err := fmt.Scanln(&r); err != nil {
		grsys.ErrorMsg(err.Error())
	}
	grsys.InitGr("")

	width := grsys.XMax - grsys.XMin
	nRow := math.Floor((width/r + 1) / 3)
	xOrig := r + (width-(r*(3*nRow-1)))/2

	height := grsys.YMax - grsys.YMin
	h := r * math.Sqrt(3)
	nCol := math.Floor(height / h)
	yOrig := h/2 + (height-(nCol*h))/2

	for i, y := int(nCol), yOrig; i > 0; i, y = i-1, y+h {
		row(nRow, xOrig, y, r)
	}
	for i, y := int(nCol)-1, yOrig; i > 0; i, y = i-1, y+h {
		row(nRow-1, xOrig+3*r/2, y+h/2, r)
	}
	grsys.EndGr()
}

func row(nRow, xOrig, y, r float64) {
	w := 3 * r
	for i, x := int(nRow), xOrig; i > 0; i, x = i-1, x+w {
		hexagon(x, y, r)
	}
}

func hexagon(xCenter, yCenter, radius float64) {
	var x1st, y1st float64
	for i := 0; i < 6; i++ {
		theta := 2 * math.Pi * float64(i) / 6
		x := xCenter + radius*math.Cos(theta)
		y := yCenter + radius*math.Sin(theta)
		if i == 0 {
			grsys.Move(x, y)
			x1st = x
			y1st = y
		} else {
			grsys.Draw(x, y)
		}
	}
	grsys.Draw(x1st, y1st)
}
