package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for index := 2; index <= int(math.Sqrt(float64(num))); index++ {
		if num%index == 0 {
			return false
		}
	}
	return true
}

func main() {
	results := ReadFile("input.txt")
	var content string
	for i := 0; i < len(results); i++ {
		if IsPrime(results[i]) {
			content += strconv.Itoa(results[i]) + "      is prime\n"
		} else {
			content += strconv.Itoa(results[i]) + "      is not prime\n"
		}
	}
	WriteFile("output.txt", content)
}
