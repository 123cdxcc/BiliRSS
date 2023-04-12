package bili

import (
	"bilibiliRSS/rss/bili/entity"
	"encoding/json"
	"errors"
	"github.com/skip2/go-qrcode"
	"io"
	"net/http"
	"strings"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

func CreateQR(context, fileName string) error {
	return qrcode.WriteFile(context, qrcode.Medium, 256, fileName)
}

func ApplyQR() (*entity.LoginQRInfo, error) {
	url := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := new(entity.Resp[entity.LoginQRInfo])
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code == 0 {
		return &resp.Data, nil
	}
	return nil, errors.New("申请失败")
}

func CheckLogin(qrKey string) (*entity.LoginSuccessInfo, error) {
	url := "https://passport.bilibili.com/x/passport-login/web/qrcode/poll?qrcode_key=" + qrKey
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := new(entity.Resp[entity.CheckLoginQR])
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%v", resp)
	if resp.Code == 0 {
		loginInfo := entity.LoginSuccessInfo{}
		switch resp.Data.Code {
		case 0:
			loginInfo.Code = 0
			loginInfo.Message = "登录成功"
			loginInfo.Cookie = strings.Join(response.Header.Values("Set-Cookie"), "")
			//fmt.Println("cookie:", response.Header)
		case 86038:
			loginInfo.Code = 86038
			loginInfo.Message = "二维码已失效"
		case 86090:
			loginInfo.Code = 86090
			loginInfo.Message = "二维码已扫码未确认"
		case 86101:
			loginInfo.Code = 86101
			loginInfo.Message = "未扫码"
		default:
			loginInfo.Code = -1
			loginInfo.Message = "未知"
		}
		return &loginInfo, nil
	}
	return nil, errors.New(resp.Message)
}
