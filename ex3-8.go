package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

const EPS = 1e-6

func main() {
	fmt.Print("Enter the coordinates of the triangle's vertices: ")
	var a, b, c grsys.Vec
	_, err := fmt.Scanln(&a.X, &a.Y, &b.X, &b.Y, &c.X, &c.Y)
	if err != nil {
		grsys.Error(err)
	}
	ab := b.Sub(a)
	bc := c.Sub(b)
	orient := ab.X*bc.Y - ab.Y*bc.X
	if math.Abs(orient) < EPS {
		grsys.Error(errors.New("vertices are collinear"))
	}
	if orient < 0 {
		a, c = c, a
		ab = b.Sub(a)
		bc = c.Sub(b)
	}
	ca := a.Sub(c)

	grsys.InitGr("")
	grsys.InitWindow()
	grsys.AppendPlot(a.X, a.Y, 0)
	grsys.AppendPlot(b.X, b.Y, 1)
	grsys.AppendPlot(c.X, c.Y, 1)
	grsys.AppendPlot(a.X, a.Y, 1)

	ab2 := ab.X*ab.X + ab.Y*ab.Y
	bc2 := bc.X*bc.X + bc.Y*bc.Y
	ca2 := ca.X*ca.X + ca.Y*ca.Y

	abLen := math.Sqrt(ab2)
	bcLen := math.Sqrt(bc2)
	caLen := math.Sqrt(ca2)

	angA := math.Acos((ca2 + ab2 - bc2) / (2 * caLen * abLen))
	angB := math.Acos((ab2 + bc2 - ca2) / (2 * abLen * bcLen))
	angC := math.Pi - angA - angB

	// find intersection of bisectors of A and B angles
	a1 := math.Tan(math.Atan2(ab.Y, ab.X) + angA/2)
	b1 := -1.0
	c1 := a.X*a1 - a.Y
	a2 := math.Tan(math.Atan2(bc.Y, bc.X) + angB/2)
	b2 := -1.0
	c2 := b.X*a2 - b.Y
	d := a1*b2 - a2*b1
	d1 := c1*b2 - c2*b1
	d2 := a1*c2 - a2*c1
	ic := grsys.Vec{X: d1 / d, Y: d2 / d}

	ai := ic.Sub(a)
	r := math.Sqrt(ai.X*ai.X+ai.Y*ai.Y) * math.Sin(angA/2)
	drawCircle(ic, r)

	// find intersection of perpendiculars of bisectors of A and B angles
	a1 = -1 / (a1 + 0.0000001)
	c1 = a.X*a1 - a.Y
	a2 = -1 / (a2 + 0.0000001)
	c2 = b.X*a2 - b.Y
	d = a1*b2 - a2*b1
	d1 = c1*b2 - c2*b1
	d2 = a1*c2 - a2*c1
	s := (abLen + bcLen + caLen) / 2
	area := r * s
	drawCircle(grsys.Vec{X: d1 / d, Y: d2 / d}, area/(s-abLen))

	// find intersection of perpendiculars of bisectors of B and C angles
	a1 = -1 / (math.Tan(math.Atan2(ca.Y, ca.X)+angC/2) + 0.0000001)
	c1 = c.X*a1 - c.Y
	d = a1*b2 - a2*b1
	d1 = c1*b2 - c2*b1
	d2 = a1*c2 - a2*c1
	drawCircle(grsys.Vec{X: d1 / d, Y: d2 / d}, area/(s-bcLen))

	// find intersection of perpendiculars of bisectors of C and A angles
	a2 = -1 / (math.Tan(math.Atan2(ab.Y, ab.X)+angA/2) + 0.0000001)
	c2 = a.X*a2 - a.Y
	d = a1*b2 - a2*b1
	d1 = c1*b2 - c2*b1
	d2 = a1*c2 - a2*c1
	drawCircle(grsys.Vec{X: d1 / d, Y: d2 / d}, area/(s-caLen))

	grsys.GenPlot()
	grsys.EndGr()
}

func drawCircle(c grsys.Vec, r float64) {
	var q grsys.Vec
	for th := 0.0; th <= 2*math.Pi; th += math.Pi / 50 {
		p := c.Add(grsys.Vec{X: r * math.Cos(th), Y: r * math.Sin(th)})
		if th == 0 {
			grsys.AppendPlot(p.X, p.Y, 0)
			q = p
		} else {
			grsys.AppendPlot(p.X, p.Y, 1)
		}
	}
	grsys.AppendPlot(q.X, q.Y, 1)
}
