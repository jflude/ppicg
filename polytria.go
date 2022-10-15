// Dividing a polygon into triangles.
package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"io"
	"os"
)

const EPS = 1e-6

var p []grsys.Vec // n vertices p[0], ..., p[n-1] are given.

func drawPolygon(p []grsys.Vec, n int) {
	p[n-1].Move()
	for i := 0; i < n; i++ {
		p[i].Draw()
	}
}

func drawTriangles(p []grsys.Vec, nrs []grsys.Trianrs, m int) {
	for j := 0; j < m; j++ {
		A := p[nrs[j].A]
		B := p[nrs[j].B]
		C := p[nrs[j].C]
		centroid := A.Add(B.Add(C)).Mul(1.0 / 3)
		A1 := centroid.Add(A.Sub(centroid).Mul(0.9))
		B1 := centroid.Add(B.Sub(centroid).Mul(0.9))
		C1 := centroid.Add(C.Sub(centroid).Mul(0.9))
		A1.Move()
		B1.Draw()
		C1.Draw()
		A1.Draw()
	}
}

func orienta(Pnr, Qnr, Rnr int) int {
	A := p[Qnr].Sub(p[Pnr])
	B := p[Rnr].Sub(p[Pnr])
	det := A.X*B.Y - A.Y*B.X
	if det < -EPS {
		return -1
	} else if det > EPS {
		return 1
	}
	return 0
}

func main() {
	var n int
	var f *os.File
	var err error
	if len(os.Args) < 2 {
		for {
			fmt.Println("Enter n, followed by the coordinate pairs (x, y)")
			fmt.Print("of n vertices, in counter-clockwise order: ")
			if _, err = fmt.Scan(&n); err != nil {
				grsys.Error(err)
			}
			if n >= 3 {
				break
			}
		}
	} else {
		if f, err = os.Open(os.Args[1]); err != nil {
			grsys.Error(err)
		}
		if _, err = fmt.Fscan(f, &n); err != nil {
			grsys.Error(err)
		}
	}
	p = make([]grsys.Vec, n)
	// Vertex number of triangles in nrspol[0]...
	nrspol := make([]int, n)
	nrs := make([]grsys.Trianrs, n-2) // At most n-2 triangles
	for i := 0; i < n; i++ {
		nrspol[i] = i
		if len(os.Args) < 2 {
			_, err = fmt.Scan(&p[i].X, &p[i].Y)
		} else {
			_, err = fmt.Fscan(f, &p[i].X, &p[i].Y)
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = io.ErrUnexpectedEOF
			}
			grsys.Error(err)
		}
	}
	ntria := grsys.Triangul(nrspol, n, nrs, orienta)
	if ntria == -1 {
		grsys.Error(errors.New("Polygon specification incorrect"))
	}
	if len(os.Args) > 2 {
		grsys.InitGr(os.Args[2])
	} else {
		grsys.InitGr("")
	}
	drawPolygon(p, n)
	drawTriangles(p, nrs, ntria)
	grsys.EndGr()
	fmt.Println(n, "vertices;", ntria, "triangles")
}
