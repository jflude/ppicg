// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
// Demonstration program dealing with pixels.
package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Print("Dimensions of elementary square (eg. 25): ")
	var k int
	if _, err := fmt.Scanln(&k); err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	n := (grsys.ImageHeight - 100) / k
	N := n * k
	xMargin := (grsys.ImageWidth - N) / 2
	yMargin := (grsys.ImageHeight - N) / 2
	for x := 0; x < n; x++ {
		x1 := xMargin + x*k
		xPlus2 := x + 2
		for y := 0; y < n; y++ {
			if gcd(xPlus2, y+2) == 1 {
				y1 := yMargin + y*k
				for i := 0; i < k; i++ {
					X := x1 + i
					for j := 0; j < k; j++ {
						grsys.PutPix(X, y1+j)
					}
				}
			}
		}
	}
	grsys.EndGr()
}
