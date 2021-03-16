package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMaxArea' function below.
 *
 * The function is expected to return a LONG_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER w
 *  2. INTEGER h
 *  3. BOOLEAN_ARRAY isVertical
 *  4. INTEGER_ARRAY distance
 */

type segmentTree struct {
	l, r, maxL, maxR, result int64
	leftChild, rightChild    *segmentTree
	bindL, bindR             bool
}

func max(args ...int64) int64 {
	result := int64(0)
	for _, value := range args {
		if result < value {
			result = value
		}
	}
	return result
}

func (tree *segmentTree) init(l, r int64) {
	tree.l, tree.r = l, r
	tree.bindL = true
	tree.bindR = true
	if l == r {
		tree.maxL = 1
		tree.maxR = 1
		tree.result = 1
		return
	}
	m := (l + r) / 2
	tree.leftChild = &segmentTree{}
	tree.rightChild = &segmentTree{}
	tree.leftChild.init(l, m)
	tree.rightChild.init(m+1, r)
	tree.maxL = tree.leftChild.maxL
	tree.maxR = tree.rightChild.maxR
	if tree.maxL == tree.leftChild.maxR {
		tree.maxL += tree.rightChild.maxL
	}
	if tree.maxR == tree.rightChild.maxL {
		tree.maxR += tree.leftChild.maxR
	}
	tree.result = max(tree.leftChild.result, tree.rightChild.result, tree.leftChild.maxR+tree.rightChild.maxL, tree.maxL, tree.maxR)
}

func (tree *segmentTree) update(index int64) {
	if tree.l > index+1 {
		return
	}
	if tree.r < index {
		return
	}
	if tree.l == index && tree.r == index {
		tree.bindR = false
		return
	}
	if tree.l == index+1 && tree.r == index+1 {
		tree.bindL = false
		return
	}
	tree.leftChild.update(index)
	tree.rightChild.update(index)
	tree.bindL = tree.leftChild.bindL
	tree.bindR = tree.rightChild.bindR

	tree.maxR = tree.rightChild.maxR
	if tree.rightChild.bindL && tree.leftChild.bindR {
		tree.maxR += tree.leftChild.maxR
	}

	tree.maxL = tree.leftChild.maxL

	if tree.leftChild.bindR && tree.rightChild.bindL {
		tree.maxL += tree.rightChild.maxL
	}

	tempResult := tree.leftChild.maxR + tree.rightChild.maxL
	if !tree.leftChild.bindR || !tree.rightChild.bindL {
		tempResult = max(tree.rightChild.maxL, tree.leftChild.maxR)
	}
	tree.result = max(tree.leftChild.result, tree.rightChild.result, tempResult, tree.maxL, tree.maxR)
}

func getMaxArea(w int32, h int32, isVertical []bool, distance []int32) []int64 {
	// Write your code here
	wTree, hTree := &segmentTree{}, &segmentTree{}
	wTree.init(1, int64(w))
	hTree.init(1, int64(h))
	// log.Println(wTree.maxL, wTree.maxR)
	// log.Println(hTree.maxL, hTree.maxR)
	result := make([]int64, len(distance))
	for index := range distance {
		if isVertical[index] {
			wTree.update(int64(distance[index]))
		} else {
			hTree.update(int64(distance[index]))
		}
		result[index] = hTree.result * wTree.result
		// log.Println(wTree.maxL, wTree.maxR)
		// log.Println(hTree.maxL, hTree.maxR)
	}
	return result
}

func main() {
	file, _ := os.Open("./inp")
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout := os.Stdout
	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	wTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	w := int32(wTemp)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	isVerticalCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var isVertical []bool

	for i := 0; i < int(isVerticalCount); i++ {
		isVerticalItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		isVerticalItem := isVerticalItemTemp != 0
		isVertical = append(isVertical, isVerticalItem)
	}

	distanceCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var distance []int32

	for i := 0; i < int(distanceCount); i++ {
		distanceItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		distanceItem := int32(distanceItemTemp)
		distance = append(distance, distanceItem)
	}

	result := getMaxArea(w, h, isVertical, distance)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
