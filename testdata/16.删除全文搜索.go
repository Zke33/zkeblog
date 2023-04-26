package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/es_ser"
)

func main() {
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.ESClient = core.EsConnect()
	es_ser.DeleteFullTextByArticleID("eNbxuocBnMkaexhcwNF7")

}
