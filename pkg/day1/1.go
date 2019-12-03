package day1

import (
	"github.com/archilkarchava/adventofcode2019/pkg/common"
	"strconv"
)

// CountFuel1 counts fuel for a single module not recursively
func CountFuel1(module uint64) uint64 {
	return module/3 - 2
}

// CountFuel2 counts fuel for a single module recursively
func CountFuel2(module uint64) uint64 {
	if int(module)/3-2 <= 0 {
		return 0
	}
	return module/3 - 2 + CountFuel2(module/3-2)
}

// Solve1 solves day 1 problem part 1
func Solve1() uint64 {
	data, err := common.ReadLines("./pkg/day1/input.txt")
	common.Check(err)
	var req1 uint64 = 0
	for _, s := range data {
		module, err := strconv.ParseUint(s, 10, 64)
		common.Check(err)
		req1 += CountFuel1(module)
	}
	return req1
}

// Solve2 solves day 1 problem part 2
func Solve2() uint64 {
	data, err := common.ReadLines("./pkg/day1/input.txt")
	common.Check(err)
	var req2 uint64 = 0
	for _, s := range data {
		module, err := strconv.ParseUint(s, 10, 64)
		common.Check(err)
		req2 += CountFuel2(module)
	}
	return req2
}
