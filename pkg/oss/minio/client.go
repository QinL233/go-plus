package minio

import (
	"fmt"
	"github.com/QinL233/go-plus/pkg/yaml"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Upload 上传文件流
func Upload(filename string, reader io.Reader) string {
	return UploadBucket(yaml.Config.Oss.Minio.Bucket, filename, reader, -1)
}

// UploadBucket 上传文件到指定bucket
func UploadBucket(bucket, filename string, reader io.Reader, size int64) string {
	object := toObject(filename)
	_, err := Driver().PutObject(bucket, object, reader, size, minio.PutObjectOptions{})
	if err != nil {
		log.Println(err)
		return ""
	}
	return object
}

// UploadForward 使用预签名put转发到minio
func UploadForward(filename string, fileLength int64, reader io.Reader) string {
	return UploadBucketForward(yaml.Config.Oss.Minio.Bucket, filename, fileLength, reader)
}

// UploadBucketForward 预上传文件到指定bucket
func UploadBucketForward(bucket, filename string, fileLength int64, reader io.Reader) string {
	object := toObject(filename)
	//获取预上传url
	server, err := Driver().PresignedPutObject(bucket, object, 10*time.Minute)
	if err != nil {
		log.Println(err)
		return ""
	}
	//构建转发
	req, err := http.NewRequest("PUT", server.String(), reader)
	if err != nil {
		log.Println(err)
		return ""
	}
	//req.Header.Set("Content-Type", c.GetHeader("Content-Type"))
	req.ContentLength = fileLength
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		//注意如果contentLenght和reader长度不一致时，会报错
		log.Println(err)
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		s, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return ""
		}
		log.Println(s)
		return ""
	}
	return object
}

// Download 获取object io
func Download(object string) io.ReadCloser {
	return DownloadBucket(yaml.Config.Oss.Minio.Bucket, object)
}

func DownloadBucket(bucket, object string) io.ReadCloser {
	client := Driver()
	file, err := client.GetObject(bucket, object, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil
	}
	return file
}

// DownloadGin 查询object的方式获取流
func DownloadGin(c *gin.Context, object string, attachment bool) {
	DownloadGinBucket(c, yaml.Config.Oss.Minio.Bucket, object, attachment)
}

func DownloadGinBucket(c *gin.Context, bucket, object string, attachment bool) {
	client := Driver()

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

func toObject(filename string) string {
	return fmt.Sprintf("%s/%s/%s", time.Now().Format("2006/01/02"), uuid(), filename)
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
