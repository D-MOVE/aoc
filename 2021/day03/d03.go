//https://adventofcode.com/2021/day/3

package day03

import (
	"strconv"
)

//PowerConsumption calculate submarine's power consumption
//determined by gamma * epsilon from diagnosticReport
func PowerConsumption(diagnosticReport []string) int {
	gamma := 0
	epsilon := 0
	n := len(diagnosticReport[0])
	for i := 0; i < n; i++ {
		zero := 0
		one := 0
		for j := 0; j < len(diagnosticReport); j++ {
			if diagnosticReport[j][i] == '0' {
				zero++
			} else {
				one++
			}
		}
		if one > zero {
			gamma = 2*gamma + 1
			epsilon = 2 * epsilon
		} else {
			gamma = 2 * gamma
			epsilon = 2*epsilon + 1
		}
	}
	return gamma * epsilon
}

//LifeSupport returns Life Support ratio
//determined by oxygen generator rating * CO2 scrubber rating
func LifeSupport(diagnosticReport []string) int {
	n := len(diagnosticReport[0])
	var ones, zeros, major, minor []string

	//oxygen generator rating
	major = diagnosticReport
	for i := 0; i < n; i++ {
		ones = []string{}
		zeros = []string{}
		if len(major) == 1 {
			break
		}
		for _, number := range major {
			if number[i] == '0' {
				zeros = append(zeros, number)
			} else {
				ones = append(ones, number)
			}
		}
		if len(ones) >= len(zeros) {
			major = ones
		} else {
			major = zeros
		}
	}
	oxygen, _ := strconv.ParseInt(major[0], 2, 64)

	// CO2 scrubber rating
	minor = diagnosticReport
	for i := 0; i < n; i++ {
		ones = []string{}
		zeros = []string{}
		if len(minor) == 1 {
			break
		}
		for _, number := range minor {
			if number[i] == '0' {
				zeros = append(zeros, number)
			} else {
				ones = append(ones, number)
			}
		}
		if len(ones) >= len(zeros) {
			minor = zeros
		} else {
			minor = ones
		}
	}
	co2, _ := strconv.ParseInt(minor[0], 2, 64)

	return int(oxygen * co2)
}
