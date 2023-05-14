package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	IsCollect bool `json:"is_collect"` //用户是否收藏
}

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Router /api/articles/{id} [get]
// @Param id path string true  "id"
// @Produce json
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_ser.NewArticleLook().Set(cr.ID)
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	isCollect := IsUserArticleColl(c, model.ID)
	var articleDetail = ArticleDetailResponse{
		ArticleModel: model,
		IsCollect:    isCollect,
	}

	res.OkWithData(articleDetail, c)
}

func IsUserArticleColl(c *gin.Context, articleID string) (isCollect bool) {
	//判断用户是否正常登录
	token := c.GetHeader("token")
	if token == "" {
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return
	}
	//判断是否在redis中
	keys := global.Redis.Keys("logout_*").Val()
	if utils.InList("logout_"+token, keys) {
		return
	}
	var count int64
	global.DB.Model(models.UserCollectModel{}).Where("user_id = ? and article_id = ?", claims.UserID, articleID).Count(&count)
	if count == 0 {
		return
	}
	return true
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
