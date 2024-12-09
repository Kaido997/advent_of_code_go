package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte
var strNums []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func parseInputV2(input []byte) []string {
	return strings.Split(string(input), "mul")
}

func part2(input []byte) int {
	parsef := parseInputV2(input)
	var toMult [][2]int
	toMultIdx := 0
	parseNum := true
	for i, v := range parsef {
		b := []byte(v)
		prev := i - 1
		if i == 0 {
			prev = 0
		}
		if strings.Contains(parsef[prev], "don't()") {
			parseNum = false
		} else if strings.Contains(parsef[prev], "do()") {
			parseNum = true
		} else {
			// do nothing
		}
		if !parseNum {
			continue
		}
		if string(b[0]) == "(" {
			idx := 1
			num1 := ""
			num2 := ""
			var nextflag int = 0

			for string(b[idx]) != ")" {
				value := string(b[idx])

				if value == "," {
					nextflag = 1
					idx++
					continue
				}

				if value == " " {
					break
				} else {
					if slices.Contains(strNums, value) {
						if nextflag == 0 {
							num1 += value
						} else {
							num2 += value
						}
					} else {
						num1 = ""
						num2 = ""
						break
					}
				}
				idx++
			}
			if num1 == "" || num2 == "" {
				continue
			}

			iNum1, err := strconv.ParseInt(num1, 0, 64)
			if err != nil {
				panic(err)
			}
			iNum2, err := strconv.ParseInt(num2, 0, 64)
			if err != nil {
				panic(err)
			}
			toMult = append(toMult, [2]int{int(iNum1), int(iNum2)})
			toMultIdx++
		}
	}
	fmt.Printf("parsef: %v\n", parsef)
	fmt.Printf("toMult: %v\n", toMult)
	sum := 0

	for _, el := range toMult {
		sum += el[0] * el[1]
	}

	fmt.Printf("sum: %v\n", sum)
	return sum
}

func part1(input []byte) int {
	parsef := parseInputV2(input)
	var toMult [][2]int
	toMultIdx := 0
	for _, v := range parsef {
		b := []byte(v)
		if string(b[0]) == "(" {
			idx := 1
			num1 := ""
			num2 := ""
			var nextflag int = 0

			for string(b[idx]) != ")" {
				value := string(b[idx])

				if value == "," {
					nextflag = 1
					idx++
					continue
				}

				if value == " " {
					break
				} else {
					if slices.Contains(strNums, value) {
						if nextflag == 0 {
							num1 += value
						} else {
							num2 += value
						}
					} else {
						num1 = ""
						num2 = ""
						break
					}
				}
				idx++
			}
			if num1 == "" || num2 == "" {
				continue
			}
			iNum1, err := strconv.ParseInt(num1, 0, 64)
			if err != nil {
				panic(err)
			}
			iNum2, err := strconv.ParseInt(num2, 0, 64)
			if err != nil {
				panic(err)
			}
			toMult = append(toMult, [2]int{int(iNum1), int(iNum2)})
			toMultIdx++
		}
	}
	//fmt.Printf("parsef: %v\n", parsef)
	//fmt.Printf("toMult: %v\n", toMult)
	sum := 0

	for _, el := range toMult {
		sum += el[0] * el[1]
	}

	fmt.Printf("sum: %v\n", sum)
	return sum
}

//go:embed test_input.txt
var tst []byte

//go:embed test_input_2.txt
var tst2 []byte

func main() {
	part1(tst2)
	part2(tst2)
	os.Exit(0)
}
