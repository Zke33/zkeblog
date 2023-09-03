package chat_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ChatListView 群聊聊天记录
// @Tags 聊天管理
// @Summary 群聊聊天记录
// @Description 群聊聊天记录
// @Router /api/chat_groups_records [get]
// @Param data query models.PageInfo    false  "表示多个参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ChatModel]}
func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	cr.Sort = "created_at desc"
	list, count, _ := common.ComList(models.ChatModel{IsGroup: true}, common.Option{
		PageInfo: cr,
	})

	if list == nil { //先判断切片是否为空，为空的切片不需要过滤
		res.OkWithList(make([]struct{}, 0), int64(count), c)
	} else {
		res.OkWithList(filter.Omit("list", list), int64(count), c)
	}
}
