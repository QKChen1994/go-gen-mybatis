package util

import "strings"

// MySQL 数据类型到 Go 数据类型的映射
var mysqlToGoType = map[string]string{
	"int":       "int",
	"tinyint":   "int8",
	"smallint":  "int16",
	"mediumint": "int32",
	"bigint":    "int64",
	"float":     "float32",
	"double":    "float64",
	"decimal":   "string", // 可以使用 string 处理，也可以自定义 Decimal 类型
	"char":      "string",
	"varchar":   "string",
	"text":      "string",
	"date":      "util.CustomTime",
	"datetime":  "util.CustomTime",
	"timestamp": "util.CustomTime",
	"time":      "util.CustomTime", // 或者使用 string
	"blob":      "[]byte",
}

// ConvertMySQLTypeToGoType 将 MySQL 字段类型转换为 Go 类型
func ConvertMySQLTypeToGoType(mysqlType string) string {
	// 去掉可能的长度限制，例如 int(11) -> int
	if idx := strings.Index(mysqlType, "("); idx != -1 {
		mysqlType = mysqlType[:idx]
	}
	if goType, ok := mysqlToGoType[mysqlType]; ok {
		return goType
	}
	return "interface{}" // 默认值，表示未知类型
}

// MySQL 数据类型到 Go 数据类型的映射
var mysqlToGoTypeXml = map[string]string{
	"int":       "int",
	"tinyint":   "int8",
	"smallint":  "int16",
	"mediumint": "int32",
	"bigint":    "int64",
	"float":     "float32",
	"double":    "float64",
	"decimal":   "string", // 可以使用 string 处理，也可以自定义 Decimal 类型
	"char":      "string",
	"varchar":   "string",
	"text":      "string",
	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "time.Time", // 或者使用 string
	"blob":      "[]byte",
}

// ConvertMySQLTypeToGoTypeXml 将 MySQL 字段类型转换为 Go xml类型
func ConvertMySQLTypeToGoTypeXml(mysqlType string) string {
	// 去掉可能的长度限制，例如 int(11) -> int
	if idx := strings.Index(mysqlType, "("); idx != -1 {
		mysqlType = mysqlType[:idx]
	}
	if goType, ok := mysqlToGoTypeXml[mysqlType]; ok {
		return goType
	}
	return "interface{}" // 默认值，表示未知类型
}
