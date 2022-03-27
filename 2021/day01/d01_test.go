package day01

import (
	"fmt"
	"testing"

	"github.com/D-MOVE/aoc"
)

var (
	example = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	myEx    = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{10, 7, 17, 33, 32, 9, 10, 2, 24, 8},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
)

func TestNumberOfIncreased(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := NumberOfIncreased(example)
		want := 7
		aoc.Report(t, got, want)
	})

	t.Run("my examples", func(t *testing.T) {
		wants := []int{9, 0, 4, 0}
		for i, want := range wants {
			got := NumberOfIncreased(myEx[i])
			aoc.Report(t, got, want)
		}
	})
}

func TestThreeSumArray(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		depths := ThreeSumArray(example)
		got := NumberOfIncreased(depths)
		want := 5
		aoc.Report(t, got, want)
	})
	t.Run("my examples", func(t *testing.T) {
		wants := []int{7, 0, 3, 0}
		for i, want := range wants {
			depths := ThreeSumArray(myEx[i])
			got := NumberOfIncreased(depths)
			aoc.Report(t, got, want)
		}
	})
}

func TestNumberOfIncreasedNInterval(t *testing.T) {
	//tests for the given example
	exampleCases := []struct {
		interval int
		want     int
	}{
		{1, 7},
		{2, 5},
		{3, 5},
		{9, 1},
		{10, 0},
	}
	for _, exampleCase := range exampleCases {
		interval := exampleCase.interval
		t.Run(fmt.Sprintf("example: int=%d", interval), func(t *testing.T) {
			got := NumberOfIncreasedNInterval(example, interval)
			want := exampleCase.want
			aoc.Report(t, got, want)
		})
	}

	//tests for my examples
	myExampleCases := []struct {
		interval int
		wants    []int
	}{
		{
			2, []int{8, 0, 5, 0},
		},
		{
			5, []int{5, 0, 1, 0},
		},
	}
	for _, myExampleCase := range myExampleCases {
		interval := myExampleCase.interval
		t.Run(fmt.Sprintf("my examples: int=%d", interval), func(t *testing.T) {
			for i, want := range myExampleCase.wants {
				got := NumberOfIncreasedNInterval(myEx[i], interval)
				aoc.Report(t, got, want)
			}
		})
	}
}
