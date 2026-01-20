package gesture

import (
	"github.com/mokiat/gog/seq"
	"github.com/mokiat/gomath/dprec"
)

// VectorLength is the fixed length of gesture vectors.
const VectorLength = 64

// CreateVector creates a Vector by resampling the given path to the fixed
// length.
func CreateVector(path Path) Vector {
	var result Vector
	for i, point := range seq.Indexed(path.ResampleIter(VectorLength)) {
		result[i] = point
	}
	return result
}

// VectorDot computes the hyperspace dot product between two vectors.
func VectorDot(a, b *Vector) float64 {
	var result float64
	for i := range VectorLength {
		result += a[i].X*b[i].X + a[i].Y*b[i].Y
	}
	return dprec.Clamp(result, -1.0, 1.0)
}

// VectorCross computes the hyperspace cross product between two vectors.
//
// Note: There is no such thing as a cross product in higher dimensions, but
// this function computes a value similar to the 3D cross product magnitude by
// summing up the 2D cross products of each corresponding point pair.
func VectorCross(a, b *Vector) float64 {
	var result float64
	for i := range VectorLength {
		result += a[i].X*b[i].Y - a[i].Y*b[i].X
	}
	return dprec.Clamp(result, -1.0, 1.0)
}

// VectorDistance computes the hyperspace distance between two vectors.
func VectorDistance(a, b *Vector) float64 {
	var result float64
	for i := range VectorLength {
		result += dprec.Sqr(a[i].X - b[i].X)
		result += dprec.Sqr(a[i].Y - b[i].Y)
	}
	return dprec.Sqrt(result)
}

// Vector is a hyperdimensional vector representing a gesture path.
//
// While it lives in hyperspace, it is constructed from points in 2D space.
// And while some operations work in hyperspace, as math would define them,
// others are defined in terms of the 2D points that make up the vector.
//
// All of this is done to facilitate gesture recognition and can be best
// understood once the protractor algorithm paper is read.
type Vector [VectorLength]dprec.Vec2

// Centroid computes the centroid of the vector points within 3D space
// (not the vector space).
func (v *Vector) Centroid() dprec.Vec2 {
	var result dprec.Vec2
	for _, point := range v {
		result = dprec.Vec2Sum(result, point)
	}
	return dprec.Vec2Quot(result, float64(VectorLength))
}

// Length computes the hyperspace length of the vector (not the 3D length of
// the path constructed from the points).
func (v *Vector) Length() float64 {
	var result float64
	for _, point := range v {
		result += dprec.Sqr(point.X) + dprec.Sqr(point.Y)
	}
	return dprec.Sqrt(result)
}

// Translate mutates the vector by translating all points by the given offset.
func (v *Vector) Translate(offset dprec.Vec2) {
	for i := range v {
		v[i] = dprec.Vec2Sum(v[i], offset)
	}
}

// Rotate mutates the vector by rotating all points by the given angle.
func (v *Vector) Rotate(angle dprec.Angle) {
	cs := dprec.Cos(angle)
	sn := dprec.Sin(angle)
	for i, value := range v {
		v[i] = dprec.Vec2{
			X: value.X*cs - value.Y*sn,
			Y: value.X*sn + value.Y*cs,
		}
	}
}

// Scale mutates the vector by scaling all points by the given factor.
func (v *Vector) Scale(scale float64) {
	for i := range v {
		v[i] = dprec.Vec2Prod(v[i], scale)
	}
}
