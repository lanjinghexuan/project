package pkr

func B() {
	//logger, _ := zap.NewProduction()
	//defer logger.Sync()
	//
	//cfg := es.Config{
	//	Addresses:  []string{"http://localhost:9200"},
	//	MaxRetries: 3,
	//	Timeout:    30 * time.Second,
	//}
	//
	//client, err := es.NewClient(cfg, logger)
	//if err != nil {
	//	log.Fatalf("创建ES客户端失败: %v", err)
	//}
	//
	//indexSvc := es.NewIndexService(client)
	//err = indexSvc.CreateIndex("video_works")
	//if err != nil {
	//	log.Fatalf("创建索引失败: %v", err)
	//}
	//
	//docSvc := es.NewDocumentService(client)
	//video := es.VideoWork{
	//	ID:    1,
	//	Title: "测试视频",
	//}
	//
	//err = docSvc.AddDocument("video_works", video, "1")
	//if err != nil {
	//	log.Fatalf("添加文档失败: %v", err)
	//}
	//
	//searchSvc := es.NewSearchService(client)
	//query := `{
	//	"query": {
	//		"match": {
	//			"title": "测试"
	//		}
	//	}
	//}`
	//
	//results, err := searchSvc.Search("video_works", query)
	//if err != nil {
	//	log.Fatalf("搜索失败: %v", err)
	//}
	//
	//log.Printf("搜索结果: %+v", results)
}
