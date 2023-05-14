package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()

	var RootCommentList []models.CommentModel

	// 先把文章下的根评论查出来
	global.DB.Find(&RootCommentList, "article_id = ? and parent_comment_id is null", "kdaw_4cBnMkaexhcatGC")
	fmt.Println(RootCommentList)

	//	// 遍历根评论，递归查根评论下的所有子评论
	//	diggInfo := redis_ser.NewCommentDigg().GetInfo()
	//	for _, model := range RootCommentList {
	//		var subCommentList, newSubCommentList []models.CommentModel
	//		FindSubComment(*model, &subCommentList)
	//		for _, commentModel := range subCommentList {
	//			digg := diggInfo[fmt.Sprintf("%d", commentModel.ID)]
	//			commentModel.DiggCount = commentModel.DiggCount + digg
	//			newSubCommentList = append(newSubCommentList, commentModel)
	//		}
	//		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
	//		model.DiggCount = model.DiggCount + modelDigg
	//		model.SubComments = newSubCommentList
	//	}
	//	return
	//}
}
