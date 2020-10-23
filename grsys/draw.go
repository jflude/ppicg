package grsys

import (
	"fmt"
	"image"
)

var pen image.Point
var outside bool

func plotCoord(x float32) int {
	return int(1000*x + 0.5)
}

func IX(x float32) int {
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

func iy(y float32) int {
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

func IY(y float32) int {
	return ImageHeight - iy(y)
}

func Move(x, y float32) {
	pen.X = IX(x)
	pen.Y = IY(y)
	if plotFile != nil {
		fmt.Fprint(plotFile, "PU;PA", plotCoord(x-XMin), ",",
			plotCoord(y-YMin), ";\n")
	}
}

func Draw(x, y float32) {
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

func DrawLine(x1, y1, x2, y2 int) {
	y1 = ImageHeight - y1
	y2 = ImageHeight - y2
	dx := x2 - x1
	if dx < 0 {
		dx = -dx
	}
	sx := -1
	if x1 < x2 {
		sx = 1
	}
	dy := y1 - y2
	if dy > 0 {
		dy = -dy
	}
	sy := -1
	if y1 < y2 {
		sy = 1
	}
	exy := dx + dy /* error value e_xy */
	for {
		PutPix(x1, y1)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * exy
		if e2 >= dy { /* e_xy + e_x > 0 */
			exy += dy
			x1 += sx
		}
		if e2 <= dx { /* e_xy + e_y < 0 */
			exy += dx
			y1 += sy
		}
	}
}

func HorizLine(xLeft, xRight, y int) {
	y = ImageHeight - y
	DrawLine(xLeft, y, xRight, y)
}
