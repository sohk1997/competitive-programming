// https://leetcode.com/problems/median-of-two-sorted-arrays/submissions/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func findMedianSortedArrays(num1 []int, num2 []int) float64 {

	leftResult := findLeftResult(num1, num2)

	if leftResult == -1 {
		leftResult = findLeftResult(num2, num1)
		leftResult = num2[leftResult]
	} else {
		leftResult = num1[leftResult]
	}

	rightResult := findRightResult(num1, num2)
	if rightResult == -1 {
		rightResult = findRightResult(num2, num1)
		rightResult = num2[rightResult]
	} else {
		rightResult = num1[rightResult]
	}
	return (float64(leftResult) + float64(rightResult)) / 2
}

func findLeftResult(num1 []int, num2 []int) int {
	result := -1
	leftTarget := (len(num1)+len(num2)+1)/2 - 1
	l, r := 0, len(num1)-1
	for l <= r {
		m := (l + r) / 2
		countL := findLeft(num2, num1[m])
		leftEqual := findLeftEqual(num2, num1[m])
		rightEqual := findRightEqual(num2, num1[m])

		equal := 0

		if rightEqual > -1 {
			equal = rightEqual - leftEqual + 1
		}
		fmt.Println("------------------")
		fmt.Println(num1, num2)
		fmt.Println(countL, m, leftTarget, equal)
		fmt.Println("------------------")
		countL += m
		if leftTarget > countL && leftTarget-countL <= equal {
			countL = leftTarget
		}
		if leftTarget == countL {
			result = m
			break
		}
		if countL < leftTarget {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return result
}

func findRightResult(num1 []int, num2 []int) int {
	result := -1
	leftTarget := (len(num1)+len(num2)+2)/2 - 1
	l, r := 0, len(num1)-1
	for l <= r {
		m := (l + r) / 2
		countL := findLeft(num2, num1[m])
		leftEqual := findLeftEqual(num2, num1[m])
		rightEqual := findRightEqual(num2, num1[m])
		equal := 0
		if rightEqual > -1 {
			equal = rightEqual - leftEqual + 1
		}
		countL += m
		log.Println(countL, num1[m])
		if leftTarget > countL && leftTarget-countL <= equal {
			countL = leftTarget
		}
		if leftTarget == countL {
			result = m
			break
		}
		if countL < leftTarget {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return result
}

func findLeft(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	result := -1
	for l <= r {
		m := (l + r) / 2
		if arr[m] >= value {
			r = m - 1
		} else {
			l = m + 1
			result = m
		}
	}
	return result + 1
}

func findLeftEqual(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	result := -1
	for l <= r {
		m := (l + r) / 2
		if arr[m] < value {
			l = m + 1
		} else {
			r = m - 1
			if arr[m] == value {
				result = m
			}
		}
	}
	return result
}

func findRightEqual(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	result := -1
	for l <= r {
		m := (l + r + 1) / 2
		if arr[m] > value {
			r = m - 1
		} else {
			l = m + 1
			if arr[m] == value {
				result = m
			}
		}
	}
	return result
}

func main() {
	var n int
	scanf("%d\n", &n)
	num1 := make([]int, n)
	for index := 0; index < n; index++ {
		scanf("%d", &num1[index])
	}
	scanf("\n")
	scanf("%d\n", &n)
	num2 := make([]int, n)
	for index := 0; index < n; index++ {
		scanf("%d", &num2[index])
	}
	result := findMedianSortedArrays(num1, num2)
	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
