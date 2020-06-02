package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	device = "8513efac-09ea-3709-b214-95b366f1a185"
	url    = "https://api.coolapk.com/v6/main/indexV8?page=1"
)

func strToMD5(str string) string {
	bt := []byte(str)
	tmp := md5.Sum(bt)
	return fmt.Sprintf("%x", tmp)
}
func getAppToken() string {
	time := time.Now().Unix()
	timeHex := fmt.Sprintf("0x%x", time)
	str := strconv.FormatInt(time, 10)
	timeMd5 := strToMD5(str)
	tmp := []byte(fmt.Sprintf("token://com.coolapk.market/c67ef5943784d09750dcfbb31020f0ab?%s$%s&com.coolapk.market", timeMd5, device))
	tmpMd5 := strToMD5(base64.StdEncoding.EncodeToString(tmp))
	result := fmt.Sprintf("%s%s%s", tmpMd5, device, timeHex)
	return result
}

func test() {
	header := map[string]string{
		"User-Agent":       "Dalvik/2.1.0 (Linux; U; Android 9; MI 8 SE MIUI/9.5.9) (#Build; Xiaomi; MI 8 SE; PKQ1.181121.001; 9) +CoolMarket/9.2.2-1905301",
		"X-App-Id":         "com.coolapk.market",
		"X-Requested-With": "XMLHttpRequest",
		"X-Sdk-Int":        "28",
		"X-Sdk-Locale":     "zh-CN",
		"X-Api-Version":    "9",
		"X-App-Version":    "9.2.2",
		"X-App-Code":       "1903501",
		"X-App-Device":     "QRTBCOgkUTgsTat9WYphFI7kWbvFWaYByO1YjOCdjOxAjOxEkOFJjODlDI7ATNxMjM5MTOxcjMwAjN0AyOxEjNwgDNxITM2kDMzcTOgsTZzkTZlJ2MwUDNhJ2MyYzM",
		"Host":             "api.coolapk.com",
		"X-Dark-Mode":      "0",
		"X-App-Token":      getAppToken(),
	}
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	for key, value := range header {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)
	if err != nil {
		os.Exit(1)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func main() {
	fmt.Println(getAppToken())
	test()
}
