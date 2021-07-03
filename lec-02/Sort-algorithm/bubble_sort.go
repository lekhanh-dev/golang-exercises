package main

func BubbleSort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr); i++ {
		for j := 0; j < len(newArr)-i-1; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j], newArr[j+1] = newArr[j+1], newArr[j]
			}
		}
	}
	return newArr
}
