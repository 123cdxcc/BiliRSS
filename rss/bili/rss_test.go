package bili

import (
	"fmt"
	"testing"
)

func TestBili_SubscriptionInfo(t *testing.T) {
	subscription, err := Bili.SubscriptionInfo(479592209)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%#v", subscription)
}
