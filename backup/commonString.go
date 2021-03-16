//https://www.hackerrank.com/challenges/common-child/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var result [][]int32

func initResult(n, m int) {
	result = make([][]int32, n+2)
	for i := 0; i <= n; i++ {
		result[i] = make([]int32, m+2)
	}
}

func max(arr ...int32) int32 {
	result := int32(0)
	for i := range arr {
		if result < arr[i] {
			result = arr[i]
		}
	}
	return result
}

func boolToInt(value bool) int32 {
	if value {
		return int32(1)
	}
	return int32(0)
}

// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int32 {
	n, m := len(s1), len(s2)
	initResult(n, m)
	arr1, arr2 := []rune(s1), []rune(s2)
	for indexi := 0; indexi < n; indexi++ {
		i := indexi + 1
		for indexj := 0; indexj < m; indexj++ {
			j := indexj + 1
			result[i][j] = max(result[i-1][j], result[i][j-1], result[i-1][j-1]+boolToInt(arr1[indexi] == arr2[indexj]))
		}
	}
	return result[n][m]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout := os.Stdout
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
