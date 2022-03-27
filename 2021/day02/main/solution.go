//https://adventofcode.com/2021/day/2

package main

import (
	"fmt"

	"github.com/D-MOVE/aoc"
	"github.com/D-MOVE/aoc/2021/day02"
)

const lineNum = 1000

func main() {
	commands := aoc.LinesN(lineNum)
	//for 1st
	os := day02.NewOldSubmarine()
	os.MultiMove(commands)
	fmt.Println(os.Product())

	//for 2nd
	ns := day02.NewNewSubmarine()
	ns.MultiMove(commands)
	fmt.Println(ns.Product())
}
