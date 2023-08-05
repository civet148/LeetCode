package main

import "fmt"

func main() {

	var nums = []int{3, 2, 4}
	var target = 6

	fmt.Printf("results [%+v] \n", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	var results []int
	for i, v := range nums {

		fmt.Printf("i=[%v] v=[%v] \n", i, v)
		for j := i + 1; j <= len(nums)-1; j++ {

			vv := nums[j]
			fmt.Printf(" j=[%v] vv=[%v] \n", j, vv)
			if v+vv == target {

				results = append(results, i)
				results = append(results, j)
				return results
			}
		}
	}

	return results
}
