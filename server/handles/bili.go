package handles

import (
	"bilibiliRSS/internal/dao"
	"bilibiliRSS/internal/model"
	"bilibiliRSS/rss/bili"
	"bilibiliRSS/server/common"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

func VideoAll(c *gin.Context) {
	pageStr, ok := c.GetQuery("page")
	if !ok {
		common.ErrorResp(c, errors.New("参数错误"))
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		common.ErrorResp(c, errors.New("参数错误"))
		return
	}
	countStr, ok := c.GetQuery("count")
	if !ok {
		common.ErrorResp(c, errors.New("参数错误"))
		return
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		common.ErrorResp(c, errors.New("参数错误"))
		return
	}
	subscriptions, err := dao.SubscriptionDao.GetAll()
	if err != nil {
		common.ErrorResp(c, errors.New("获取数据错误"))
		return
	}
	if len(subscriptions) == 0 {
		common.ErrorResp(c, errors.New("当前没有订阅者"))
		return
	}
	wait := &sync.WaitGroup{}
	videoMetadata := make([]*model.Video, 0, len(subscriptions)*count)
	for _, subscription := range subscriptions { // 循环所有订阅
		wait.Add(1)
		go func(subscription *model.Subscription) { // 开启协程查询每个订阅的视频
			videoList, err := bili.Bili.VideoList(subscription.Mid, page, count)
			if err != nil {
				wait.Done()
				fmt.Println(err)
				return
			}
			for _, video := range videoList { // 循环查询查到的视频
				wait.Add(1)
				go func(video *model.Video) { // 开启协程查询视频直链
					video.Head = subscription.Pic
					cid, err := bili.Bili.AidToCid(video.Aid)
					if err != nil {
						wait.Done()
						fmt.Println(err)
						return
					}
					video.Cid = cid
					url, err := bili.Bili.DownloadUrl(video.Aid, cid, 80)
					if err != nil {
						wait.Done()
						fmt.Println(err)
						return
					}
					video.DownloadUrl = url
					videoMetadata = append(videoMetadata, video)
					wait.Done()
				}(video)
			}
			wait.Done()
		}(subscription)
	}
	wait.Wait()
	common.SuccessResp(c, videoMetadata)
}
func VideoBySubscription(c *gin.Context) {
	mid, err := common.GetQueryInt(c, "mid")
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	page, err := common.GetQueryInt(c, "page")
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	count, err := common.GetQueryInt(c, "count")
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	wait := &sync.WaitGroup{}
	videoMetadata := make([]*model.Video, 0, count)
	subscription, err := dao.SubscriptionDao.GetByMid(int64(mid))
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	videoList, err := bili.Bili.VideoList(subscription.Mid, page, count)
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	for _, video := range videoList { // 循环查询查到的视频
		wait.Add(1)
		go func(video *model.Video) { // 开启协程查询视频直链
			video.Head = subscription.Pic
			cid, err := bili.Bili.AidToCid(video.Aid)
			if err != nil {
				wait.Done()
				fmt.Println(err)
				return
			}
			video.Cid = cid
			url, err := bili.Bili.DownloadUrl(video.Aid, cid, 80)
			if err != nil {
				wait.Done()
				fmt.Println(err)
				return
			}
			video.DownloadUrl = url
			videoMetadata = append(videoMetadata, video)
			wait.Done()
		}(video)
	}
	wait.Wait()
	common.SuccessResp(c, videoMetadata)
}
func Add(c *gin.Context) {
	s := new(model.Subscription)
	err := c.BindJSON(s)
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	if s.Mid == 0 {
		s.Mid, err = bili.Bili.ShareUrlToMid(s.ShareUrl)
		if err != nil {
			common.ErrorResp(c, err)
			return
		}
	}
	subscription, err := bili.Bili.SubscriptionInfo(s.Mid)
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	subscription.ShareUrl = s.ShareUrl
	err = dao.SubscriptionDao.Insert(subscription)
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	common.SuccessResp(c, subscription)
}
func SubscriptAll(c *gin.Context) {
	subscriptions, err := dao.SubscriptionDao.GetAll()
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	common.SuccessResp(c, subscriptions)
}
