// This program reads data from any text file containing lines of the
// form x y code (code = 1: pen down; code = 0: pen up), and displays
// lines, all fitting into a given viewport.  The file name is to be
// supplied as a program argument.
package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		grsys.ErrorMsg("No valid input file as program argument")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		grsys.ErrorMsg(err.Error())
	}

	grsys.InitWindow()
	var x, y float64
	var code int
	for {
		_, err := fmt.Fscanln(f, &x, &y, &code)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			grsys.ErrorMsg(err.Error())
		}
		grsys.UpdateWindowBoundaries(x, y)
	}
	if _, err := f.Seek(0, os.SEEK_SET); err != nil {
		grsys.ErrorMsg(err.Error())
	}

	grsys.InitGr("adjust.hpg") // plot file desired
	grsys.ViewportBoundaries(grsys.XMin, grsys.XMax,
		grsys.YMin, grsys.YMax, 0.9)

	grsys.Move(grsys.XMin, grsys.YMin)
	grsys.Draw(grsys.XMax, grsys.YMin)
	grsys.Draw(grsys.XMax, grsys.YMax)
	grsys.Draw(grsys.XMin, grsys.YMax)
	grsys.Draw(grsys.XMin, grsys.YMin) // show viewport

	for {
		if _, err := fmt.Fscanln(f, &x, &y, &code); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			grsys.ErrorMsg(err.Error())
		}
		x := grsys.XViewport(x)
		y := grsys.YViewport(y)
		if code != 0 {
			grsys.Draw(x, y)
		} else {
			grsys.Move(x, y)
		}
	}
	grsys.EndGr()
}
