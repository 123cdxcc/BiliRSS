package bili

import (
	"fmt"
	"testing"
)

func TestVideoList(t *testing.T) {
	videos, err := VideoList("479592209", 1, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(videos)
	for i, video := range videos {
		fmt.Printf("第%d个视频：标题%s，发布时间：%d，作者：%s，时长：%s，Aid：%s\n", i, video.Title, video.CreateTime, video.Author, video.Length, video.Aid)
	}
}
