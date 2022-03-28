//https://adventofcode.com/2021/day/1

package main

import (
	"fmt"

	"github.com/D-MOVE/aoc"
	"github.com/D-MOVE/aoc/2021/day01"
)

const lineNum = 2000

func main() {
	depths := aoc.IntsN(lineNum)

	//for 1st
	//fmt.Println(day01.NumberOfIncreasedNInterval(depths, 1))

	//for 2nd
	fmt.Println(day01.NumberOfIncreasedNInterval(depths, 3))
}
