package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func main() {
	fmt.Print("Enter the number of intervals for hatching: ")
	var m int
	if _, err := fmt.Scanln(&m); err != nil {
		grsys.ErrorMsg(err.Error())
	}

	grsys.Initgr("")
	var width, height, xOrg, yOrg float64
	if grsys.ImageWidth > grsys.ImageHeight {
		height = grsys.YMax - grsys.YMin
		width = height
		xOrg = (grsys.XMax - grsys.YMax) / 2
	} else {
		width = grsys.XMax - grsys.XMin
		height = width
		yOrg = (grsys.YMax - grsys.XMax) / 2
	}
	xOpp := xOrg + width
	yOpp := yOrg + height

	grsys.Move(xOrg, yOrg)
	grsys.Draw(xOpp, yOrg)
	grsys.Draw(xOpp, yOpp)
	grsys.Draw(xOrg, yOpp)
	grsys.Draw(xOrg, yOrg)

	xSqInc := width / 8
	ySqInc := height / 8
	for x := xOrg; x <= xOpp; x += xSqInc {
		grsys.Move(x, yOrg)
		grsys.Draw(x, yOpp)
	}
	for y := yOrg; y <= yOpp; y += ySqInc {
		grsys.Move(xOrg, y)
		grsys.Draw(xOpp, y)
	}

	xHchInc := xSqInc / float64(m)
	yHchInc := ySqInc / float64(m)
	firstBlack := true
	for ySq := yOrg; ySq < yOpp; ySq += ySqInc {
		firstBlack = !firstBlack
		isBlack := firstBlack
		for xSq := xOrg; xSq < xOpp; xSq += xSqInc {
			if isBlack = !isBlack; !isBlack {
				continue
			}
			for i := 0.0; i < float64(m); i += 1 {
				grsys.Move(xSq+i*xHchInc, ySq)
				grsys.Draw(xSq+xSqInc, ySq+ySqInc-i*yHchInc)
				grsys.Move(xSq, ySq+i*yHchInc)
				grsys.Draw(xSq+xSqInc-i*xHchInc, ySq+ySqInc)
			}
		}
	}
	grsys.Endgr()
}
