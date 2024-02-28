package minio

import (
	"fmt"
	"github.com/QinL233/go-plus/pkg/yaml"
	"github.com/minio/minio-go"
	"log"
)

var driver *minio.Core

func Init() {
	// Initialize minio client object.
	config := yaml.Config.Oss.Minio
	if config.Endpoint == "" || config.Port == 0 {
		log.Println("minio config is empty !")
		return
	}
	var err error
	driver, err = minio.NewCore(fmt.Sprintf("%s:%d", config.Endpoint, config.Port), config.AccessKey, config.SecretKey, config.Secure)
	//minioClient, err = minio.NewCore(global.MinioConfig.Endpoint, global.MinioConfig.AccessKey, global.MinioConfig.SecretKey, global.MinioConfig.Secure)
	if err != nil {
		log.Fatalf("connect minio err: %v", err)
	}
	if flag, err := driver.BucketExists(config.Bucket); err != nil {
		log.Fatalf("connect minio err: %v", err)
	} else if !flag {
		log.Fatalf("connect minio err: %s", config.Bucket)
	}
	log.Println("minio connect success!")
}

func Driver() *minio.Client {
	return driver.Client
}
