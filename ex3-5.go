package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math"
)

const EPS = 1e-6

func calcCircumcenter(p, q, r grsys.Vec) grsys.Vec {
	var a1, b1, c1 float64
	pq := q.Sub(p)
	pq2 := p.Add(pq.Mul(0.5))
	if math.Abs(pq.Y) > EPS {
		a1 = -pq.X / pq.Y
		b1 = -1
		c1 = a1*pq2.X - pq2.Y
	} else {
		a1 = 1
		b1 = 0
		c1 = pq2.X
	}
	var a2, b2, c2 float64
	pr := r.Sub(p)
	pr2 := p.Add(pr.Mul(0.5))
	if math.Abs(pr.Y) > EPS {
		a2 = -pr.X / pr.Y
		b2 = -1
		c2 = a2*pr2.X - pr2.Y
	} else {
		a2 = 1
		b2 = 0
		c2 = pr2.X
	}
	D := a1*b2 - a2*b1
	D1 := c1*b2 - c2*b1
	D2 := a1*c2 - a2*c1
	return grsys.Vec{X: D1 / D, Y: D2 / D}
}

func main() {
	fmt.Print("Enter the coordinates of the triangle vertices: ")
	var a, b, c grsys.Vec
	_, err := fmt.Scanln(&a.X, &a.Y, &b.X, &b.Y, &c.X, &c.Y)
	if err != nil {
		grsys.Error(err)
	}
	grsys.InitGr("")
	a.Move()
	b.Draw()
	c.Draw()
	a.Draw()
	center := calcCircumcenter(a, b, c)
	v := a.Sub(center)
	r := math.Sqrt(v.X*v.X + v.Y*v.Y)
	var q grsys.Vec
	for i := 0; i < 50; i++ {
		th := 2 * math.Pi * float64(i) / 50
		p := grsys.Vec{X: r * math.Cos(th), Y: r * math.Sin(th)}
		p = center.Add(p)
		if i == 0 {
			q = p
			p.Move()
		} else {
			p.Draw()
		}
	}
	q.Draw()
	grsys.EndGr()
}
