package gesture_test

import (
	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gesture"
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Path", func() {
	var path gesture.Path

	BeforeEach(func() {
		path = gesture.Path{
			dprec.NewVec2(0.0, 0.0),
			dprec.NewVec2(0.0, 3.0),
			dprec.NewVec2(4.0, 3.0),
		}
	})

	Describe("Length", func() {
		It("returns the total length of the path", func() {
			Expect(path.Length()).To(dprectest.EqualFloat64(7.0))
		})
	})

	Describe("ResampleIter", func() {
		const pointCount = 8

		It("yields the correct number of points", func() {
			points := slices.Collect(path.ResampleIter(pointCount))
			Expect(points).To(HaveLen(pointCount))
			Expect(points[0]).To(dprectest.HaveVec2Coords(0.0, 0.0))
			Expect(points[1]).To(dprectest.HaveVec2Coords(0.0, 1.0))
			Expect(points[2]).To(dprectest.HaveVec2Coords(0.0, 2.0))
			Expect(points[3]).To(dprectest.HaveVec2Coords(0.0, 3.0))
			Expect(points[4]).To(dprectest.HaveVec2Coords(1.0, 3.0))
			Expect(points[5]).To(dprectest.HaveVec2Coords(2.0, 3.0))
			Expect(points[6]).To(dprectest.HaveVec2Coords(3.0, 3.0))
			Expect(points[7]).To(dprectest.HaveVec2Coords(4.0, 3.0))
		})
	})
})
