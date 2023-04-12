package main

import (
	"bilibiliRSS/rss/bili"
	"bilibiliRSS/server/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	/*qr, err := bili.ApplyQR()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = bili.CreateQR(qr.Url, "./qrcode.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("申请成功，请扫码登录")
	checkLogin(qr.QrcodeKey)*/

	r := gin.Default()
	route.InitRouter(r)
	err := r.Run(":45321")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func checkLogin(key string) {
	for {
		successInfo, err := bili.CheckLogin(key)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		switch successInfo.Code {
		case 0:
			fmt.Println(successInfo.Message)
			fmt.Println(successInfo.Cookie)
			return
		default:
			fmt.Println(successInfo.Message)
		}
		time.Sleep(2 * time.Second)
	}
}
