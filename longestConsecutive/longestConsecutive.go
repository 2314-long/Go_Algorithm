package main

func main() {
	// Example usage
	nums := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutive(nums)
	println(result) // Output: 4 (the longest consecutive sequence is [1, 2, 3, 4])
}

func longestConsecutive(nums []int) int {
	ans := 1
	map1 := make(map[int]bool)
	for _, num := range nums {
		map1[num] = true
	}
	for num := range map1 {
		cnt := 1
		if !map1[num-1] {
			for map1[num+1] {

				cnt++
				num = num + 1
				if cnt > ans {
					ans = cnt
				}

			}
		}
	}
	return ans
}

// max 函数返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
