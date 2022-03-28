package day03

import (
	"testing"

	"github.com/D-MOVE/aoc"
)

var example = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPowerConsumption(t *testing.T) {
	got := PowerConsumption(example)
	want := 198
	aoc.Report(t, got, want)
}

func TestLifeSupport(t *testing.T) {
	got := LifeSupport(example)
	want := 230
	aoc.Report(t, got, want)
}
