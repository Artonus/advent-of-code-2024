package day01

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getLocationLists(fileName string) (locations1, locations2 []int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	locations1, locations2 = make([]int, 0), make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, "   ")
		val0, _ := strconv.Atoi(vals[0])
		val1, _ := strconv.Atoi(vals[1])
		locations1 = append(locations1, val0)
		locations2 = append(locations2, val1)
	}
	return locations1, locations2
}

func getDistances(locations1, locations2 []int) []int {
	distances := make([]int, len(locations1))
	if len(locations1) != len(locations2) {
		panic("Locations have different lengths")
	}
	length := len(locations1)
	for i := 0; i < length; i++ {
		distances[i] = int(math.Abs(float64(locations1[i] - locations2[i])))
	}
	return distances
}

func Day01(step1 bool) {
	fileName := "day01/data.txt"
	locations1, locations2 := getLocationLists(fileName)
	slices.Sort(locations1)
	if step1 == false {
		slices.Sort(locations2)
		distances := getDistances(locations1, locations2)
		sum := 0
		for _, val := range distances {
			sum += val
		}
		print(sum)
		return
	}
	sum := 0
	numberCounts := getNumberCounts(locations2)
	for _, val := range locations1 {
		count, _ := numberCounts[val]
		sum += val * count
	}
	print(sum)
}

func getNumberCounts(numbers []int) map[int]int {
	numberCounts := make(map[int]int)

	for _, number := range numbers {
		//_, exists := numberCounts[number]
		numberCounts[number]++
	}
	return numberCounts
}
