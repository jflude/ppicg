package grsys

type Vec struct {
	X, Y float64
}

func (v Vec) Add(u Vec) Vec {
	return Vec{v.X + u.X, v.Y + u.Y}
}

func (v Vec) Sub(u Vec) Vec {
	return Vec{v.X - u.X, v.Y - u.Y}
}

func (v Vec) Mul(c float64) Vec {
	return Vec{c * v.X, c * v.Y}
}

func (v Vec) Move() {
	Move(v.X, v.Y)
}

func (v Vec) Draw() {
	Draw(v.X, v.Y)
}

func Rotate(p, c Vec, cosphi, sinphi float64) Vec {
	dx := p.X - c.X
	dy := p.Y - c.Y
	return Vec{c.X + dx*cosphi - dy*sinphi,
		c.Y + dx*sinphi + dy*cosphi}
}
