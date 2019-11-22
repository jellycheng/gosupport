package gosupport

//求和
func IntSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}



