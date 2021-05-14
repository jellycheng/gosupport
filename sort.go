package gosupport

import (
	"sort"
)

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


// 对map的key进行排序，返回排序后的key切片
func KeysOfMap(m map[string]string) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	keys.Sort()
	return keys
}

type MyMapStringInt struct {
	Key   string
	Value int
}

type MyMapStringIntList []MyMapStringInt

func (p MyMapStringIntList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p MyMapStringIntList) Len() int           { return len(p) }
func (p MyMapStringIntList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// 对map的int值排序， ret := gosupport.MapSortByValue(map[string]int{"abc":123, "xyz":4}, "desc")
func MapSortByValue(m map[string]int, so string) MyMapStringIntList {
	p := make(MyMapStringIntList, 0)
	for k, v := range m {
		p = append(p, MyMapStringInt{Key:k, Value:v})
	}

	if so == "desc" {
		sort.Sort(sort.Reverse(p))
	} else {
		sort.Sort(p)
	}
	return p
}
