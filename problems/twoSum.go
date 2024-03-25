package problems

func TwoSum(nums []int, target int) []int {
	var res []int
	visited := make(map[int]int)
	for index, val := range nums {
		_, ok := visited[target-val]
		if ok {
			res = append(res, index)
			res = append(res, visited[target-val])
		}
		visited[val] = index
	}
	return res
}
