// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
// Initialization and output primitives.
package grsys

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"
)

var (
	XMin    float64 = 0
	XMax    float64 = 10
	YMin    float64 = 0
	YMax    float64
	XCenter float64
	YCenter float64
	RMax    float64
	Density float64

	plotFile io.WriteCloser
	inGrMode bool
)

func InitGr(filename string) {
	NColors = GetMaxColor() + 1
	palette = make([]color.Color, NColors)
	for i := range palette {
		palette[i] = defaultPalette[i]
	}
	bounds := image.Rect(0, 0, ImageWidth, ImageHeight)
	canvas = image.NewPaletted(bounds, palette)

	SetRGBPalette(0, 0, 0, 63)
	SetRGBPalette(14, 63, 63, 0)
	BackGrColor = 0
	ForeGrColor = 14

	Density = float64(ImageWidth) / (XMax - XMin)
	YMax = YMin + float64(ImageHeight)/Density
	XCenter = (XMin + XMax) / 2
	YCenter = (YMin + YMax) / 2
	if XCenter < YCenter {
		RMax = XCenter
	} else {
		RMax = YCenter
	}
	if filename != "" {
		var err error
		plotFile, err = os.Create(filename)
		if err != nil {
			Error(err)
		}
		fmt.Fprint(plotFile, "IN;SP0;SC0,10000,0,",
			plotCoord(YMax-YMin), ";\n")
	}
	inGrMode = true
}

func EndGr() {
	inGrMode = false
	if outside {
		log.SetFlags(0)
		log.Println("Warning: attempts to draw outside the screen")
	}
	if plotFile != nil {
		if err := plotFile.Close(); err != nil {
			Error(err)
		}
		plotFile = nil
	}
	encode("ppicg.gif")
}

func Error(err error) {
	log.SetFlags(0)
	log.Fatal(err)
}
