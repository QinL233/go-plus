package web

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"mime"
	"strings"
	"time"
)

/**
异步解析multipart/form-data的文件
*/

func AsyncFormFile(c *gin.Context, handler func(filename string, size int64, file io.Reader)) {
	//读取文件头信息
	filename, size, file := parseMultipartFileHeader(c)

	//异步读取
	pr, pw := io.Pipe()

	finish := make(chan bool, 1)

	//直接开启一个线程运行方法（此时reader不一定有，在使用的时候才会阻塞线程）
	go func(filename string, size int64, reader *io.PipeReader) {
		defer func() {
			if err := recover(); err != nil {
				Fail(c, 500, err.(error))
			}
			finish <- true
		}()
		//注意如何handler没有使用reader会直接结束掉
		handler(filename, size, reader)
		//handler执行完毕即判断关闭reader
		reader.Close()
	}(filename, size, pr)

	//异步将file传输到reader
	go func() {
		io.Copy(pw, file)
		pw.Close()
	}()

	select {
	case <-time.After(time.Minute * 5):
		panic(errors.New("超过5分钟"))
	case <-finish:
	}
}

// 仅当有且仅有一个file时有效
func parseMultipartFileHeader(c *gin.Context) (string, int64, io.Reader) {
	//报文内容
	/*
		----------------------------626289440185114288335838
		Content-Disposition: form-data; name="file"; filename="1.txt"
		Content-Type: text/plain

		...文件真实内容

		----------------------------626289440185114288335838--

	*/
	r := bufio.NewReader(c.Request.Body)
	i := 0
	fileHeaderSize := 0
	LastBoundarySize := 0
	//读取前4行内容
	var filename string
	for i < 4 {
		i++
		content, err := r.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		//每行都需要+分行符号(2byte)
		fileHeaderSize += len(content)
		//第一行为文件信息分隔符(末尾也存在一行分隔符信息，因此需要减去该行)
		if i == 1 {
			//第一行为分隔符
			//结尾为(分行符号+(空行)分行符号+分隔符+(空行)分行符号)
			LastBoundarySize += 4 + len(content)
			continue
		} else if i == 2 {
			//第二行有文件名信息
			contentDisposition := string(content)
			if strings.HasPrefix(contentDisposition, "Content-Disposition: ") {
				contentDisposition = contentDisposition[21:]
			}
			_, dispositionParams, err := mime.ParseMediaType(contentDisposition)
			if err != nil {
				panic(err)
			}
			filename = dispositionParams["filename"]
		}
	}
	if filename == "" {
		panic("无法解析到文件头部信息")
	}
	fileLength := c.Request.ContentLength - int64(fileHeaderSize+LastBoundarySize)
	return filename, fileLength, io.LimitReader(c.Request.Body, fileLength)
}
