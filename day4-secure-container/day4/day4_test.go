package day4

import "testing"

import "github.com/stretchr/testify/assert"

func TestMatchesRules(t *testing.T) {
	assert.True(t, matchesRules(111111, 1))
	assert.False(t, matchesRules(223450, 1))
	assert.False(t, matchesRules(123789, 1))
	assert.True(t, matchesRules(112233, 2))
	assert.False(t, matchesRules(123444, 2))
	assert.True(t, matchesRules(111122, 2))
}

func TestNumberToSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, numberToSlice(123))
	assert.Equal(t, []int{7}, numberToSlice(7))
	assert.Equal(t, []int{7, 7, 8, 8, 9, 9}, numberToSlice(778899))
}

func TestHasTwoAdjacentDigits(t *testing.T) {
	assert.True(t, hasTwoAdjacentDigits([]int{1, 2, 2, 3, 4, 5}))
	assert.True(t, hasTwoAdjacentDigits([]int{1, 2, 5, 3, 5, 5}))
	assert.True(t, hasTwoAdjacentDigits([]int{1, 1, 3, 4, 5, 6}))
	assert.True(t, hasTwoAdjacentDigits([]int{1, 1, 3, 4, 4, 6}))
	assert.False(t, hasTwoAdjacentDigits([]int{1, 2, 3, 4, 5, 6}))
}

func TestHasTwoAdjacentDigitsAdvanced(t *testing.T) {
	assert.True(t, hasTwoAdjacentDigitsAdvanced([]int{1, 2, 2, 3, 4, 5}))
	assert.True(t, hasTwoAdjacentDigitsAdvanced([]int{1, 2, 5, 3, 5, 5}))
	assert.False(t, hasTwoAdjacentDigitsAdvanced([]int{1, 1, 1, 4, 5, 6}))
	assert.True(t, hasTwoAdjacentDigitsAdvanced([]int{1, 1, 3, 4, 4, 4}))
	assert.False(t, hasTwoAdjacentDigitsAdvanced([]int{1, 2, 3, 4, 5, 6}))
}

func TestHasNeverDecreasingDigits(t *testing.T) {
	assert.True(t, hasNeverDecreasingDigits([]int{1, 2, 3, 4, 5, 6}))
	assert.True(t, hasNeverDecreasingDigits([]int{1, 1, 3, 4, 4, 4}))
	assert.False(t, hasNeverDecreasingDigits([]int{1, 2, 3, 4, 5, 0}))
	assert.False(t, hasNeverDecreasingDigits([]int{9, 8, 7, 6, 5, 4}))
}
