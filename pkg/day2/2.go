package day2

import (
	"strconv"
	"strings"

	"github.com/archilkarchava/adventofcode2019/pkg/common"
)

func readArr() []int {
	dat, err := common.ReadLines("./pkg/day2/input.txt")
	common.Check(err)
	data := dat[0]
	strArr := strings.Split(data, ",")
	res := make([]int, len(strArr))
	for i, strNum := range strArr {
		res[i], err = strconv.Atoi(strNum)
		common.Check(err)
	}
	return res
}

func intComputer(arr []int) []int {
	resArr := make([]int, len(arr))
	copy(resArr, arr)
loop:
	for i, opcode := range resArr {
		if i%4 != 0 {
			continue
		}
		switch opcode {
		case 99:
			break loop
		case 1:
			resArr[resArr[i+3]] = resArr[resArr[i+1]] + resArr[resArr[i+2]]
			break
		case 2:
			resArr[resArr[i+3]] = resArr[resArr[i+1]] * resArr[resArr[i+2]]
			break
		default:
			panic("opcode can be either 1, 2 or 99")
		}
	}
	return resArr
}

// Solve1 solves day 2 part 1 problem
func Solve1() int {
	arr := readArr()
	arr[1] = 12
	arr[2] = 2
	return intComputer(arr)[0]
}

// Solve2 solves day 2 part 2 problem
func Solve2() int {
	desiredResult := 19690720
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			arr := readArr()
			arr[1] = i
			arr[2] = j
			if intComputer(arr)[0] == desiredResult {
				return 100*i + j
			}
		}
	}
	return -1
}
