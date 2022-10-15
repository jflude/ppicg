// Demonstration program for polygon filling.
// The name of an input file must be given as a program argument.
// This file must contain n pairs of pixel coordinates.
package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"io"
	"os"
)

func main() {
	var X, Y []int
	if len(os.Args) < 2 {
		grsys.Error(errors.New("No valid input file as program argument"))
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		grsys.Error(err)
	}
	for {
		var x, y int
		if _, err := fmt.Fscanln(f, &x, &y); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			grsys.Error(err)
		}
		X = append(X, x)
		Y = append(Y, y)
	}
	grsys.InitGr("")
	grsys.Fill(X, Y) // This draws and fills the polygon.
	grsys.EndGr()
}
