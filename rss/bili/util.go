package bili

import (
	"io"
	"net/http"
)

var cookie = "SESSDATA=53d53e40%2C1696524299%2C44215%2A42; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT; HttpOnly; Securebili_jct=803b5cb90c05ca91e8683f89d8b1dba8; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID=23132528; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTDedeUserID__ckMd5=f315f0b03aa596b9; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMTsid=eytrhxcp; Path=/; Domain=bilibili.com; Expires=Thu, 05 Oct 2023 16:44:59 GMT"

func (b *bili) get(url string, header map[string]string) ([]byte, error) {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
