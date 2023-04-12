package model

type Video struct {
	Pic         string `json:"pic"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Length      string `json:"length"`
	Aid         string `json:"aid"`
	BvId        string `json:"bv_id"`
	Cid         string `json:"cid"`
	DownloadUrl string `json:"download_url"`
	CreateTime  int    `json:"create_time"`
}
