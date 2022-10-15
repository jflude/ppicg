// Copyright (c) 2020 Justin Flude.
// Use of this source code is governed by the COPYING.md file.
// Implementation of Cohen-Sutherland line clipping.
package grsys

var xMinClip, xMaxClip, yMinClip, yMaxClip float64

func SetClipBoundaries(x1, x2, y1, y2 float64) {
	xMinClip = x1
	xMaxClip = x2
	yMinClip = y1
	yMaxClip = y2
}

func clipCode(x, y float64) int {
	var c int
	if x < xMinClip {
		c |= 8
	} else if x > xMaxClip {
		c |= 4
	}
	if y < yMinClip {
		c |= 2
	} else if y > yMaxClip {
		c |= 1
	}
	return c
}

func ClipDraw(xP, yP, xQ, yQ float64) {
	cP := clipCode(xP, yP)
	cQ := clipCode(xQ, yQ)
	for (cP | cQ) != 0 {
		if (cP & cQ) != 0 {
			return
		}
		dx := xQ - xP
		dy := yQ - yP
		if cP != 0 {
			if (cP & 8) != 0 {
				yP += (xMinClip - xP) * dy / dx
				xP = xMinClip
			} else if (cP & 4) != 0 {
				yP += (xMaxClip - xP) * dy / dx
				xP = xMaxClip
			} else if (cP & 2) != 0 {
				xP += (yMinClip - yP) * dx / dy
				yP = yMinClip
			} else if (cP & 1) != 0 {
				xP += (yMaxClip - yP) * dx / dy
				yP = yMaxClip
			}
			cP = clipCode(xP, yP)
		} else {
			if (cQ & 8) != 0 {
				yQ += (xMinClip - xQ) * dy / dx
				xQ = xMinClip
			} else if (cQ & 4) != 0 {
				yQ += (xMaxClip - xQ) * dy / dx
				xQ = xMaxClip
			} else if (cQ & 2) != 0 {
				xQ += (yMinClip - yQ) * dx / dy
				yQ = yMinClip
			} else if (cQ & 1) != 0 {
				xQ += (yMaxClip - yQ) * dx / dy
				yQ = yMaxClip
			}
			cQ = clipCode(xQ, yQ)
		}
	}
	Move(xP, yP)
	Draw(xQ, yQ)
}
