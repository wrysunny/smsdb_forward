package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Pushresp struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func send(ReceivingDateTime string, SenderNumber string, TextDecoded string) bool {
	postdata := fmt.Sprintf("接收时间：%s\n发送者：%s\n短信内容：%s\n", ReceivingDateTime, SenderNumber, TextDecoded)
	content := "title=手机号2570收到一条短信 时间： " + time.Now().Format("01-02 15:04:05") + "&content=" + postdata

	var pushurl = "https://push.showdoc.com.cn/server/api/push/3b"
	defaultCipherSuites := []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009,
		0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}
	httpclient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{CipherSuites: append(defaultCipherSuites[8:], defaultCipherSuites[:8]...)}},
	}
	req, _ := http.NewRequest("POST", pushurl, bytes.NewBufferString(content))
	// add header value
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("langType", "zh_CN")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) weapp/4.3.21/public//1.0//2")

	resp, _ := httpclient.Do(req)
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("read push resp error:", err.Error())
	}

	var pushresp Pushresp
	err = json.Unmarshal(all, &pushresp)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("sign resp content unmarshal error:", err.Error())
	}
	if pushresp.ErrorCode == 0 || pushresp.ErrorMessage == "ok" {
		log.SetPrefix("[Info] ")
		log.Println("push success.")
		send2(ReceivingDateTime, SenderNumber, TextDecoded)
		return true
	} else {
		log.SetPrefix("[Info] ")
		log.Println("push failed.")
		send2(ReceivingDateTime, SenderNumber, TextDecoded)
		return false
	}
}

func send2(ReceivingDateTime string, SenderNumber string, TextDecoded string) {
	postdata := fmt.Sprintf("接收时间：%s\n发送者：%s\n短信内容：%s\n", ReceivingDateTime, SenderNumber, TextDecoded)
	content := "title=手机号2570收到一条短信 时间： " + time.Now().Format("01-02 15:04:05") + "&content=" + postdata

	var pushurl = "https://push.showdoc.com.cn/server/api/push/a9"
	defaultCipherSuites := []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009,
		0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}
	httpclient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{CipherSuites: append(defaultCipherSuites[8:], defaultCipherSuites[:8]...)}},
	}
	req, _ := http.NewRequest("POST", pushurl, bytes.NewBufferString(content))
	// add header value
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("langType", "zh_CN")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) weapp/4.3.21/public//1.0//2")

	resp, _ := httpclient.Do(req)
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("read push resp error:", err.Error())
	}

	var pushresp Pushresp
	err = json.Unmarshal(all, &pushresp)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Println("sign resp content unmarshal error:", err.Error())
	}
	if pushresp.ErrorCode == 0 || pushresp.ErrorMessage == "ok" {
		log.SetPrefix("[Info] ")
		log.Println("push success.")
	} else {
		log.SetPrefix("[Info] ")
		log.Println("push failed.")
	}
}
