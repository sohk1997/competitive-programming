package main

//https://www.hackerrank.com/challenges/maximum-palindromes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// define type of segment tree
type SegmentTree struct {
	l, r                  int64
	count                 []int64
	leftChild, rightChild *SegmentTree
}

func (tree *SegmentTree) Initialize(charArr []rune, l, r int64) {
	tree.l = l
	tree.r = r
	tree.count = make([]int64, 27)

	if l != r {
		m := (l + r) / 2
		tree.leftChild = &SegmentTree{}
		tree.rightChild = &SegmentTree{}
		tree.leftChild.Initialize(charArr, l, m)
		tree.rightChild.Initialize(charArr, m+1, r)

		for index := 0; index < 26; index++ {
			tree.count[index] = tree.leftChild.count[index] + tree.rightChild.count[index]
		}

	} else {
		index := charArr[l] - 'a'
		tree.count[index] = 1
	}
}

func (tree *SegmentTree) query(l, r int64) *SegmentTree {
	var result SegmentTree
	result.count = make([]int64, 27)
	if tree.l >= l && tree.r <= r {
		return tree
	}

	if tree.r < l {
		return &result
	}
	if tree.l > r {
		return &result
	}
	result.count = make([]int64, 27)

	left := tree.leftChild.query(l, r)
	right := tree.rightChild.query(l, r)
	for index := 0; index < 26; index++ {
		result.count[index] = left.count[index] + right.count[index]
	}

	return &result
}

func combination(n, r int64) int64 {
	result := (permution[r] * permution[n-r]) % MODULE
	result = expon(result, MODULE-2)
	result = (result * permution[n]) % MODULE
	return result
}

func expon(a int64, ex int64) int64 {
	if ex == 1 {
		return a
	}
	temp := expon(a, ex/2)
	result := (temp * temp) % MODULE
	if ex%2 == 1 {
		result = (result * a) % MODULE
	}
	return result
}

var tree SegmentTree
var permution []int64

const MODULE = int64(1e9 + 7)

/*
 * Complete the 'initialize' function below.
 *
 * The function accepts STRING s as parameter.
 */

func initialize(s string) {
	// This function is called once before all queries.
	runeArr := []rune(s)
	initPer()
	tree.Initialize(runeArr, 0, int64(len(runeArr)-1))
}

func initPer() {
	permution = make([]int64, int(1e5)+10)
	permution[0] = 1
	for index := range permution {
		if index == 0 {
			continue
		}
		permution[index] = (int64(index) * permution[index-1]) % MODULE
	}
}

/*
 * Complete the 'answerQuery' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER l
 *  2. INTEGER r
 */

func answerQuery(l int32, r int32) int32 {
	// Return the answer for this query modulo 1000000007.
	t := tree.query(int64(l-1), int64(r-1))

	var (
		result int64 = 1
		sum    int64
		count  int64
	)
	for index := range t.count {
		sum += t.count[index] / 2
		count += t.count[index] % 2
	}
	for index := range t.count {
		result = (result * combination(sum, t.count[index]/2)) % MODULE
		sum = sum - t.count[index]/2
	}

	if count > 0 {
		result = (result * count) % MODULE
	}
	return int32(result)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout := os.Stdout
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	initialize(s)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		lTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		l := int32(lTemp)

		rTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		r := int32(rTemp)

		result := answerQuery(l, r)

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
