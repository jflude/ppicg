// Generation of a random curve.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("How many line segments (eg. 500): ")
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		grsys.Error(err)
	}

	grsys.InitWindow()
	var x, y float64
	grsys.AppendPlot(x, y, 0)
	rand.Seed(time.Now().UnixNano())

	var alpha int
	var phi float64
	for ; n > 0; n-- {
		theta := rand.Int()%91 - 45
		alpha = (alpha + theta) / 2
		phi += float64(alpha) * math.Pi / 180
		x += math.Cos(phi)
		y += math.Sin(phi)
		grsys.AppendPlot(x, y, 1)
	}

	grsys.InitGr("curve.hpg") // plot file desired
	grsys.Move(grsys.XMin, grsys.YMin)
	grsys.Draw(grsys.XMax, grsys.YMin)
	grsys.Draw(grsys.XMax, grsys.YMax)
	grsys.Draw(grsys.XMin, grsys.YMax)
	grsys.Draw(grsys.XMin, grsys.YMin) // show viewport
	grsys.GenPlot()
	grsys.EndGr()
}
