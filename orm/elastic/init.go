package elastic

import (
	"github.com/QinL233/go-plus/yaml"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

var client *elastic.Client

func Init() {
	config := yaml.Config.Orm.Elastic
	if config.Url == "" {
		log.Println("elastic config is empty")
		return
	}
	//1.初始化连接，得到一个client
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(config.Url),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stdout, "", 0)),
	)
	if err != nil {
		log.Fatalf("elastic connect err: %v", err)
	}
	if !client.IsRunning() {
		log.Fatalf("elastic server is not running")
	}
	log.Println("elastic connect success!")
}

func Driver() *elastic.Client {
	return client
}
