package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func WriteFile(filePath string, content string) {

	// 获取文件所在的目录
	directory := filepath.Dir(filePath)

	// 创建目录，如果已经存在则忽略错误
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Printf("创建目录失败: %s\n", err)
		return
	}

	// 创建并打开文件，如果文件存在则清空内容
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// 确保在函数结束时关闭文件
	defer file.Close()

	// 写入数据到文件
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to", filePath)
}

func WriteTemplateToFile(filePath string, templateFilePath string, data any) {

	// 获取文件所在的目录
	directory := filepath.Dir(filePath)

	// 创建目录，如果已经存在则忽略错误
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Printf("创建目录失败: %s\n", err)
		return
	}

	// 创建函数映射
	funcMap := template.FuncMap{
		"hasPrefix": hasPrefix, // 将函数添加到映射中
	}

	content, err := readFileToString(templateFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析模板文件
	tmpl, err := template.New("111").Funcs(funcMap).Parse(content)
	//tmpl, err := template.ParseFiles(templateFilePath)
	//tmpl = template.New("111")
	//tmpl, err = tmpl.ParseFiles(templateFilePath)

	tmpl.Name()
	//tmpl.Funcs(funcMap)
	//
	if err != nil {
		panic(err)
	}

	// 创建并打开文件，如果文件存在则清空内容
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// 确保在函数结束时关闭文件
	defer file.Close()

	// 执行模板并将数据写入文件
	err = tmpl.Execute(file, data) // 输出到文件
	if err != nil {
		panic(err)
	}

	fmt.Println("Data written to", filePath)
}

// 定义一个函数，用于检查字符串前缀
func hasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func readFileToString(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
