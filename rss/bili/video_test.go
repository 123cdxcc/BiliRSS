package bili

import (
	"fmt"
	"testing"
)

func TestMid(t *testing.T) {
	mid, err := ShareToMid("https://b23.tv/FLlkm9j")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取mid成功")
	videos, err := VideoList(mid, 1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取投稿成功")
	cid, err := AidToCid(videos[0].Aid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("aid转cid成功")
	cookie := "SESSDATA=53d53e40%2C1696524299%2C44215%2A42; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT; HttpOnly; Securebili_jct=803b5cb90c05ca91e8683f89d8b1dba8; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID=23132528; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID__ckMd5=f315f0b03aa596b9; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTsid=eytrhxcp; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT"
	url, err := GetDownloadUrl(cookie, videos[0].Aid, cid, 80)
	if err != nil {
		fmt.Println("失败", err)
		return
	}
	fmt.Println("url", url)
	/*req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("referer", "https://api.bilibili.com")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54")
	fmt.Println("开始下载")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("下载失败")
		return
	}
	defer res.Body.Close()
	fmt.Println("下载成功")
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = os.WriteFile("./a.mp4", bytes, 0666)
	if err != nil {
		fmt.Println("写入失败")
		return
	}*/
}

//过膝袜 肉感的神b@其实我是我是[movie=100%*100%]https://v.scflover.cf/bili/BV1tY411z7JJ|http://i1.hdslb.com/bfs/archive/62658b5d39f7f0af0c56872f8d12f3fb1376e1fa.jpg[/movie]
