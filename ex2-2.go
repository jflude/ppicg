package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	fmt.Print("Enter coordinates of first triangle's vertices: ")
	var xA, yA, xB, yB, xC, yC float64
	if _, err := fmt.Scanln(&xA, &yA, &xB, &yB, &xC, &yC); err != nil {
		grsys.Error(err)
	}
	fmt.Print("Enter coordinates of center of rotation: ")
	var xP, yP float64
	if _, err := fmt.Scanln(&xP, &yP); err != nil {
		grsys.Error(err)
	}

	phi := 3 * math.Pi / 180
	cosphi := math.Cos(phi)
	sinphi := math.Sin(phi)
	c1 := xP - xP*cosphi + yP*sinphi
	c2 := yP - xP*sinphi - yP*cosphi

	grsys.InitGr("")
	for i := 0; i < 30; i++ {
		grsys.Move(xA, yA)
		grsys.Draw(xB, yB)
		grsys.Draw(xC, yC)
		grsys.Draw(xA, yA)

		xA = xA*cosphi - yA*sinphi + c1
		yA = xA*sinphi + yA*cosphi + c2
		xB = xB*cosphi - yB*sinphi + c1
		yB = xB*sinphi + yB*cosphi + c2
		xC = xC*cosphi - yC*sinphi + c1
		yC = xC*sinphi + yC*cosphi + c2
	}
	grsys.EndGr()
}
