package minio

import (
	"fmt"
	"github.com/QinL233/go-plus/yaml"
	"github.com/minio/minio-go"
	"io"
	"log"
	"math/rand"
	"time"
)

// Upload 上传[]byte
func Upload(filename string, reader io.Reader) string {
	bucket := yaml.Config.Oss.Minio.Bucket
	object := fmt.Sprintf("%s/%s/%s", time.Now().Format("2006/01/02"), uuid(), filename)
	_, err := Driver().PutObject(bucket, object, reader, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Println(err)
		return ""
	}
	return object
}

func uuid() string {
	// 创建一个长度为 16 字节的切片
	r := make([]byte, 16)

	// 从加密随机数生成器中读取随机字节
	_, err := rand.Read(r)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return ""
	}

	// 设置 UUID 版本号和变体号
	r[6] = (r[6] & 0x0f) | 0x40
	r[8] = (r[8] & 0x3f) | 0x80

	// 将 UUID 转换为字符串并输出
	return fmt.Sprintf("%x-%x-%x-%x-%x", r[0:4], r[4:6], r[6:8], r[8:10], r[10:])
}
