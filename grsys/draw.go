// Line-drawing primitives.
package grsys

import (
	"fmt"
	"image"
)

var (
	ImageWidth  int = 800 // X__max
	ImageHeight int = 600 // Y__max

	canvas  *image.Paletted
	pen     image.Point
	outside bool
)

func plotCoord(x float64) int {
	return int(1000*x + 0.5)
}

func IX(x float64) int {
	ix := int(Density * (x - XMin))
	if ix < 0 {
		ix = 0
		outside = true
	} else if ix > ImageWidth {
		ix = ImageWidth
		outside = true
	}
	return ix
}

func iy(y float64) int {
	iy := int(Density * (y - YMin))
	if iy < 0 {
		iy = 0
		outside = true
	} else if iy > ImageHeight {
		iy = ImageHeight
		outside = true
	}
	return iy
}

func IY(y float64) int {
	return ImageHeight - iy(y)
}

func Move(x, y float64) {
	pen.X = IX(x)
	pen.Y = IY(y)
	if plotFile != nil {
		fmt.Fprint(plotFile, "PU;PA", plotCoord(x-XMin), ",",
			plotCoord(y-YMin), ";\n")
	}
}

func Draw(x, y float64) {
	to := image.Pt(IX(x), IY(y))
	DrawLine(pen.X, pen.Y, to.X, to.Y)
	pen = to
	if plotFile != nil {
		fmt.Fprint(plotFile, "PD;PA", plotCoord(x-XMin), ",",
			plotCoord(y-YMin), ";\n")
	}
}

func PutPix(x, y int) {
	y = ImageHeight - y
	if (image.Point{x, y}.In(canvas.Bounds())) {
		i := canvas.PixOffset(x, y)
		canvas.Pix[i] = uint8(ForeGrColor)
	}
}

func DrawLine(xP, yP, xQ, yQ int) {
	yP = ImageHeight - yP
	yQ = ImageHeight - yQ
	dx := xQ - xP
	dy := yQ - yP
	xInc := 1
	yInc := 1
	if dx < 0 {
		xInc = -1
		dx = -dx
	}
	if dy < 0 {
		yInc = -1
		dy = -dy
	}
	x := xP
	y := yP
	D := 0
	if dy < dx {
		c := 2 * dx
		M := 2 * dy
		for x != xQ {
			PutPix(x, y)
			x += xInc
			D += M
			if D > dx {
				y += yInc
				D -= c
			}
		}
	} else {
		c := 2 * dy
		M := 2 * dx
		for y != yQ {
			PutPix(x, y)
			y += yInc
			D += M
			if D > dy {
				x += xInc
				D -= c
			}
		}
	}
}

func HorLine(xLeft, xRight, y int) {
	DrawLine(xLeft, y, xRight, y)
}
