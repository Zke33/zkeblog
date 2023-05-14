package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
	router.POST("image", app.ImageUploadDataView)
	router.GET("images", app.ImageListView)
	router.GET("images_names", app.ImageNameListView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
}
