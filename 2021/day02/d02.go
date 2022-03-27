package day02

import (
	"strconv"
	"strings"
)

const splitter = " "

//Mover is an interface for submarines
type Mover interface {
	forward(unit int)
	up(unit int)
	down(unit int)
}

//Move makes s move as a com indicates
func move(m Mover, com string) bool {
	splittedStr := strings.Split(com, splitter)
	if len(splittedStr) != 2 {
		return false
	}
	unit, err := strconv.Atoi(splittedStr[1])
	if err != nil {
		return false
	}
	switch splittedStr[0] {
	case "forward":
		m.forward(unit)
	case "up":
		m.up(unit)
	case "down":
		m.down(unit)
	default:
		return false
	}
	return true
}

//multiMove moves m multiple times according to coms
func multiMove(m Mover, coms []string) bool {
	ret := true
	for _, com := range coms {
		ret = ret && move(m, com)
	}
	return ret
}

//-----------------------------------------------

//OldSubmarine stores its potisin
type OldSubmarine struct {
	horizontalPos, depth int
}

//NewOldSubmarine serves factory function for OldSubmarine
func NewOldSubmarine() *OldSubmarine {
	return new(OldSubmarine)
}

//MultiMove moves s multiple times according to coms
func (s *OldSubmarine) MultiMove(coms []string) bool {
	return multiMove(s, coms)
}

func (s *OldSubmarine) forward(unit int) {
	s.horizontalPos += unit
}

func (s *OldSubmarine) up(unit int) {
	s.depth -= unit
}

func (s *OldSubmarine) down(unit int) {
	s.depth += unit
}

//Product returns the value of s.horizontalPos * s.depth
func (s OldSubmarine) Product() int {
	return s.horizontalPos * s.depth
}

//-----------------------------------------------

//NewSubmarine stores its potisin
type NewSubmarine struct {
	horizontalPos, depth, aim int
}

//NewNewSubmarine serves factory function for NewSubmarine
func NewNewSubmarine() *NewSubmarine {
	return new(NewSubmarine)
}

//MultiMove moves s multiple times according to coms
func (s *NewSubmarine) MultiMove(coms []string) bool {
	return multiMove(s, coms)
}

func (s *NewSubmarine) forward(unit int) {
	s.horizontalPos += unit
	s.depth += s.aim * unit
}

func (s *NewSubmarine) up(unit int) {
	s.aim -= unit
}

func (s *NewSubmarine) down(unit int) {
	s.aim += unit
}

//Product returns the value of s.horizontalPos * s.depth
func (s NewSubmarine) Product() int {
	return s.horizontalPos * s.depth
}
