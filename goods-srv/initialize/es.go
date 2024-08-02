package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-blog/goods-srv/global"
	"go-blog/goods-srv/model"
	"log"
	"os"
)

func InitEs() {
	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsConfig.Host, global.ServerConfig.EsConfig.Port)
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}
	//新建mapping和index
	exist, err := global.EsClient.IndexExists(model.EsGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exist {
		_, err := global.EsClient.CreateIndex(model.EsGoods{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
