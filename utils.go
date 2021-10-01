package main

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"time"
)

// NetWorkStatus 检测网络系统调用速度最快
func NetWorkStatus() bool {
	cmd := exec.Command("ping", "www.baidu.com", "-c", "4", "-W", "5")
	t1 := time.Now()
	err := cmd.Run()
	fmt.Println("waist time :", time.Now().Sub(t1))
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println("Net Status , OK")
	}
	return true
}

// CheckServer 检测网络 TCP 实现速度稍慢
func CheckServer() {
	timeout := time.Duration(5 * time.Second)
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", "www.baidu.com:443", timeout)
	fmt.Println("waist time :", time.Now().Sub(t1))
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
		return
	}
	fmt.Println("tcp server is ok")
}

// CheckServerHttp 检测网络 HTTP 实现速度最慢
func CheckServerHttp() {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	t1 := time.Now()
	resp, err := client.Get("https://www.baidu.com")
	fmt.Println("waist time :", time.Now().Sub(t1))
	if err != nil {
		fmt.Println("client.Get, error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
