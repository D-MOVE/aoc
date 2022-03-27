package day02

import (
	"testing"

	"github.com/D-MOVE/aoc"
)

var example = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestOldSubmarineMovement(t *testing.T) {
	t.Run("example movement for old", func(t *testing.T) {
		s := NewOldSubmarine()
		s.forward(5)
		s.down(5)
		s.forward(8)
		s.up(3)
		s.down(8)
		s.forward(2)
		want := 150
		got := s.Product()
		aoc.Report(t, got, want)
	})
}

func TestOldSubmarineCommand(t *testing.T) {
	t.Run("example command for old", func(t *testing.T) {
		s := NewOldSubmarine()
		s.MultiMove(example)
		want := 150
		got := s.Product()
		aoc.Report(t, got, want)
	})
	t.Run("wrong command for old", func(t *testing.T) {
		s := NewOldSubmarine()
		want := false
		got := move(s, "backword 10")
		aoc.Report(t, got, want)
		got = move(s, "up 1 2")
		aoc.Report(t, got, want)
		got = move(s, "down")
		aoc.Report(t, got, want)
	})
}

func TestNewSubmarineCommand(t *testing.T) {
	t.Run("example command for new", func(t *testing.T) {
		s := NewNewSubmarine()
		s.MultiMove(example)
		want := 900
		got := s.Product()
		aoc.Report(t, got, want)
	})
}
