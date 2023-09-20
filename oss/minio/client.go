package minio

import (
	"fmt"
	"github.com/QinL233/go-plus/yaml"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"io"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"strconv"
	"strings"
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

// Download 获取object io
func Download(object string) io.Reader {
	client := Driver()
	bucket := yaml.Config.Oss.Minio.Bucket
	file, err := client.GetObject(bucket, object, minio.GetObjectOptions{})
	defer file.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	return file
}

// DownloadGin 查询object的方式获取流
func DownloadGin(c *gin.Context, object string, attachment bool) {
	client := Driver()
	bucket := yaml.Config.Oss.Minio.Bucket

	info, err := client.StatObject(bucket, object, minio.StatObjectOptions{})
	if err != nil {
		log.Println(err)
	}
	//设置文件的类型
	contentType := "application/octet-stream"
	filename := "file"
	if i := strings.LastIndex(object, "/"); i > 0 {
		filename = object[i+1:]
		if j := strings.LastIndex(filename, "."); j > 0 {
			if ext := filename[j:]; ext != "" {
				contentType = mime.TypeByExtension(ext)
			}
		}
	}
	c.Header("Content-Type", contentType)

	//是否强制让弹出下载窗口浏览器下载
	if attachment {
		c.Header("Content-Disposition", "attachment; filename="+filename)
	}

	//设置文章的长度
	c.Header("Content-Length", strconv.FormatInt(info.Size, 10))

	//告诉浏览器分块返回
	c.Header("Accept-Ranges", "bytes")
	//使用range判断请求是否分段
	options := minio.GetObjectOptions{}
	rangeHeader := c.GetHeader("Range")
	if rangeHeader != "" {
		c.Status(http.StatusPartialContent)
		//获取偏移量
		var start, end int64
		if strings.HasSuffix(rangeHeader, "-") {
			if _, err = fmt.Sscanf(rangeHeader, "bytes=%d-", &start); err != nil {
				log.Println(err)
			}
			end = info.Size - 1
		} else if strings.Contains(rangeHeader, "-") {
			if _, err = fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end); err != nil {
				log.Println(err)
			}
		}
		//告诉浏览器当前块数据的实际偏移量
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, info.Size))
		//分块获取minio数据
		options.SetRange(start, end)
	}
	file, err := client.GetObject(bucket, object, options)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	//忽视通讯IO时的异常
	io.CopyBuffer(c.Writer, file, make([]byte, 1024))
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
