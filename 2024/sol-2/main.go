package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func parseInput(input []byte) ([][]int, error) {
	var output [][]int
	line := ""
	for _, v := range input {
		char := string(v)
		if char == "\n" {
			values := strings.Split(line, " ")
			var arr []int
			for _, ints := range values {
				toInt0, err := strconv.ParseInt(ints, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("Error for input %v Parse Int error: %s", input, err)
				}
				arr = append(arr, int(toInt0))

			}
			output = append(output, arr)
			line = ""
			continue
		}
		line += char
	}
	return output, nil
}

func findProblems(report []int, excludeIndex int) []int {
	var problems []int
	order := -1 // ascending
	leng := len(report) - 1
	for i := 0; i < leng; i++ {
		diff := report[i] - report[i+1]

		if i == 0 {
			if diff > 0 {
				order = 1
			} else {
				order = -1
			}
		}

		if (diff > 0 && diff < 4) || (diff > -4 && diff < 0) {
			if order > 0 && diff > 0 {
				continue
			} else if order < 0 && diff < 0 {
				continue
			} else {
				idx := i
				if excludeIndex != -1 {
					if excludeIndex == i {
						continue
					}
				}
				if i+1 == leng {
					if excludeIndex == leng {
						continue
					}
					idx += 1
				}
				problems = append(problems, idx)
			}
		} else {
			idx := i
			if excludeIndex != -1 {
				if excludeIndex == i {
					continue
				}
			}
			if i+1 == leng {
				if excludeIndex == leng {
					continue
				}
				idx += 1
			}
			problems = append(problems, idx)
		}
	}
	return problems

}
func isSafe(problems []int) bool {
	if len(problems) == 0 {
		return true
	} else {
		return false
	}
}

func problemDampener(report []int, problemsIdx []int) bool {
	for _, id := range problemsIdx {
		tmpproblems := findProblems(report, id)
		if isSafe(tmpproblems) {
			return true
		} else {
			continue
		}
	}
	return false
}

func part2(input []byte) int {
	var safeReports int
	data, err := parseInput(input)
	if err != nil {
		log.Fatalf("Sol2 Input parsing error %s", err)
	}
	for _, report := range data {
		fmt.Println(report)
		problems := findProblems(report, -1)
		fmt.Println(report)
		if isSafe(problems) {
			if len(problems) > 0 {
				fmt.Printf("Report: %v Problems idx: %v\n", report, problems)
			}
			//fmt.Printf("Safe report idx: %d, report: %v\n", i, report)
			safeReports++
		} else {
			fmt.Println(report)
			if problemDampener(report, problems) {
				fmt.Println(report)
				if len(problems) > 0 {
					fmt.Printf("Report: %v Problems idx: %v\n", report, problems)
				}
				fmt.Println(report)
				//fmt.Printf("Safe report idx: %d, report: %v\n", i, report)
				safeReports++
			} else {
				continue
			}
		}
	}
	return safeReports
}

func part1(input []byte) int {
	var safeReports int
	data, err := parseInput(input)
	if err != nil {
		log.Fatalf("Sol2 Input parsing error %s", err)
	}
	var order int
	for _, report := range data {

		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]
			if i == 0 {
				if diff < 0 {
					order = 0
				} else {
					order = 1
				}
			}
			if diff < 0 {
				diff = diff * -1
			}
			if diff > 0 && diff < 4 {
				if order == 1 && report[i] > report[i+1] {
					if i+1 == len(report)-1 {
						safeReports++
					}
					continue
				} else if order == 0 && report[i] < report[i+1] {
					if i+1 == len(report)-1 {
						safeReports++
					}
					continue
				}
				break
			}
			break
		}
		order = 0
	}

	return safeReports
}

//go:embed test_input.txt
var tst []byte

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
	os.Exit(0)
}
