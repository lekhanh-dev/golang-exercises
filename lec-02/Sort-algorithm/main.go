package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filename string) []int {
	f, err := os.Open(filename)
	results := make([]int, 0)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		results = append(results, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return results
}

func main() {
	results := ReadFile("input.txt")
	// fmt.Println(results)
	// arrayByBubbleSort := BubbleSort(results)
	// fmt.Println(arrayByBubbleSort)
	// fmt.Println(results)

	// arrayByMergeSort := MergeSort(results)
	// fmt.Println(arrayByMergeSort)

	arrayByQuickSort := QuickSort(results)
	fmt.Println(arrayByQuickSort)
}
