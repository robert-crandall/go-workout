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
				repsList:        expectedRepsList,
				weightsLBList:   nil,
				percentageList:  expectedPercentageList,
				lastSetsIsAMRAP: false,
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
				repsList:        expectedRepsList,
				weightsLBList:   nil,
				percentageList:  expectedPercentageList,
				lastSetsIsAMRAP: false,
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
				repsList:        expectedRepsList,
				weightsLBList:   nil,
				percentageList:  expectedPercentageList,
				lastSetsIsAMRAP: false,
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
				repsList:        expectedRepsList,
				weightsLBList:   nil,
				percentageList:  expectedPercentageList,
				lastSetsIsAMRAP: false,
			}

			It("should be a 3x5 starting at 70%", func() {
				Expect(session).To(Equal(expectedSets))
			})
		})
	})
})
