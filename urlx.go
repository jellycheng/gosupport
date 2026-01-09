package gosupport

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// 提取给定 URL 字符串的 Path 部分（不含查询参数和片段）
func GetURLPath(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return u.Path, nil
}

// 返回 URL 中所有 query 参数名,并按字典序排序
func SortedQueryKeys(rawURL string) ([]string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(u.Query()))
	for k := range u.Query() { // Query() 返回 map[string][]string
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, nil
}

// 从完整URL中提取基础地址（不含参数）
func GetBaseURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("URL解析失败: %v", err)
	}

	// 构建基础URL
	baseURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path)
	return baseURL, nil
}

// 解析URL参数字符串为map, 格式 k1=值&k2=值2
func ParseURLParams(paramStr string) (map[string]string, error) {
	result := make(map[string]string)
	if len(paramStr) == 0 {
		return result, nil
	}
	// 如果包含完整的URL，先提取查询参数部分
	if strings.Contains(paramStr, "?") {
		parts := strings.Split(paramStr, "?")
		if len(parts) > 1 {
			paramStr = parts[1]
		}
	}

	// 解析参数
	params, err := url.ParseQuery(paramStr)
	if err != nil {
		return nil, fmt.Errorf("参数解析失败: %v", err)
	}

	for key, values := range params {
		if len(values) > 0 {
			result[key] = values[0] // 只取第一个值
		}
	}

	return result, nil
}
