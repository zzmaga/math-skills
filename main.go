package main

import (
	"bufio"
	"fmt"
	"math"
	"mathskills/mathskills"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage go run main.go data.txt")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid number in file:", scanner.Text())
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("No valid numbers found in file.")
		return
	}

	average := mathskills.CalculateAverage(numbers)
	median := mathskills.CalculateMedian(numbers)
	variance := mathskills.CalculateVariance(numbers)
	stdDev := mathskills.CalculateStandardDeviation(numbers)

	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
