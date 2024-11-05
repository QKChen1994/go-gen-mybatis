package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func Test_str(t *testing.T) {
	template := "My name is {{2}} and I am {{1}} years old."

	result, err := ReplacePlaceholders(template, "30", "Alice")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}

// ReplacePlaceholders 根据占位符替换字符串中的变量
func ReplacePlaceholders(template string, values ...string) (string, error) {
	re := regexp.MustCompile(`\{\{(\d+)\}\}`)
	result := re.ReplaceAllStringFunc(template, func(match string) string {
		// 获取占位符的数字
		index := match[2] - '1' // match[2] 是数字字符
		if int(index) < len(values) {
			return values[int(index)]
		}
		return match // 如果没有匹配，保持原样
	})
	return result, nil
}

func Test_dll(t *testing.T) {
	list := []map[string]any{
		{
			"filePath": "D:\\智云\\查打一体\\dll\\machine_model_11\\shanghai_test.dll",
			"modelId":  1788482580825473026,
		},
		{
			"filePath": "D:\\智云\\查打一体\\dll\\machine_model_11\\shanghai_test1.dll",
			"modelId":  1790978190155509762,
		},
		{
			"filePath": "D:\\智云\\查打一体\\dll\\machine_model_11\\shanghai_test2.dll",
			"modelId":  1803392800023719937,
		},
		{
			"filePath": "D:\\智云\\查打一体\\dll\\machine_model_11\\shanghai_test3.dll",
			"modelId":  1803240739950428162,
		},
		{
			"filePath": "D:\\智云\\查打一体\\dll\\machine_model_11\\shanghai_test4.dll",
			"modelId":  1817853714863951873,
		},
	}

	for _, item := range list {
		go callHttp(item["filePath"].(string), item["modelId"].(int))
	}
	time.Sleep(1000 * time.Second)
}

func callHttp(filePath string, modelId int) {
	url := "http://127.0.0.1:55822/stflowrenode/uploadDllFile"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filePath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(filePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("machine_id", strconv.Itoa(modelId))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", " QEtAbkFmQKPl9Pnsk7RmOM6OwJpGHkb8dIflwVKFeiLERhb7jmBb/vsn62WuDLOEVo8bErnejdTgut3ev/vdsa0W7lUWHW2pkhAeSXG4wFyi/Ecq/fZL+oV16kmuVNtz7Q2D/iji3XLcNRtYmLOzBUCPO7qjc5ieOM9X2EpcQevp+GHz+GBdQcn170dukNFzllrfpX1rNZGuJUOI9ziTdg==")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("modelId:%v,调用报错：%v", modelId, err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("modelId:%v,转换报错：%v", modelId, err)
		return
	}
	fmt.Printf("modelId:%v,返回：%v \n", modelId, string(body))
}
