package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service/image_ser"
	"gvb_server/utils"
	"path"
	"strings"
)

// ImageUploadDataView 上传单个图片，返回图片的url
// @Tags 图片管理
// @Summary 上传单个图片，返回图片的url
// @Description 上传单个图片，返回图片的url
// @Accept multipart/form-data
// @Param images formData file true  "文件上传"
// @Router /api/image [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUploadDataView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("参数校验失败", c)
		return
	}
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)

	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, image_ser.WhiteImageList) {
		res.FailWithMessage("非法文件", c)
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		res.FailWithMessage(msg, c)
		return
	}
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData("/"+filePath, c)
	return
}
