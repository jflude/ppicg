package grsys

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"os"
	"os/user"
	"path"
)

var (
	XMin    float32 = 0
	XMax    float32 = 10
	YMin    float32 = 0
	YMax    float32
	XCenter float32
	YCenter float32
	RMax    float32
	Density float32

	ImageWidth  int = 1024 // X__max
	ImageHeight int = 768  // Y__max
	NColors     int

	ForeGrColor int = 14
	BackGrColor int = 0

	palette  color.Palette
	canvas   *image.Paletted
	plotFile io.WriteCloser
)

func Initgr(filename string) {
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

	Density = float32(ImageWidth) / (XMax - XMin)
	YMax = YMin + float32(ImageHeight)/Density
	XCenter = (XMin + XMax) / 2
	YCenter = (YMin + YMax) / 2
	if XCenter < YCenter {
		RMax = XCenter
	} else {
		RMax = YCenter
	}
	if filename != "" {
		var err error
		plotFile, err = os.Create(getWorkPath(filename))
		if err != nil {
			ErrorMsg(err.Error())
		}
		fmt.Fprint(plotFile, "IN;SP0;SC0,10000,0,",
			plotCoord(YMax-YMin), ";\n")
	}
}

func Endgr() {
	if outside {
		log.SetFlags(0)
		log.Println("Warning: attempts to draw outside the screen")
	}
	if plotFile != nil {
		if err := plotFile.Close(); err != nil {
			ErrorMsg(err.Error())
		}
		plotFile = nil
	}
	f, err := os.Create(getWorkPath("ppicg.gif"))
	if err != nil {
		ErrorMsg(err.Error())
	}
	if err := gif.Encode(f, canvas, nil); err != nil {
		f.Close()
		ErrorMsg(err.Error())
	}
	if err := f.Close(); err != nil {
		ErrorMsg(err.Error())
	}
}

func ErrorMsg(msg string) {
	log.SetFlags(0)
	log.Fatal(msg)
}

func getWorkPath(filename string) string {
	u, err := user.Current()
	if err != nil {
		ErrorMsg(err.Error())
	}
	if u.HomeDir == "" {
		return filename
	}
	return path.Join(u.HomeDir, "Desktop", filename)
}
