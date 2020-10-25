// This program draws 50 squares inside each other.
package main

import "github.com/jflude/ppicg/grsys"

func main() {
	q := 0.05 // q = lambda (see discussion below)
	p := 1 - q
	grsys.InitGr("")

	r := 0.95 * grsys.RMax
	xA := grsys.XCenter - r
	yA := grsys.YCenter - r
	xB := grsys.XCenter + r
	yB := yA
	xC := xB
	yC := grsys.YCenter + r
	xD := xA
	yD := yC

	for i := 0; i < 50; i++ {
		grsys.Move(xA, yA)
		grsys.Draw(xB, yB)
		grsys.Draw(xC, yC)
		grsys.Draw(xD, yD)
		grsys.Draw(xA, yA)

		xT := xA
		yT := yA

		xA = p*xA + q*xB
		yA = p*yA + q*yB
		xB = p*xB + q*xC
		yB = p*yB + q*yC
		xC = p*xC + q*xD
		yC = p*yC + q*yD
		xD = p*xD + q*xT
		yD = p*yD + q*yT
	}

	grsys.EndGr()
}
