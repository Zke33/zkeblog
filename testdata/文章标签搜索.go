package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.ESClient = core.EsConnect()

	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))

	query := elastic.NewBoolQuery()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	fmt.Println(string(result.Aggregations["tags"]))
}
