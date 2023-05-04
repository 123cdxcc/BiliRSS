package bili

import (
	"bilibiliRSS/internal/model"
	"bilibiliRSS/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
)

type bili struct {
	base string
}

var Bili = bili{
	base: "https://api.bilibili.com",
}

func (b *bili) ShareUrlToMid(shareUrl string) (int64, error) {
	request, err := http.NewRequest(http.MethodGet, shareUrl, nil)
	if err != nil {
		return 0, err
	}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	mid, err := strconv.ParseInt(response.Request.URL.Path[1:], 10, 64)
	if err != nil {
		return 0, err
	}
	return mid, nil
}
func (b *bili) SubscriptionInfo(mid int64) (*model.Subscription, error) {
	url := fmt.Sprintf("%s/x/space/acc/info?mid=%d", b.base, mid)
	res, err := b.get(url, map[string]string{"cookie": cookie})
	if err != nil {
		return nil, err
	}
	data := utils.Json.Get(res, "data")
	fmt.Println(data.Get("face").ToString())
	name := data.Get("name").ToString()
	pic := data.Get("face").ToString()
	return &model.Subscription{
		Name: name,
		Mid:  mid,
		Pic:  pic,
	}, nil
}
func (b *bili) VideoList(mid int64, page, count int) ([]*model.Video, error) {
	url := fmt.Sprintf("%s/x/space/wbi/arc/search?mid=%d&pn=%d&ps=%d", b.base, mid, page, count)
	res, err := b.get(url, nil)
	if err != nil {
		return nil, err
	}
	videos := make([]*model.Video, 0, count)
	vlist := utils.Json.Get(res, "data").Get("list").Get("vlist")
	for i := 0; i < vlist.Size(); i++ {
		v := vlist.Get(i)
		video := model.Video{
			Pic:         v.Get("pic").ToString(),
			Title:       v.Get("title").ToString(),
			Description: v.Get("description").ToString(),
			Author:      v.Get("author").ToString(),
			Length:      v.Get("length").ToString(),
			Aid:         v.Get("aid").ToString(),
			BvId:        v.Get("bvid").ToString(),
			CreateTime:  v.Get("created").ToInt(),
		}
		videos = append(videos, &video)
	}
	return videos, nil
}
func (b *bili) AidToCid(aid string) (string, error) {
	url := b.base + "/x/player/pagelist?aid=" + aid
	res, err := b.get(url, nil)
	if err != nil {
		return "", err
	}
	cid := utils.Json.Get(res, "data").Get(0).Get("cid").ToString()
	return cid, nil
}
func (b *bili) DownloadUrl(aid, cid string, qn int) (string, error) {
	url := fmt.Sprintf("%s/x/player/playurl?platform=html5&avid=%s&cid=%s&qn=%d", b.base, aid, cid, qn)
	res, err := b.get(url, map[string]string{"cookie": cookie})
	if err != nil {
		return "", err
	}
	durl := utils.Json.Get(res, "data").Get("durl")
	if durl.Size() == 0 {
		return "", nil
	}
	downloadUrl := durl.Get(0).Get("url").ToString()
	return downloadUrl, nil
}
