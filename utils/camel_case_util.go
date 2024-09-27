package utils

import (
	"strings"
	"unicode"
)

// ToCamelCase 将下划线格式的字符串转换为驼峰格式
func ToCamelCase(s string) string {
	// 分割字符串
	parts := strings.Split(s, "_")
	for i := range parts {
		// 首字母大写，其他字母小写
		if len(parts[i]) > 0 {
			parts[i] = strings.Title(parts[i])
		}
	}
	// 拼接结果并返回
	return strings.Join(parts, "")
}

// toLowerFirstChar 将输入字符串的第一个字符转换为小写
func ToLowerFirstChar(s string) string {
	if len(s) == 0 {
		return s // 如果字符串为空，直接返回
	}

	// 获取第一个字符
	firstChar := s[0]

	// 将第一个字符转换为小写
	lowerFirstChar := byte(unicode.ToLower(rune(firstChar)))

	// 返回结果：将小写字符与剩余部分拼接
	return string(lowerFirstChar) + s[1:]
}
