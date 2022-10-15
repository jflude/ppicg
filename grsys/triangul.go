// Triangulation of polygons.
// Triangulation of a polygon with successive vertex numbers
// pol[0], ..., pol[n-1], in counter-clockwise order.  With
// three given vertex numbers P, Q, R, function orienta must
// determine their orientation.
//   Negative = clockwise
//   Zero     = collinear
//   Positive = counter-clockwise
// If triangulation is possible, the resulting triangles are
// successively stored in array 'nrs'.  Triangle j has vertex
// numbers nrs[j].A, nrs[j].B, nrs[j].C.
// Memory space for slice 'nrs' must be supplied by the caller.
// Return value: the number of triangles found, or -1 if no
// proper polygon, or vertices are clockwise
package grsys

type Trianrs struct {
	A, B, C int
}

func Triangul(pol []int, n int, nrs []Trianrs,
	orienta func(int, int, int) int) int {
	if n < 3 {
		return -1 // no polygon
	}
	if n == 3 {
		nrs[0].A = pol[0]
		nrs[0].B = pol[1]
		nrs[0].C = pol[2]
		return 1 // only one triangle
	}
	polconvex := true
	ort := make([]int, n) // ort[i] = 1 if vertex i is convex
	for {
		collinear := false
		for i := 0; i < n; i++ {
			var i1, i2 int
			if i < n-1 {
				i1 = i + 1
			}
			if i1 < n-1 {
				i2 = i1 + 1
			}
			ort[i1] = orienta(pol[i], pol[i1], pol[i2])
			if ort[i1] == 0 {
				collinear = true
				for j := i1; j < n-1; j++ {
					pol[j] = pol[j+1]
				}
				n--
				break
			}
			if ort[i1] < 1 {
				polconvex = false
			}
		}
		if !collinear {
			break
		}
	}
	if n < 3 {
		return -1
	}
	if polconvex { // Use diagonals through vertex 0:
		for j := 0; j < n-2; j++ {
			nrs[j].A = pol[0]
			nrs[j].B = pol[j+1]
			nrs[j].C = pol[j+2]
		}
		return n - 2
	}
	ptr := make([]int, n) // ptr[i] is i's successor in list
	// Set up a circular list of vertex numbers.
	for i := 1; i < n; i++ {
		ptr[i-1] = i
	}
	ptr[n-1] = 0
	q := 0
	qA := ptr[q]
	qB := ptr[qA]
	qC := ptr[qB]
	r := -1
	j := 0                   // j triangles stored so far
	for m := n; m > 2; m-- { // m remaining nodes in circular list
		for k := 0; k < m; k++ {
			// Try triangle ABC:
			ortB := ort[qB]
			ok := false
			// B is a candidate if it is convex:
			if ortB > 0 {
				A := pol[qA]
				B := pol[qB]
				C := pol[qC]
				ok = true
				r = ptr[qC]
				for r != qA && ok {
					P := pol[r] // ABC counter-clockwise:
					ok = (P == A || P == B || P == C ||
						orienta(A, B, P) < 0 ||
						orienta(B, C, P) < 0 ||
						orienta(C, A, P) < 0)
					r = ptr[r]
				}
				// ok means: P coinciding with A, B or C
				// or outside ABC
				if ok {
					nrs[j].A = pol[qA]
					nrs[j].B = pol[qB]
					nrs[j].C = pol[qC]
					j++
				}
			}
			if ok || ortB == 0 {
				// Cut off triangle ABC from polygon:
				ptr[qA] = qC
				qB = qC
				qC = ptr[qC]
				if ort[qA] < 1 {
					ort[qA] = orienta(pol[q], pol[qA],
						pol[qB])
				}
				if ort[qB] < 1 {
					ort[qB] = orienta(pol[qA], pol[qB],
						pol[qC])
				}
				for ort[qA] == 0 && m > 2 {
					ptr[q] = qB
					qA = qB
					qB = qC
					qC = ptr[qC]
					m--
				}
				for ort[qB] == 0 && m > 2 {
					ptr[qA] = qC
					qB = qC
					qC = ptr[qC]
					m--
				}
				break
			}
			q = qA
			qA = qB
			qB = qC
			qC = ptr[qC]
		}
	}
	return j // j triangles
}
