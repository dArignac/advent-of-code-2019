package day4

import (
	"strconv"
	"strings"
)

// ResolvePasswordCount resolves the password count by checking the rules for all numbers in the range.
func ResolvePasswordCount(part uint8) int {
	counter := 0
	for i := 256310; i <= 732736; i++ {
		if matchesRules(i, part) {
			counter++
		}
	}
	return counter
}

func matchesRules(number int, part uint8) bool {
	digits := numberToSlice(number)
	if part == 1 {
		return hasTwoAdjacentDigits(digits) && hasNeverDecreasingDigits(digits)
	} else if part == 2 {
		return hasTwoAdjacentDigitsAdvanced(digits) && hasNeverDecreasingDigits(digits)
	}
	return false
}

func hasNeverDecreasingDigits(digits []int) bool {
	for i, v := range digits {
		if i+1 == len(digits) {
			break
		}
		if digits[i+1] < v {
			return false
		}
	}
	return true
}

func hasTwoAdjacentDigits(digits []int) bool {
	for i, v := range digits {
		if i+1 == len(digits) {
			break
		}
		if v == digits[i+1] {
			return true
		}
	}
	return false
}

func hasTwoAdjacentDigitsAdvanced(digits []int) bool {
	results := make(map[int]uint)

	for i, v := range digits {
		if i+1 == len(digits) {
			break
		}
		if v == digits[i+1] {
			if results[v] == 0 {
				results[v] = 2
			} else {
				results[v]++
			}
		}
	}

	for _, v := range results {
		if v == 2 {
			return true
		}
	}

	return false
}

func numberToSlice(num int) []int {
	var result []int
	for _, v := range strings.Split(strconv.Itoa(num), "") {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		result = append(result, i)
	}
	return result
}
