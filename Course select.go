package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func httpDo(url, cookie string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", cookie)

	fmt.Println("请求中...")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return resp.Status, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.Status, err
	}

	fmt.Println(string(body))
	return resp.Status, nil
}

func main() {
	var str string
	fmt.Print("请先确定想选择几门课（只能为正整数）：")
	std := bufio.NewScanner(os.Stdin)
	std.Scan()
	str = std.Text()
	value, err := strconv.Atoi(str)
	if err != nil || value <= 0 {
		fmt.Println("类型错误或值错误")
		time.Sleep(time.Hour)
		return
	}

	url, cookie := make([]string, value), make([]string, value)
	for i := 0; i < value; i++ {
		fmt.Print("请输入url：")
		std.Scan()
		url[i] = std.Text()
		fmt.Print("请输入cookie：")
		std.Scan()
		cookie[i] = std.Text()
	}
	for {
		for i := 0; i < value; i++ {
			fmt.Println(httpDo(url[i], cookie[i]))
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second)
	}
}
