//https://adventofcode.com/2021/day/3

package main

import (
	"fmt"

	"github.com/D-MOVE/aoc"
	"github.com/D-MOVE/aoc/2021/day03"
)

const lineNum = 1000

func main() {
	diagnosticReport := aoc.LinesN(lineNum)

	//for 1st
	fmt.Println(day03.PowerConsumption(diagnosticReport))

	//for 2nd
	fmt.Println(day03.LifeSupport(diagnosticReport))
}
