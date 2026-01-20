package gesture_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gesture"
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Vector", func() {

	Describe("CreateVector", func() {
		It("creates a vector by resampling the path", func() {
			path := gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(float64(gesture.VectorLength-1), -float64(gesture.VectorLength-1)),
			}
			vector := gesture.CreateVector(path)
			for i, point := range vector {
				Expect(point).To(dprectest.HaveVec2Coords(float64(i), float64(-i)))
			}
		})
	})

	Describe("VectorDot", func() {
		It("returns 1.0 when vectors are identical", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			dot := gesture.VectorDot(&vectorA, &vectorB)
			Expect(dot).To(dprectest.EqualFloat64(1.0))
		})

		It("returns -1.0 when vectors are opposite", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(-1.0, 0.0),
			})
			dot := gesture.VectorDot(&vectorA, &vectorB)
			Expect(dot).To(dprectest.EqualFloat64(-1.0))
		})
	})

	Describe("VectorCross", func() {
		It("returns 0.0 when vectors are identical", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			cross := gesture.VectorCross(&vectorA, &vectorB)
			Expect(cross).To(dprectest.EqualFloat64(0.0))
		})

		It("returns 1.0 when vectors are perpendicular in positive direction", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(0.0, 1.0),
			})
			cross := gesture.VectorCross(&vectorA, &vectorB)
			Expect(cross).To(dprectest.EqualFloat64(1.0))
		})

		It("returns -1.0 when vectors are perpendicular in negative direction", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(0.0, -1.0),
			})
			cross := gesture.VectorCross(&vectorA, &vectorB)
			Expect(cross).To(dprectest.EqualFloat64(-1.0))
		})
	})

	Describe("VectorDistance", func() {
		It("computes the correct distance between two vectors", func() {
			vectorA := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 0.0),
			})
			vectorA.Scale(1.0 / vectorA.Length())

			vectorB := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(0.0, 1.0),
			})
			vectorB.Scale(1.0 / vectorB.Length())

			distance := gesture.VectorDistance(&vectorA, &vectorB)
			Expect(distance).To(dprectest.EqualFloat64(dprec.Sqrt(2.0)))
		})
	})

	Describe("#Centroid", func() {
		It("computes the correct centroid of the vector points", func() {
			vector := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
			})
			centroid := vector.Centroid()
			Expect(centroid).To(dprectest.HaveVec2Coords(1.0, 1.0))
		})
	})

	Describe("#Length", func() {
		It("computes the correct length of the vector", func() {
			path := make(gesture.Path, gesture.VectorLength)
			for i := range path {
				switch i % 2 {
				case 0:
					path[i] = dprec.NewVec2(1.0, 0.0)
				case 1:
					path[i] = dprec.NewVec2(-1.0, 0.0)
				}
			}
			vector := gesture.CreateVector(path)
			length := vector.Length()
			Expect(length).To(dprectest.EqualFloat64(dprec.Sqrt(64.0)))
		})
	})

	Describe("#Translate", func() {
		It("translates all points by the given offset", func() {
			vector := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(1.0, 1.0),
			})
			vector.Translate(dprec.NewVec2(1.0, -1.0))
			Expect(vector[0]).To(dprectest.HaveVec2Coords(1.0, -1.0))
			Expect(vector[gesture.VectorLength-1]).To(dprectest.HaveVec2Coords(2.0, 0.0))
		})
	})

	Describe("#Rotate", func() {
		It("rotates all points by the given angle", func() {
			vector := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(63.0, 0.0),
			})
			vector.Rotate(dprec.Degrees(90.0))
			for i, point := range vector {
				Expect(point).To(dprectest.HaveVec2Coords(0.0, float64(i)))
			}
		})
	})

	Describe("#Scale", func() {
		It("scale all points by the given factor", func() {
			vector := gesture.CreateVector(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(63.0, 0.0),
			})
			vector.Scale(2.0)
			for i, point := range vector {
				Expect(point).To(dprectest.HaveVec2Coords(float64(i*2), 0.0))
			}
		})
	})
})
