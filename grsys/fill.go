// Polygon fill function, using integers only.
package grsys

func Fill(X []int, Y []int) {
	type element struct {
		xP, yQ    int
		dx, dy, E int
		next      *element
	}
	n := len(X)
	x := X[n-1]
	y := Y[n-1]
	var ymin, ymax int
	for i := 0; i < n; i++ {
		DrawLine(x, y, X[i], Y[i])
		x = X[i]
		y = Y[i]
		if y < ymin {
			ymin = y
		}
		if y > ymax {
			ymax = y
		}
	}
	ny := ymax - ymin + 1
	table := make([]*element, ny)
	for i := 0; i < n; i++ {
		i1 := i + 1
		if i1 == n {
			i1 = 0 // i1 is i's successor
		}
		xP := X[i]
		yP := Y[i]
		xQ := X[i1]
		yQ := Y[i1]
		if yP == yQ {
			continue
		}
		if yQ < yP {
			xP, xQ = xQ, xP
			yP, yQ = yQ, yP
		}
		j := yP - ymin
		p := &element{xP, yQ, xQ - xP, yQ - yP, 0, table[j]}
		table[j] = p
	}
	start := new(element)
	end := start // Sentinel
	for j := 0; j < ny; j++ {
		y = ymin + j // Build or update active edge list:
		for p := start; p != end; {
			if p.yQ == y {
				// Delete list element *p:
				q := p.next
				if q == end {
					end = p
				} else {
					*p = *q
				}
			} else { // Update list element *p:
				dx := p.dx
				if dx != 0 {
					x = p.xP
					dy := p.dy
					E := p.E
					m := dx / dy // Integer division!
					dyQ := 2 * dy
					x += m
					E += 2*dx - m*dyQ
					if E > dy || E < -dy {
						if dx > 0 {
							x++
							E -= dyQ
						} else {
							x--
							E += dyQ
						}
					}
					p.xP = x
					p.E = E
				}
				p = p.next
			}
		}
		// End of updating the elements (if any) of active list.
		// Edges may now be added to active edge list:
		for p := table[j]; p != nil; {
			x = p.xP
			end.xP = x
			yQ := p.yQ
			dx := p.dx
			dy := p.dy
			q := start
			for q.xP < x ||
				(q.xP == x && q != end && q.dx*dy < dx*q.dy) {
				q = q.next
			}
			p0 := p
			p = p.next
			if q == end {
				end = p0
			} else {
				*p0 = *q
			}
			q.xP = x
			q.yQ = yQ
			q.dx = dx
			q.dy = dy
			q.E = 0
			q.next = p0
		}
		// Draw line segments:
		for p := start; p != end; p = p.next {
			xleft := p.xP + 1
			p = p.next
			xright := p.xP
			if xleft <= xright {
				HorLine(xleft, xright, y)
			}
		}
	}
}
