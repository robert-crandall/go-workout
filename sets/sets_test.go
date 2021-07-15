package sets

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSessions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sessions")
}

const (
	expectedRestTimeSeconds = 60
)

var _ = Describe("Workout Sets", func() {

	Context("static percentage", func() {

		staticList := NewSets(
			WithSetCount(5),
			WithRestTimer(expectedRestTimeSeconds),
			WithRepCount(5),
			WithWeightPercentage(0.7),
		)

		expectedSetList := []set{
			{5, 0.7},
			{5, 0.7},
			{5, 0.7},
			{5, 0.7},
			{5, 0.7},
		}

		It("should be a 5x5 at 70%", func() {
			staticList.Static()
			Expect(staticList.SetList).To(Equal(expectedSetList))
			Expect(staticList.RestTimeSeconds).To(Equal(expectedRestTimeSeconds))
		})
	})

	Context("RPT list percentage", func() {

		rptList := NewSets(
			WithSetCount(3),
			WithRestTimer(expectedRestTimeSeconds),
			WithRepCount(5),
			WithWeightPercentage(0.85),
		)

		expectedRPTList := []set{
			{5, 0.85},
			{7, 0.80},
			{9, 0.75},
		}

		It("should be a 3x5-7", func() {
			rptList.RPT(2, 0.05)
			Expect(rptList.SetList).To(Equal(expectedRPTList))
			Expect(rptList.RestTimeSeconds).To(Equal(expectedRestTimeSeconds))
		})
	})
})
