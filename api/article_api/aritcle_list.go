package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommList(cr.Key, cr.Page, cr.Limit)
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	// json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}
	res.OkWithList(data, int64(count), c)
}
