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

func main() {
	results := ReadFile("input.txt")
	var listValueCheck [5]int = [5]int{33, 1, 54, 0, 222}
	var content string
	for i := 0; i < len(listValueCheck); i++ {
		var check bool = false
		for j := 0; j < len(results); j++ {
			if listValueCheck[i] == results[j] {
				content += strconv.Itoa(listValueCheck[i]) + "     is exists\n"
				check = true
				break
			}
		}
		if !check {
			content += strconv.Itoa(listValueCheck[i]) + "     is not exists\n"
		}
	}
	WriteFile("output.txt", content)
}
