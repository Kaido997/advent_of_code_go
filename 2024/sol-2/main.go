package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func parseInput(input []byte) ([]int, []int, error) {
	var output1 []int
	var output2 []int
	line := ""
	for _, v := range input {
		char := string(v)
		if char == "\n" {
			values := strings.Split(line, "   ")
			toInt0, err := strconv.ParseInt(values[0], 0, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("Int.0 strconv error %s", err)
			}
			toInt1, err := strconv.ParseInt(values[1], 0, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("Int.1 strconv error %s", err)
			}

			output1 = append(output1, int(toInt0))
			output2 = append(output2, int(toInt1))
			line = ""
			continue
		}
		line += char
	}
	return output1, output2, nil

}

func part1(input []byte) int {
	arr1, arr2, err := parseInput(input)
	if len(arr1) != len(arr2) {
		log.Fatal("input len not match")
	}
	if err != nil {
		log.Fatalf("%s", err)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	var totDistance int

	if len(arr1) <= 10 {
		fmt.Println(arr1, arr2)
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			totDistance += arr1[i] - arr2[i]
		} else {
			totDistance += arr2[i] - arr1[i]
		}
	}

	return totDistance
}

func part2(input []byte) int {
	arr1, arr2, err := parseInput(input)
	if len(arr1) != len(arr2) {
		log.Fatal("input len not match")
	}
	if err != nil {
		log.Fatalf("%s", err)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	var totSim int

	if len(arr1) <= 10 {
		fmt.Println(arr1, arr2)
	}

	for i := 0; i < len(arr1); i++ {
		idx, found := slices.BinarySearch(arr2, arr1[i])
		if found {
			var count int = 1
			for j := idx + 1; j < len(arr2); j++ {
				if arr2[j] != arr1[i] {
					break
				}
				count++
			}
			totSim += arr1[i] * count
		}
	}

	return totSim
}

func main() {
	log.Println(part1(input))
}
