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

func KeysOfMapV2(m map[string]interface{}) []string {
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

//	myInt64 := []int64{9, 8,10,7}
//	sort.Sort(gosupport.Int64Slice(myInt64)) // 升序，或者 gosupport.Int64s(myInt64) 或者 gosupport.Int64Slice(myInt64).Sort()
//	fmt.Println(myInt64) //[7 8 9 10]
type Int64Slice []int64
func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x Int64Slice) Sort() { sort.Sort(x) }
// 升序
func Int64s(a []int64)  {
	sort.Sort(Int64Slice(a))
}
// 倒序
func ReverseInt64s(a []int64) {
	sort.Sort(sort.Reverse(Int64Slice(a)))
}


