package algorithm

import "F.A.N.A.P/models"

func Overlap(m1 models.Point, m2 models.Point, e1 models.Point, e2 models.Point) bool {
	if e2.XR <= m1.XR {
		return false
	}
	if e1.XR >= m2.XR {
		return false
	}
	if e2.YR <= m1.YR {
		return false
	}
	if e1.YR >= m2.YR {
		return false
	}

	return true
}
