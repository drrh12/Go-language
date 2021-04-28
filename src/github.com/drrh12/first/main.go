package main

import "fmt"

var result []int

func nTestCase(idx int, testCase int) {
	if idx == testCase {
		PrintN(result, 0)
		return
	}
	var n int
	fmt.Scanln(&n)

	all := make([]int, n)
	ReadN(all, 0, n) // reading array elements
	result = append(result, sumOfSquare(all, 0, n))
	nTestCase(idx+1, testCase)
}

func sumOfSquare(values []int, idx int, arrayLen int) int {
	if idx == arrayLen {
		return 0
	}
	if values[idx] > 0 { // ignoring negative numbers in the array
		return values[idx]*values[idx] + sumOfSquare(values, idx+1, arrayLen)
	}
	return sumOfSquare(values, idx+1, arrayLen)
}

func ReadN(all []int, idx, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&all[idx]); m != 1 {
		panic(err)
	}
	ReadN(all, idx+1, n-1)
}

func PrintN(values []int, idx int) {
	if len(values) == idx {
		return
	}
	fmt.Println(values[idx])
	PrintN(values, idx+1)
}

func main() {

	var testCaseNum int
	fmt.Scanln(&testCaseNum)

	nTestCase(0, testCaseNum)
}
