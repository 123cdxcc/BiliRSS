package entity

type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    T      `json:"data"`
}
type LoginQRInfo struct {
	Url       string `json:"url"`
	QrcodeKey string `json:"qrcode_key"`
}

type CheckLoginQR struct {
	Url          string `json:"url"`
	RefreshToken string `json:"refresh_token"`
	Timestamp    int64  `json:"timestamp"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

type LoginSuccessInfo struct {
	Code    int
	Message string
	Cookie  string
}
type VideoInfo struct {
	List struct {
		Vlist []struct {
			Comment        int    `json:"comment"`
			Typeid         int    `json:"typeid"`
			Play           int    `json:"play"`
			Pic            string `json:"pic"`
			Subtitle       string `json:"subtitle"`
			Description    string `json:"description"`
			Copyright      string `json:"copyright"`
			Title          string `json:"title"`
			Review         int    `json:"review"`
			Author         string `json:"author"`
			Mid            int    `json:"mid"`
			Created        int    `json:"created"`
			Length         string `json:"length"`
			VideoReview    int    `json:"video_review"`
			Aid            int    `json:"aid"`
			Bvid           string `json:"bvid"`
			HideClick      bool   `json:"hide_click"`
			IsPay          int    `json:"is_pay"`
			IsUnionVideo   int    `json:"is_union_video"`
			IsSteinsGate   int    `json:"is_steins_gate"`
			IsLivePlayback int    `json:"is_live_playback"`
			Meta           *struct {
				Id        int    `json:"id"`
				Title     string `json:"title"`
				Cover     string `json:"cover"`
				Mid       int    `json:"mid"`
				Intro     string `json:"intro"`
				SignState int    `json:"sign_state"`
				Attribute int    `json:"attribute"`
				Stat      struct {
					SeasonId int `json:"season_id"`
					View     int `json:"view"`
					Danmaku  int `json:"danmaku"`
					Reply    int `json:"reply"`
					Favorite int `json:"favorite"`
					Coin     int `json:"coin"`
					Share    int `json:"share"`
					Like     int `json:"like"`
					Mtime    int `json:"mtime"`
					Vt       int `json:"vt"`
					Vv       int `json:"vv"`
				} `json:"stat"`
				EpCount  int `json:"ep_count"`
				FirstAid int `json:"first_aid"`
				Ptime    int `json:"ptime"`
				EpNum    int `json:"ep_num"`
			} `json:"meta"`
			IsAvoided int `json:"is_avoided"`
			Attribute int `json:"attribute"`
		} `json:"vlist"`
	} `json:"list"`
	Page struct {
		Pn    int `json:"pn"`
		Ps    int `json:"ps"`
		Count int `json:"count"`
	} `json:"page"`
	EpisodicButton struct {
		Text string `json:"text"`
		Uri  string `json:"uri"`
	} `json:"episodic_button"`
	IsRisk      bool        `json:"is_risk"`
	GaiaResType int         `json:"gaia_res_type"`
	GaiaData    interface{} `json:"gaia_data"`
}

type Video struct {
	Pic         string `json:"pic"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Length      string `json:"length"`
	Aid         string `json:"aid"`
	BvId        string `json:"bv_id"`
	DownloadUrl string `json:"download_url"`
	CreateTime  int    `json:"create_time"`
}
