package gosupport

import "sort"

// 排名算法
// 示例： ret := gosupport.Float64Rank([]float64{-1,0,2.9,3,2.9}, "desc", false)
func Float64Rank(data []float64, order string, isRepeatOrder bool) map[float64]int {
	// 排名
	ret := map[float64]int{}
	newData := make([]float64, len(data))
	copy(newData, data)
	if order == "desc" {
		sort.Sort(sort.Reverse(sort.Float64Slice(newData)))
	} else if order == "asc" {
		sort.Float64s(newData)
	} else {
		return ret
	}
	rank := 1 // 排名
	temp := map[float64]struct{}{}
	for _, v := range newData {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			ret[v] = rank
			rank++
		} else if isRepeatOrder == false {// 不去重排序
			rank++
		}
	}
	return ret
}

