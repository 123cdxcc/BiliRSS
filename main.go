package main

import (
	_ "bilibiliRSS/internal/db"
	"bilibiliRSS/server/route"
	"fmt"
	"github.com/gin-gonic/gin"
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

	/*
	   	2 6
	      30 39
	      15 29 42 1 44 1

	*/

	/*n := 0
	m := 0
	fmt.Scan(&n, &m)
	ns := make([]int, 0, n)
	ms := make([]int, 0, m)
	for j := 0; j < n; j++ {
		x := 0
		fmt.Scan(&x)
		ns = append(ns, x)
	}
	for j := 0; j < m; j++ {
		x := 0
		fmt.Scan(&x)
		ms = append(ms, x)
	}
	nSum := 0
	for i := range ns {
		nSum += ns[i]
	}
	//fmt.Println("nsum", nSum)
	stp := make(map[int][]int)
	var mSum int
	for i := range ms {
		mSum = 0
		arr := make([]int, 0, m)
		for j := i; j < len(ms); j++ {
			mSum += ms[j]
			arr = append(arr, ms[j])
			z := int(math.Abs(float64(mSum - nSum)))
			stp[z] = arr
			//fmt.Println("i:", i, "j:", j, "ms[j]:", ms[j], "z:", z, "stp[z]:", stp)
		}
	}

	min := ms[0]
	arr := []int{ms[0]}
	for k, v := range stp {
		if k < min {
			min = k
			arr = v
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}*/
}

/*func checkLogin(key string) {
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
*/
