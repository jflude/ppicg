package main

import (
	"errors"
	"fmt"
	"github.com/jflude/ppicg/grsys"
	"io"
	"math"
)

type vec4 [4]float64
type mat4x4 [4][4]float64

var rgen mat4x4

func (v *vec4) mulVec(u *vec4) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]

}

func (m *mat4x4) mulVec(v *vec4) (r vec4) {
	r[0] = v[0]*m[0][0] + v[1]*m[1][0] + v[2]*m[2][0] + v[3]*m[3][0]
	r[1] = v[0]*m[0][1] + v[1]*m[1][1] + v[2]*m[2][1] + v[3]*m[3][1]
	r[2] = v[0]*m[0][2] + v[1]*m[1][2] + v[2]*m[2][2] + v[3]*m[3][2]
	r[3] = v[0]*m[0][3] + v[1]*m[1][3] + v[2]*m[2][3] + v[3]*m[3][3]
	return
}

func (m *mat4x4) mulMat(n *mat4x4) {
	var r mat4x4
	r[0][0] = m[0][0]*n[0][0] + m[0][1]*n[1][0] + m[0][2]*n[2][0] + m[0][3]*n[3][0]
	r[0][1] = m[0][0]*n[0][1] + m[0][1]*n[1][1] + m[0][2]*n[2][1] + m[0][3]*n[3][1]
	r[0][2] = m[0][0]*n[0][2] + m[0][1]*n[1][2] + m[0][2]*n[2][2] + m[0][3]*n[3][2]
	r[0][3] = m[0][0]*n[0][3] + m[0][1]*n[1][3] + m[0][2]*n[2][3] + m[0][3]*n[3][3]

	r[1][0] = m[1][0]*n[0][0] + m[1][1]*n[1][0] + m[1][2]*n[2][0] + m[1][3]*n[3][0]
	r[1][1] = m[1][0]*n[0][1] + m[1][1]*n[1][1] + m[1][2]*n[2][1] + m[1][3]*n[3][1]
	r[1][2] = m[1][0]*n[0][2] + m[1][1]*n[1][2] + m[1][2]*n[2][2] + m[1][3]*n[3][2]
	r[1][3] = m[1][0]*n[0][3] + m[1][1]*n[1][3] + m[1][2]*n[2][3] + m[1][3]*n[3][3]

	r[2][0] = m[2][0]*n[0][0] + m[2][1]*n[1][0] + m[2][2]*n[2][0] + m[2][3]*n[3][0]
	r[2][1] = m[2][0]*n[0][1] + m[2][1]*n[1][1] + m[2][2]*n[2][1] + m[2][3]*n[3][1]
	r[2][2] = m[2][0]*n[0][2] + m[2][1]*n[1][2] + m[2][2]*n[2][2] + m[2][3]*n[3][2]
	r[2][3] = m[2][0]*n[0][3] + m[2][1]*n[1][3] + m[2][2]*n[2][3] + m[2][3]*n[3][3]

	r[3][0] = m[3][0]*n[0][0] + m[3][1]*n[1][0] + m[3][2]*n[2][0] + m[3][3]*n[3][0]
	r[3][1] = m[3][0]*n[0][1] + m[3][1]*n[1][1] + m[3][2]*n[2][1] + m[3][3]*n[3][1]
	r[3][2] = m[3][0]*n[0][2] + m[3][1]*n[1][2] + m[3][2]*n[2][2] + m[3][3]*n[3][2]
	r[3][3] = m[3][0]*n[0][3] + m[3][1]*n[1][3] + m[3][2]*n[2][3] + m[3][3]*n[3][3]
	*m = r
}

func initRotate(a1, a2, a3, v1, v2, v3, alpha float64) {
	rho := math.Sqrt(v1*v1 + v2*v2 + v3*v3)
	theta := math.Atan2(v2, v1)
	phi := math.Acos(v3 / rho)

	ca := math.Cos(alpha)
	sa := math.Sin(alpha)
	cp := math.Cos(phi)
	sp := math.Sin(phi)
	ct := math.Cos(theta)
	st := math.Sin(theta)

	rgen = mat4x4{ // T(inv)
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{-a1, -a2, -a3, 1},
	}
	rgen.mulMat(&mat4x4{ // Rz(inv)
		{ct, -st, 0, 0},
		{st, ct, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
	rgen.mulMat(&mat4x4{ // Ry(inv) (book errata - uses phi, not alpha)
		{cp, 0, sp, 0},
		{0, 1, 0, 0},
		{-sp, 0, cp, 0},
		{0, 0, 0, 1},
	})
	rgen.mulMat(&mat4x4{ // Rv
		{ca, sa, 0, 0},
		{-sa, ca, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
	rgen.mulMat(&mat4x4{ // Ry (book errata - uses phi, not alpha)
		{cp, 0, -sp, 0},
		{0, 1, 0, 0},
		{sp, 0, cp, 0},
		{0, 0, 0, 1},
	})
	rgen.mulMat(&mat4x4{ // Rz
		{ct, st, 0, 0},
		{-st, ct, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
	rgen.mulMat(&mat4x4{ // T
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{a2, a2, a3, 1},
	})
}

func rotate(x, y, z float64) (px1, py1, pz1 float64) {
	p := vec4{x, y, z, 1}
	q := rgen.mulVec(&p)
	return q[0], q[1], q[2]
}

func main() {
	var a1, a2, a3, v1, v2, v3, alpha float64
	fmt.Print("Enter the starting point's coordinates (a1, a2, a3): ")
	if _, err := fmt.Scanln(&a1, &a2, &a3); err != nil {
		grsys.Error(err)
	}
	fmt.Print("Enter the vector (v1, v2, v3) to rotate around: ")
	if _, err := fmt.Scanln(&v1, &v2, &v3); err != nil {
		grsys.Error(err)
	}
	fmt.Print("Enter the angle alpha to rotate around: ")
	if _, err := fmt.Scanln(&alpha); err != nil {
		grsys.Error(err)
	}
	initRotate(a1, a2, a3, v1, v2, v3, alpha*math.Pi/180)
	for {
		var x, y, z float64
		fmt.Print("Enter coordinates (x, y, z) of a point: ")
		if _, err := fmt.Scanln(&x, &y, &z); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			grsys.Error(err)
		}
		x2, y2, z2 := rotate(x, y, z)
		fmt.Printf("(%v, %v, %v) rotates to (%v, %v, %v)\n",
			x, y, z, x2, y2, z2)
	}
}
