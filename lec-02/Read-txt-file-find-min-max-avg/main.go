package main

import (
	"bufio"
	"fmt"
	"log"
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

func WriteFile(filename string, content string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func FindMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func FindMin(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func FindAverage(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum / len(arr)
}

func main() {
	results := ReadFile("input.txt")
	content := "Max: " + strconv.Itoa(FindMax(results)) +
		"\n" + "Min: " + strconv.Itoa(FindMin(results)) +
		"\n" + "Average: " + strconv.Itoa(FindAverage(results))
	WriteFile("output.txt", content)
}
