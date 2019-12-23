package day7

import (
	"fmt"
	"sync"

	"github.com/dArignac/advent-of-code-2019/day5-sunny-with-a-chance-of-asteroids/day5"
	"github.com/dArignac/advent-of-code-2019/helper"
)

// Phases defines the phase settings for the amplifiers
type Phases struct {
	a             int
	b             int
	c             int
	d             int
	e             int
	thrusterValue int
}

// ToString returns a string representation of a Phases instance
func (s Phases) ToString() string {
	return fmt.Sprintf("%d%d%d%d%d max thruster signal: %d", s.a, s.b, s.c, s.d, s.e, s.thrusterValue)
}

// FindLargestOutputSignalForThrusters finds the largest thruster output for the given code
func FindLargestOutputSignalForThrusters(codeString string) Phases {
	// convert code for usage
	code, err := helper.SplitStringToIntArray(codeString)
	if err != nil {
		fmt.Println("Failed to convert program code")
	}

	// generate permutations
	sequences := generatePermutations([]int{0, 1, 2, 3, 4})

	// for each combination start a go subroutine and write result to list
	var wg sync.WaitGroup
	wg.Add(len(sequences))
	results := []Phases{}
	for _, seq := range sequences {
		go generateThrusterOutput(code, seq, &wg, &results)
	}
	wg.Wait()

	// find highest thruster value
	finalSequence := results[0]
	for _, res := range results {
		if res.thrusterValue > finalSequence.thrusterValue {
			finalSequence = res
		}
	}

	return finalSequence
}

func generateThrusterOutput(code []int, phases []int, wg *sync.WaitGroup, results *[]Phases) {
	defer (*wg).Done()

	getImmutableCodeCopy := func(code *[]int) []int {
		zeCopy := make([]int, len(*code))
		copy(zeCopy, *code)
		return zeCopy
	}

	// amp a
	outputs := []int{}
	inputs := []int{phases[0], 0}
	day5.RunProgramCode(getImmutableCodeCopy(&code), &inputs, 0, &outputs)

	// amp b
	inputs = []int{phases[1], outputs[0]}
	outputs = []int{}
	day5.RunProgramCode(getImmutableCodeCopy(&code), &inputs, 0, &outputs)

	// amp c
	inputs = []int{phases[2], outputs[0]}
	outputs = []int{}
	day5.RunProgramCode(getImmutableCodeCopy(&code), &inputs, 0, &outputs)

	// amp d
	inputs = []int{phases[3], outputs[0]}
	outputs = []int{}
	day5.RunProgramCode(getImmutableCodeCopy(&code), &inputs, 0, &outputs)

	// amp e
	inputs = []int{phases[4], outputs[0]}
	outputs = []int{}
	day5.RunProgramCode(getImmutableCodeCopy(&code), &inputs, 0, &outputs)

	// write result
	s := Phases{
		a:             phases[0],
		b:             phases[1],
		c:             phases[2],
		d:             phases[3],
		e:             phases[4],
		thrusterValue: outputs[0],
	}
	*results = append(*results, s)
}

// code from https://stackoverflow.com/a/30226442/364244
func generatePermutations(alphabet []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(alphabet, len(alphabet))

	return res
}

func runAmplificationCircuit(code []int, in Phases) Phases {
	// FIXME implements
	return Phases{0, 0, 0, 0, 0, 0}
}
