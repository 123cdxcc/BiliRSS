package dao

import (
	"bilibiliRSS/internal/model"
	"fmt"
	"testing"
)

func TestSubscriptionDaoInsert(t *testing.T) {
	s := &model.Subscription{
		Name: "朱志文环球骑行",
		Mid:  "479592209",
	}
	err := SubscriptionDao.Insert(s)
	if err != nil {
		fmt.Println(err)
	}
}
