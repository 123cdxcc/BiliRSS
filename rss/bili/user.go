package bili

import (
	"bilibiliRSS/rss/bili/entity"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ShareToMid(url string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return response.Request.URL.Path[1:], nil
}

func VideoList(mid string, page, count int) ([]entity.Video, error) {
	cookie := "SESSDATA=53d53e40%2C1696524299%2C44215%2A42; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT; HttpOnly; Securebili_jct=803b5cb90c05ca91e8683f89d8b1dba8; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID=23132528; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID__ckMd5=f315f0b03aa596b9; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTsid=eytrhxcp; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT"
	url := fmt.Sprintf("https://api.bilibili.com/x/space/wbi/arc/search?mid=%s&pn=%d&ps=%d", mid, page, count)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := new(entity.Resp[entity.VideoInfo])
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code == 0 {
		videos := make([]entity.Video, count)
		for i, v := range resp.Data.List.Vlist {
			videos[i] = entity.Video{
				Pic:         v.Pic,
				Title:       v.Title,
				Description: v.Description,
				Author:      v.Author,
				Length:      v.Length,
				Aid:         fmt.Sprintf("%d", v.Aid),
				BvId:        v.Bvid,
				CreateTime:  v.Created,
			}
			cid, err := AidToCid(videos[i].Aid)
			if err != nil {
				continue
			}
			downloadUrl, err := GetDownloadUrl(cookie, videos[i].Aid, cid, 80)
			if err != nil {
				continue
			}
			videos[i].DownloadUrl = downloadUrl
		}
		return videos, nil
	}
	return nil, errors.New(resp.Message)
}
