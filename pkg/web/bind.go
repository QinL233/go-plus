package web

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime"
	"mime/multipart"
)

/**
bind解析multipart/form-data的文件
回调得到文件名、文件大小、文件内容
【注】原bind或readForm方法会读取整个multipart内容，计算出文件大小并将[]byte存放到内存中（一次io）
此处还是使用multipart.reader封装的流，但是会直接将文件流抛给handler用户执行，以此减少重复读取io
1、单个文件时，由于通过头尾字符规则，可以通过content-length计算出文件大小
2、多个文件时，无法获取到文件大小
*/

func BindMultipartFile(c *gin.Context, handler func(filename string, fileLength int64, file io.Reader)) {
	/*
		----------------------------436879025358800755340764【\r\n】
		Content-Disposition: form-data; name="file"; filename="1.gif"【\r\n】
		Content-Type: image/gif【\r\n】
		【\r\n】
		...文件真实内容
		【\r\n】
		----------------------------436879025358800755340764--【\r\n】
		【\r\n】
	*/
	//1、解析contentType获取分隔符
	_, params, err := mime.ParseMediaType(c.GetHeader("Content-Type"))
	if err != nil {
		panic(err)
	}
	boundary, ok := params["boundary"]
	if !ok {
		panic("not boundary")
	}
	//2、生成multipart.reader解析（里边包装了form表单的解析信息）
	r := multipart.NewReader(c.Request.Body, boundary)
	//第一次part时拿到头部，是文件描述信息
	p, err := r.NextPart()
	if err != nil {
		return
	}
	//重新构建描述信息，通过一系列规则计算出文件大小以便oss执行put转发
	var partHeader bytes.Buffer
	partHeader.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	for k, v := range p.Header {
		vs := ""
		for _, s := range v {
			vs += s
		}
		partHeader.WriteString(fmt.Sprintf("%s: %s\r\n", k, vs))
	}
	partHeader.WriteString("\r\n")
	var partFooter bytes.Buffer
	partFooter.WriteString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	//文件实际大小等于掐头去尾
	fileLength := c.Request.ContentLength - int64(partHeader.Len()) - int64(partFooter.Len())
	//执行回调
	handler(p.FileName(), fileLength, p)
}

func BindMultipartFiles(c *gin.Context, handler func(filename string, file io.Reader)) {
	/*
		----------------------------436879025358800755340764【\r\n】
		Content-Disposition: form-data; name="files"; filename="1.gif"【\r\n】
		Content-Type: image/gif【\r\n】
		【\r\n】
		...文件真实内容1
		【\r\n】
		----------------------------436879025358800755340764--【\r\n】
		Content-Disposition: form-data; name="files"; filename="2.gif"【\r\n】
		Content-Type: image/gif【\r\n】
		【\r\n】
		...文件真实内容2
		【\r\n】
		----------------------------436879025358800755340764--【\r\n】
		【\r\n】
	*/
	//1、解析contentType获取分隔符
	_, params, err := mime.ParseMediaType(c.GetHeader("Content-Type"))
	if err != nil {
		panic(err)
	}
	boundary, ok := params["boundary"]
	if !ok {
		panic("not boundary")
	}
	//2、生成multipart.reader解析（里边包装了form表单的解析信息）
	r := multipart.NewReader(c.Request.Body, boundary)
	//第一次part时拿到头部，是文件描述信息
	for {
		p, err := r.NextPart()
		if err != nil {
			return
		}
		//执行回调
		handler(p.FileName(), p)
	}
}
