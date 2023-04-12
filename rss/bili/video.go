package bili

import (
	"bilibiliRSS/rss/bili/entity"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetDownloadUrl(cookie, aid, cid string, qn int) (string, error) {
	url := fmt.Sprintf("https://api.bilibili.com/x/player/playurl?platform=html5&avid=%s&cid=%s&qn=%d", aid, cid, qn)
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("cookie", cookie)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resp := new(entity.Resp[map[string]interface{}])
	err = json.Unmarshal(body, resp)
	if err != nil {
		return "", err
	}
	if resp.Code == 0 {
		durlAny, ok := resp.Data["durl"].([]interface{})
		if ok && len(durlAny) > 0 {
			durl, ok := durlAny[0].(map[string]interface{})
			if ok {
				return durl["url"].(string), nil
			}
		}
	}
	return "", errors.New(resp.Message)
}

func AidToCid(id string) (string, error) {
	url := "https://api.bilibili.com/x/player/pagelist?aid=" + id
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resp := new(entity.Resp[[]map[string]interface{}])
	err = json.Unmarshal(body, resp)
	if err != nil {
		return "", err
	}
	if resp.Code == 0 && len(resp.Data) > 0 {
		return strconv.Itoa(int(resp.Data[0]["cid"].(float64))), nil
	}
	return "", errors.New("转换失败")
}
