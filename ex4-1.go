package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

func main() {
	fmt.Print("Enter start and end coordinates of the line: ")
	var x1, y1, x2, y2 float64
	if _, err := fmt.Scanln(&x1, &y1, &x2, &y2); err != nil {
		grsys.Error(err)
	}
	fmt.Print("Enter the width of the line: ")
	var w float64
	if _, err := fmt.Scanln(&w); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	th := math.Atan2(x1-x2, y2-y1)
	dx := w * math.Cos(th)
	dy := w * math.Sin(th)
	grsys.Fill(
		[]int{
			grsys.IX(x1 - dx),
			grsys.IX(x1 + dx),
			grsys.IX(x2 + dx),
			grsys.IX(x2 - dx),
		},
		[]int{
			grsys.IY(y1 - dy),
			grsys.IY(y1 + dy),
			grsys.IY(y2 + dy),
			grsys.IY(y2 - dy),
		})
	grsys.EndGr()
}
