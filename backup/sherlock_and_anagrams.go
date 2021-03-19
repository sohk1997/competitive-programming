//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	count := make([]int32, 30)
	runeArr := []rune(s)
	countMap := make(map[string]int32)
	for i := 0; i < len(runeArr); i++ {
		count[runeArr[i]-'a']++
		currentCount := make([]int32, 30)
		for j := 0; j <= i; j++ {
			if j > 0 {
				currentCount[runeArr[j-1]-'a']++
			}
			s := arrayToString(subArr(count, currentCount))
			countMap[s]++
		}
	}
	result := int32(0)
	for _, value := range countMap {
		result += (value * (value - 1)) / 2
	}
	return result
}

func arrayToString(arr []int32) string {
	result := ""
	for i := range arr {
		result += fmt.Sprintf("%3d", arr[i])
	}
	return result
}

func subArr(first []int32, second []int32) []int32 {
	result := make([]int32, len(first))
	for index := range first {
		result[index] = first[index] - second[index]
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
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
