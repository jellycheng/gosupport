package gosupport

import "sort"

// 获取map的所有key
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SortedMapKeys[K comparable, V any](m map[K]V, less func(a, b K) bool) []K {
	keys := GetMapKeys(m)
	// 排序
	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})

	return keys
}

// 带过滤条件的key
func FilterMapKeys[K comparable, V any](m map[K]V, filter func(K, V) bool) []K {
	if m == nil {
		return nil
	}
	keys := make([]K, 0)
	for k, v := range m {
		if filter(k, v) {
			keys = append(keys, k)
		}
	}
	return keys
}

// 检查key是否存在
func MapKeyExists[K comparable, V any](m map[K]V, key K) bool {
	if m == nil {
		return false
	}
	_, exists := m[key]
	return exists
}

// 多个map的合并，返回去重后的key
func GetMapUniqueKeys[K comparable, V any](maps ...map[K]V) []K {
	keySet := make(map[K]struct{})
	for _, m := range maps {
		for k := range m {
			keySet[k] = struct{}{}
		}
	}

	return GetMapKeys(keySet)
}
