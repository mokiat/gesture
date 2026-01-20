package gesture_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gesture"
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Segment", func() {
	var segment gesture.Segment

	BeforeEach(func() {
		segment = gesture.Segment{
			From: dprec.NewVec2(1.0, 2.0),
			To:   dprec.NewVec2(4.0, 6.0),
		}
	})

	Describe("Interpolation", func() {
		It("returns the correct interpolated point at t=0.0", func() {
			point := segment.Interpolation(0.0)
			Expect(point).To(dprectest.HaveVec2Coords(1.0, 2.0))
		})

		It("returns the correct interpolated point at t=1.0", func() {
			point := segment.Interpolation(1.0)
			Expect(point).To(dprectest.HaveVec2Coords(4.0, 6.0))
		})

		It("returns the correct interpolated point at t=0.5", func() {
			point := segment.Interpolation(0.5)
			Expect(point).To(dprectest.HaveVec2Coords(2.5, 4.0))
		})
	})

	Describe("Length", func() {
		It("returns the length of the segment", func() {
			Expect(segment.Length()).To(dprectest.EqualFloat64(5.0))
		})
	})
})
