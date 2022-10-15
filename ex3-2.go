// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func drawSquares(n int, x1, y1, x2, y2 float64) {
	if n == 0 {
		return
	}
	grsys.Move(x1, y1)
	grsys.Draw(x2, y1)
	grsys.Draw(x2, y2)
	grsys.Draw(x1, y2)
	grsys.Draw(x1, y1)

	d := (x2 - x1) / 3
	grsys.Move(x1+d, y1)
	grsys.Draw(x1+d, y2)
	grsys.Move(x2-d, y1)
	grsys.Draw(x2-d, y2)

	grsys.Move(x1, y1+d)
	grsys.Draw(x2, y1+d)
	grsys.Move(x1, y2-d)
	grsys.Draw(x2, y2-d)

	n--
	drawSquares(n, x1, y1, x1+d, y1+d)
	drawSquares(n, x2-d, y1, x2, y1+d)
	drawSquares(n, x2-d, y2-d, x2, y2)
	drawSquares(n, x1, y2-d, x1+d, y2)
}

func main() {
	fmt.Print("Enter the maximum depth of recursion: ")
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	drawSquares(n,
		grsys.XCenter-grsys.RMax,
		grsys.YCenter-grsys.RMax,
		grsys.XCenter+grsys.RMax,
		grsys.YCenter+grsys.RMax)
	grsys.EndGr()
}
