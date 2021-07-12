package sets

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"main/lifts"
	"testing"
)

func TestSessions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sessions")
}

var _ = Describe("FSL531", func() {
	var (
		expectedRepsList       []int
		expectedPercentageList []float64
	)

	Describe("With a FSL531 session", func() {
		Context("First week", func() {
			session := FSL531(1)
			expectedRepsList = []int{5, 5, 5, 5, 5}
			expectedPercentageList = []float64{0.7, 0.7, 0.7, 0.7, 0.7}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 60,
			}

			It("should be a 5x5 at 70%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("Second week", func() {
			session := FSL531(2)
			expectedRepsList = []int{5, 5, 5, 5, 5}
			expectedPercentageList = []float64{0.725, 0.725, 0.725, 0.725, 0.725}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 60,
			}

			It("should be a 5x5 at 72.5%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})

	Describe("With a Main531 session", func() {
		Context("First week", func() {
			session := Main531(1)
			expectedRepsList = []int{5, 5, 5}
			expectedPercentageList = []float64{0.65, 0.75, 0.85}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be a 3x5 starting at 65%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("Second week", func() {
			session := Main531(2)
			expectedRepsList = []int{3, 3, 3}
			expectedPercentageList = []float64{0.70, 0.80, 0.90}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be a 3x5 starting at 70%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})

	Describe("With a StaticSets session", func() {
		Context("First week", func() {
			session := StaticSets(1, 5, 0.75)
			expectedRepsList = []int{5, 5, 5, 5, 5}
			expectedPercentageList = []float64{0.75, 0.75, 0.75, 0.75, 0.75}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be a 5x5 at 75%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("Fourth week", func() {
			session := StaticSets(4, 5, 0.75)
			expectedRepsList = []int{5, 5, 5, 5, 5}
			expectedPercentageList = []float64{0.65, 0.65, 0.65, 0.65, 0.65}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be a 5x5 at 65%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})

	Describe("With an RPTIncreaseWeight session", func() {
		Context("First week upper", func() {
			session := RPTIncreaseWeight(1, lifts.TargetUpper)
			expectedRepsList = []int{6, 8, 10}
			expectedPercentageList = []float64{0.80, 0.75, 0.70}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be an RPT set", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("First week lower", func() {
			session := RPTIncreaseWeight(1, lifts.TargetLower)
			expectedRepsList = []int{4, 5, 6}
			expectedPercentageList = []float64{0.85, 0.8, 0.75}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be an RPT set", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("Third week lower", func() {
			session := RPTIncreaseWeight(3, lifts.TargetLower)
			expectedRepsList = []int{4, 5, 6}
			expectedPercentageList = []float64{0.90, 0.85, 0.80}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should be an RPT set", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})

	Describe("With a StaticSetsIncreaseWeekly session", func() {
		Context("First week", func() {
			session := StaticSetsIncreaseWeekly(1, 3, 1)
			expectedRepsList = []int{5, 5, 5}
			expectedPercentageList = []float64{0.75, 0.75, 0.75}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should match expected result", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})

		Context("Third week at 90%", func() {
			session := StaticSetsIncreaseWeekly(1, 3, .9)
			expectedRepsList = []int{5, 5, 5}
			expectedPercentageList = []float64{0.675, 0.675, 0.675}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should match expected result", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})

	Describe("With a StaticSetsIncreaseReps session", func() {
		Context("First week", func() {
			session := StaticSetsIncreaseReps(3, 6, 3)
			expectedRepsList = []int{6, 9, 12}
			expectedPercentageList = []float64{1, 1, 1}

			expectedSets := Set{
				RepsList:        expectedRepsList,
				WeightsLBList:   nil,
				PercentageList:  expectedPercentageList,
				LastSetsIsAMRAP: false,
				RestTimeSeconds: 90,
			}

			It("should match expected result", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})
})
