package gesture_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gesture"
	"github.com/mokiat/gomath/dprec"
)

var _ = Describe("Pattern", func() {
	const matchEpsilon = 0.0001

	Describe("SimilarityDot", func() {
		It("returns 1.0 for identical patterns, even if rotated", func() {
			patternA := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
			})
			patternB := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
			})
			Expect(gesture.SimilarityDot(patternA, patternB)).To(BeNumerically("~", 1.0, matchEpsilon))
		})

		It("returns less than 1.0 for distinct patterns", func() {
			patternRect := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
			})
			patternTriangle := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(-2.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(0.0, 2.0),
			})
			Expect(gesture.SimilarityDot(patternRect, patternTriangle)).To(BeNumerically("<", 0.9))
		})
	})

	Describe("SimilarityDistance", func() {
		It("returns 1.0 for identical patterns, even if rotated", func() {
			patternA := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
			})
			patternB := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
			})
			Expect(gesture.SimilarityDistance(patternA, patternB)).To(BeNumerically("~", 1.0, matchEpsilon))
		})

		It("returns less than 1.0 for distinct patterns", func() {
			patternRect := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(2.0, 2.0),
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(0.0, 0.0),
			})
			patternTriangle := gesture.CreatePattern(gesture.Path{
				dprec.NewVec2(0.0, 2.0),
				dprec.NewVec2(-2.0, 0.0),
				dprec.NewVec2(2.0, 0.0),
				dprec.NewVec2(0.0, 2.0),
			})
			Expect(gesture.SimilarityDistance(patternRect, patternTriangle)).To(BeNumerically("<", 0.9))
		})
	})

})
