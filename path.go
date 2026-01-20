package gesture

import (
	"iter"

	"github.com/mokiat/gomath/dprec"
)

// Path represents a sequence of points in 2D space that make up a shape.
type Path []dprec.Vec2

// SegmentsIter returns an iterator over the segments of the path (i.e. the
// segments between each two consecutive points).
func (p Path) SegmentsIter() iter.Seq[Segment] {
	return func(yield func(Segment) bool) {
		for i := range len(p) - 1 {
			segment := Segment{
				From: p[i],
				To:   p[i+1],
			}
			if !yield(segment) {
				return
			}
		}
	}
}

// Length returns the total length of the path, which is the sum of the lengths
// of its segments.
func (p Path) Length() float64 {
	var result float64
	for segment := range p.SegmentsIter() {
		result += segment.Length()
	}
	return result
}

// ResampleIter returns an iterator that yields a specified number of points
// evenly distributed along the path.
func (p Path) ResampleIter(pointCount int) iter.Seq[dprec.Vec2] {
	if pointCount <= 1 {
		panic("at least two points are required for resampling")
	}
	return func(yield func(dprec.Vec2) bool) {
		targetSegmentLength := p.Length() / float64(pointCount-1)
		remainingLength := 0.0 // start with zero to yield the first point immediately
		count := 0
		for segment := range p.SegmentsIter() {
			segmentLength := segment.Length()
			if segmentLength < dprec.Epsilon {
				continue // prevent division by zero
			}
			for remainingLength <= segmentLength {
				point := segment.Interpolation(remainingLength / segmentLength)
				count++
				if !yield(point) {
					return
				}
				segment.From = point // shorten segment
				segmentLength -= remainingLength
				remainingLength = targetSegmentLength
			}
			remainingLength -= segmentLength
		}
		if count < pointCount { // due to precision issues we might be one short
			lastPoint := p[len(p)-1]
			if !yield(lastPoint) {
				return
			}
		}
	}
}
