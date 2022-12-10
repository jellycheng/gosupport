package utils

import "sort"

//对取出map的所有key，并按key排序
func KeysOfMap(m map[string]string) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	keys.Sort()
	return []string(keys)
}
