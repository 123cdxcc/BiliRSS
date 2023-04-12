package route

import (
	"bilibiliRSS/server/handles"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	bili := r.Group("bili")

	video := bili.Group("video")
	video.GET("all", handles.VideoAll)

	subscription := bili.Group("subscription")
	subscription.GET("all", handles.SubscriptAll)
	subscription.POST("add", handles.Add)
}
