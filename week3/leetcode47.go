package main

import "sort"

var (
	ans [][]int
	temp,arr []int
	visited []bool
)

func permuteUnique(nums []int) [][]int {
	ans = [][]int{}
	arr = nums
	visited = make([]bool, len(arr))
	sort.Ints(arr)
	dfs()
	return ans
}

func dfs() {
	if len(temp) == len(arr) {
		ans = append(ans, append([]int{}, temp...))
		return
	}

	for i:= 0; i < len(arr); i++ {
		if visited[i] || (i >0 && !visited[i - 1] && arr[i] == arr[i - 1]){
			continue
		}
		visited[i] = true
		temp = append(temp, arr[i])
		dfs()
		temp = temp[: len(temp) - 1]
		visited[i] = false
		//if i + 1 < len(arr) && arr[i] == arr[i + 1] {
		//	i++
		//}
	}
}

func main() {
	permuteUnique([]int{3, 3, 0, 3})
}
