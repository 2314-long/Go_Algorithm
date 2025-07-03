package main

import "fmt"

func towSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, item := range nums {

		if value, exits := m[target-item]; exits {
			if index != value {
				ans := []int{index, value}
				return ans
			}

		}
		m[item] = index
	}
	return []int{0, 0}
}
func main() {
	ans := towSum([]int{3, 3}, 6)
	fmt.Println(ans)

}
