package gesture

import "github.com/mokiat/gomath/dprec"

// CreatePattern creates a gesture pattern from the given path. Two gesture
// patterns can be compared for similarity after being created.
func CreatePattern(path Path) Pattern {
	vector := CreateVector(path)

	position := vector.Centroid()
	vector.Translate(dprec.InverseVec2(position))

	angle := dprec.Atan2(vector[0].Y, vector[0].X)
	vector.Rotate(-angle)

	scale := vector.Length()
	vector.Scale(1.0 / scale)

	return Pattern{
		vector:   vector,
		position: position,
		angle:    angle,
		scale:    scale,
	}
}

// Pattern represents a normalized gesture pattern that is ready for matching
// (i.e. similarity checks).
type Pattern struct {
	vector   Vector
	position dprec.Vec2
	angle    dprec.Angle
	scale    float64
}

// SimilarityDot computes the similarity between two patterns using
// the dot product of their vectors.
//
// The returned value is in the range [0.0, 1.0], where 1.0 means identical
// patterns and 0.0 means completely dissimilar patterns.
func SimilarityDot(a Pattern, b Pattern) float64 {
	aVector := &a.vector
	bVector := &b.vector
	alignVectors(aVector, bVector)

	dot := VectorDot(aVector, bVector)
	arcdot := dprec.Acos(dot).Radians()
	score := 1.0 - arcdot/dprec.Pi
	return dprec.Clamp(score, 0.0, 1.0)
}

// SimilarityDistance computes the similarity between two patterns using
// the distance between their vectors.
//
// The returned value is in the range [0.0, 1.0], where 1.0 means identical
// patterns and 0.0 means completely dissimilar patterns.
func SimilarityDistance(a Pattern, b Pattern) float64 {
	aVector := &a.vector
	bVector := &b.vector
	alignVectors(aVector, bVector)

	score := 1.0 - VectorDistance(aVector, bVector)/2.0
	return dprec.Clamp(score, 0.0, 1.0)
}

func alignVectors(a, b *Vector) {
	cross := VectorCross(a, b)
	dot := VectorDot(a, b)
	correctionAngle := dprec.Atan2(cross, dot)
	a.Rotate(correctionAngle)
}
