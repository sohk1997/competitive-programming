package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the absolutePermutation function below.
func absolutePermutation(n int32, k int32) []int32 {
	result := make([]int32, n)
	mark := make([]bool, n)
	falseResult := []int32{-1}
	for i := int32(1); i <= n; i++ {
		index := i - 1
		if mark[i] {
			continue
		}
		if i+k > n {
			return falseResult

		}
		if mark[i+k] {
			return falseResult
		}
		result[index] = i + k
		result[index+k] = i
		mark[i] = true
		mark[i+k] = true
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout := os.Stdout

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		result := absolutePermutation(n, k)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
