package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	fmt.Print("Enter x0, y0, R and N: ")
	var x0, y0, R float64
	var N int
	if _, err := fmt.Scanln(&x0, &y0, &R, &N); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	for i := 1; i < N; i++ {
		first := true
		for phi := 0.0; phi <= 360; phi += 6 {
			theta := phi * math.Pi / 180
			x := x0 + float64(i)*R/float64(N)*math.Cos(theta)
			y := y0 + float64(N-i)*R/float64(N)*math.Sin(theta)
			if first {
				first = false
				grsys.Move(x, y)
			} else {
				grsys.Draw(x, y)
			}
		}
	}
	grsys.EndGr()
}
