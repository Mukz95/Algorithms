package main

import "fmt"

func LinearSearch(input_array []int, target int) (int, []int) {
	temp_array := make([]int, 0)
	count := 0
	for _, value := range input_array {
		if value == target {
			temp_array = append(temp_array, value)
			count++

		}
	}
	return count, temp_array
}

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func CommonOccurencesTwoArrays(first_array []int, second_array []int) (int, []int) {
	temp_array := make([]int, 0)
	count := 0
	for i := 0; i < len(first_array); i++ {
		for j := 0; j < len(second_array); j++ {
			if first_array[i] == second_array[j] {
				count++
				removeIndex(second_array, j)
				temp_array = append(temp_array, first_array[i])
			}
		}
	}

	return count, temp_array
}

/*work in progress**/
func UniqueOccurencesUsingArrays(input_array []int) {
	temp_array := make([]int, 0)
	for i := 0; i < len(input_array); i++ {
		for j := 0; j < i; j++ {
			if input_array[i] == input_array[j] {
				fmt.Println(input_array[i])
			} else {
				temp_array = append(temp_array, input_array[i])
			}
		}

	}
	fmt.Println(temp_array)
}

func UniqueOccurencesUsingMap(input_array []int) {
	tempMap := make(map[int]bool)
	for _, index := range input_array {
		if _, ok := tempMap[index]; ok {
			fmt.Println(index, "is a duplicate")

		} else {
			tempMap[index] = true
		}
	}

	for key, _ := range tempMap {
		fmt.Println(key)
	}
}

func main() {

	//INITIALIZE ARRAY

	first_input := []int{48, 88, 3, 81, 92, 92, 44, 96}
	second_input := []int{96, 3, 92, 1, 3, 71}

	compare_count, compare_array := CommonOccurencesTwoArrays(first_input, second_input)
	count, array := LinearSearch(first_input, 92)

	UniqueOccurencesUsingArrays(compare_array)
	fmt.Println("occurences:", count)
	fmt.Println("list of occurences:", array)

	fmt.Println("occurences:", compare_count)
	fmt.Println("list of occurences:", compare_array)

	//fmt.Println("List Of Unique Occurences:", uniquearray)

	//	fmt.Sprint("count of occurences: %v, list of occurences: %v", compare_count, compare_array)

}
