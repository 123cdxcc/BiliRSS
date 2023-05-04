package service

import (
	"bilibiliRSS/internal/model"
	"fmt"
	"testing"
)

func TestUserDao_GetUserByUsernameAndPassword(t *testing.T) {
	user, err := UserService.GetUserByUsernameAndPassword("demo", "demo")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("user", user)
}
func TestUserSubscriptionDao_GetAllByUserID(t *testing.T) {
}
func TestUserDao_Insert(t *testing.T) {
	UserService.Insert("demo", "demo")
}
func TestSubscriptionDaoInsert(t *testing.T) {
	s := &model.Subscription{
		Name: "朱志文环球骑行",
		Mid:  479592209,
	}
	err := SubscriptionService.Insert(s)
	if err != nil {
		fmt.Println(err)
	}
}
