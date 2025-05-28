package mathskills

import (
	"math"
	"sort"
)

func CalculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func CalculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)

	if n%2 == 1 {
		return numbers[n/2]
	}
	return (numbers[n/2-1] + numbers[n/2]) / 2
}

func CalculateVariance(numbers []float64) float64 {
	avg := CalculateAverage(numbers)
	sum := 0.0
	for _, num := range numbers {
		sum += (num - avg) * (num - avg)
	}
	return sum / float64(len(numbers))
}

func CalculateStandardDeviation(numbers []float64) float64 {
	return math.Sqrt(CalculateVariance(numbers))
}
