package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func main() {
	var xA, yA, xB, yB, xC, yC float32
	if _, err := fmt.Scanln(&xA, &yA, &xB, &yB, &xC, &yC); err != nil {
		grsys.ErrorMsg(err.Error())
	}
	var q float32 = 0.05
	p := 1 - q

	grsys.Initgr()
	for i := 0; i < 50; i++ {
		grsys.Move(xA, yA)
		grsys.Draw(xB, yB)
		grsys.Draw(xC, yC)
		grsys.Draw(xA, yA)

		xT, yT := xA, yA
		xA, yA = p*xA+q*xB, p*yA+q*yB
		xB, yB = p*xB+q*xC, p*yB+q*yC
		xC, yC = p*xC+q*xT, p*yC+q*yT
	}
	grsys.Endgr()
}
