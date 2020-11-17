package main

import (
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"math/rand"
	"time"
)

func Pythagoras(A, B grsys.Vec, n int) {
	type call struct {
		A, B grsys.Vec
		n    int
	}
	q := []call{call{A, B, n}}
	colors := rand.Perm(grsys.NColors)
	for len(q) > 0 {
		A := q[0].A
		B := q[0].B
		n := q[0].n
		q = q[1:]
		if n <= 0 {
			continue
		}
		C := B.Add(grsys.Vec{X: A.Y - B.Y, Y: B.X - A.X})
		D := A.Add(C.Sub(B))
		E := D.Add(C.Sub(A).Mul(0.5))
		grsys.SetColor(colors[n%grsys.NColors])
		grsys.Fill(
			[]int{
				grsys.IX(A.X),
				grsys.IX(B.X),
				grsys.IX(C.X),
				grsys.IX(D.X),
			},
			[]int{
				grsys.IY(A.Y),
				grsys.IY(B.Y),
				grsys.IY(C.Y),
				grsys.IY(D.Y),
			})
		grsys.SetColor(colors[(grsys.NColors-n)%grsys.NColors])
		grsys.Fill(
			[]int{
				grsys.IX(C.X),
				grsys.IX(D.X),
				grsys.IX(E.X),
			},
			[]int{
				grsys.IY(C.Y),
				grsys.IY(D.Y),
				grsys.IY(E.Y),
			})
		q = append(q, call{D, E, n - 1})
		q = append(q, call{E, C, n - 1})
	}
}

func main() {
	fmt.Print("Enter xA, yA, xB, yB and recursion depth n: ")
	var A, B grsys.Vec
	var n int
	if _, err := fmt.Scanln(&A.X, &A.Y, &B.X, &B.Y, &n); err != nil {
		grsys.Error(err)
	}
	rand.Seed(time.Now().UnixNano())
	grsys.InitGr("")
	Pythagoras(A, B, n)
	grsys.EndGr()
}
