package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func setOfSquares(xA, yA float64, m int, p, a float64) {
	xB := xA + a
	yB := yA
	xC := xB
	yC := yA + a
	xD := xA
	yD := yC
	q := 1 - p

	for i := 0; i < m; i++ {
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
}

func main() {
	fmt.Println("There will be n x n sets of squares.")
	fmt.Print("Enter n (eg. 8 for a chess board): ")
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		grsys.ErrorMsg(err.Error())
	}
	halfn := float64(n) / 2
	fmt.Print("How many squares in each set? (eg. 10): ")
	var m int
	if _, err := fmt.Scanln(&m); err != nil {
		grsys.ErrorMsg(err.Error())
	}
	fmt.Print("Enter interpolation factor between 0 and 1 (eg. 0.2): ")
	var lambda float64
	if _, err := fmt.Scanln(&lambda); err != nil {
		grsys.ErrorMsg(err.Error())
	}
	grsys.Initgr("")
	a := 1.9 * grsys.RMax / float64(n) // length of side of largest square
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var f float64
			if (i+j)%2 == 1 {
				f = lambda
			} else {
				f = 1 - lambda
			}
			setOfSquares(grsys.XCenter+(float64(i)-halfn)*a,
				grsys.YCenter+(float64(j)-halfn)*a,
				m, f, a)
		}
	}
	grsys.Endgr()
}
