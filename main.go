package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the steadyGene function below.
func steadyGene(gene string) int32 {
	arr := []rune(gene)
	n := len(gene)
	count := make([]int, 30)
	for index := range arr {
		x := arr[index] - 'A'
		count[x]++
	}
	for index := range count {
		if count[index] > n/4 {
			count[index] -= n / 4
		} else {
			count[index] = 0
		}
	}
	l, r := 0, 0
	current := make([]int, 30)
	current[int(arr[0]-'A')] = 1
	result := -1
	for r < n-1 {
		r++
		current[int(arr[r]-'A')]++
		for l < n && checkEqual(current, count) {
			if result > r-l+1 || result == -1 {
				result = r - l + 1
			}
			current[arr[l]-'A']--
			l++
		}

	}
	return int32(result)
}

func checkEqual(check, target []int) bool {
	for index := range check {
		if check[index] < target[index] {
			return false
		}
	}
	return true
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	_ = int32(nTemp)

	gene := readLine(reader)

	result := steadyGene(gene)

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
