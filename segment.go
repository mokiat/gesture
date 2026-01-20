package gesture

import "github.com/mokiat/gomath/dprec"

// Segment represents a line segment between two points in 2D space.
type Segment struct {

	// From is the starting point of the segment.
	From dprec.Vec2

	// To is the ending point of the segment.
	To dprec.Vec2
}

// Interpolation returns a point along the segment at parameter t,
func (s *Segment) Interpolation(t float64) dprec.Vec2 {
	return dprec.Vec2{
		X: dprec.Mix(s.From.X, s.To.X, t),
		Y: dprec.Mix(s.From.Y, s.To.Y, t),
	}
}

// Length returns the length of the segment.
func (s *Segment) Length() float64 {
	return dprec.Vec2Diff(s.To, s.From).Length()
}
