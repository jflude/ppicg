// Conversion from world coordinates to viewport coordinates.
package grsys

import "math"

var (
	xMinWorld, xMaxWorld         float64
	yMinWorld, yMaxWorld         float64
	xCWorld, yCWorld             float64
	xCView, yCView               float64
	fScale                       float64
	iwCalled, wbCalled, vbCalled bool
	vpPlots                      []vpNode
)

func InitWindow() {
	xMinWorld = math.MaxFloat64
	yMinWorld = math.MaxFloat64
	xMaxWorld = -math.MaxFloat64
	yMaxWorld = -math.MaxFloat64
	iwCalled = true
}

func UpdateWindowBoundaries(x, y float64) {
	if !iwCalled {
		ErrorMsg("Call InitWindow before UpdateWindowBoundaries")
	}
	if x < xMinWorld {
		xMinWorld = x
	}
	if x > xMaxWorld {
		xMaxWorld = x
	}
	if y < yMinWorld {
		yMinWorld = y
	}
	if y > yMaxWorld {
		yMaxWorld = y
	}
	wbCalled = true
}

func ViewportBoundaries(Xmin, Xmax, Ymin, Ymax, reductionFactor float64) {
	if !inGrMode {
		ErrorMsg("Call InitGr before ViewportBoundaries")
	}
	if !wbCalled {
		ErrorMsg("Call UpdateWindowBoundaries before ViewportBoundaries")
	}
	xCView = (XMin + XMax) / 2
	yCView = (YMin + YMax) / 2
	fx := (XMax - XMin) / (xMaxWorld - xMinWorld + 1e-7) // +1e-7 to prevent
	fy := (YMax - YMin) / (yMaxWorld - yMinWorld + 1e-7) // division by zero
	if fx < fy {
		fScale = fx
	} else {
		fScale = fy
	}
	fScale *= reductionFactor
	xCWorld = (xMinWorld + xMaxWorld) / 2
	yCWorld = (yMinWorld + yMaxWorld) / 2
	vbCalled = true
}

func XViewport(x float64) float64 {
	if !vbCalled {
		ErrorMsg("Call ViewportBoundaries before XViewport")
	}
	return xCView + fScale*(x-xCWorld)
}

func YViewport(y float64) float64 {
	if !vbCalled {
		ErrorMsg("Call ViewportBoundaries before YViewport")
	}
	return yCView + fScale*(y-yCWorld)

}

type vpNode struct {
	x, y float64
	code int // 0 = pen up; 1 = pen down
}

func AppendPlot(x, y float64, code int) {
	// Store point (x, y) and plotcode (0=up, 1=down) in queue.
	if !iwCalled {
		ErrorMsg("Call InitWindow before AppendPlot")
	}
	vpPlots = append(vpPlots, vpNode{x, y, code})
	UpdateWindowBoundaries(x, y)
}

func GenPlot() {
	if !inGrMode {
		ErrorMsg("Call InitGr before GenPlot")
	}
	ViewportBoundaries(XMin, XMax, YMin, YMax, 0.9)
	for _, p := range vpPlots {
		x := XViewport(p.x)
		y := YViewport(p.y)
		if p.code != 0 {
			Draw(x, y)
		} else {
			Move(x, y)
		}
	}
}
