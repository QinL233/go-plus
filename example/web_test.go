package example

import (
	"fmt"
	"github.com/QinL233/go-plus"
	"github.com/QinL233/go-plus/oss/minio"
	"github.com/QinL233/go-plus/web"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sync"
	"testing"
)

/*
Controller 负责请求出入口处理
1、负责路由入口【api、request_type】
2、【重点】负责参数校验【query、path、body..取参数、校验】
3、负责结果包装【不仅是统一的response，还有具体data的控制权】
4、负责api错误的【汇总以及打印】

Service 负责做业务处理
1、有且仅有一个service接口，一个文件夹是一个实现，一个.go实现一个方法【便于管理】
2、参数从方法参数传递，输出使用指针：func(db,param) (result)【统一】
3、service.go中定义的param和result尽量不要循环引用
4、函数的实现保证参数校验、原子性、健壮性【规范】
5、在抽象类中定义全局以保证多个service传递中实现事务、全局变量的传递性【db传递】
6、统一使用 panic 处理异常，可以减少if err 处理【*争议!】
7、子协程异常处理需自己捕获(牺牲代码健壮性，由调用方自行判断异常，以提高整体代码的可读性)

Dao db-driver处理
1、处理单表的基础增删改查【crud-可封装】
2、处理数据的事务和异常，确保数据的完整性和一致性【db传递事务】

Model 负责定义数据结构
1、定义db表数据结构，包括表的字段、索引、结构体与表结构一致【table】
2、定义request和response结构，包括请求参数、返回结果、返回的错误信息【class】
*/

//1、定义param和result的class
type DemoParam struct {
	Name  string                `form:"name" binding:"required"`
	Ids   []int                 `form:"ids" binding:"required"`
	File  *multipart.FileHeader `form:"file"`
	Token string                `header:"token" binding:"required"`
}

type DemoResult struct {
	//web.Response
	Username string `json:"username"`
	Ids      []int  `json:"ids"`
}

type SysUser struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
}

func server(db *gorm.DB, param DemoParam) (r DemoResult) {
	//1、测试入参情况
	fmt.Println(param)
	res, err := http.Get("https://pss.bdstatic.com/static/superman/img/logo/bd_logo1-66368c33f8.png")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	go e(res.Body)

	//2、测试driverif
	//db == nil {
	//		db = mysql.Driver()
	//	}
	//user := dao.TryOne[SysUser](db, "id = ?", param.Name)
	////user := dao.One[SysUser](db, "",param.Name)
	//err = db.Transaction(func(tx *gorm.DB) error {
	//	dao.Create(tx, &SysUser{Username: param.Name})
	//	panic(errors.New("dao.Create error"))
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(user)
	//r.Username = user.Username

	//3、测试异常
	var dw sync.WaitGroup
	dw.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
				dw.Done()
			}()
			if i%2 == 0 {
				panic(fmt.Errorf("test panic %d", i))
			}
		}()
	}
	dw.Wait()
	testErr(func() {
		if param.Name == "err" {
			fmt.Printf("100")
			//【*注意】此时抛出异常会被主线程recover拦截到
			panic(errors.New("test"))
			fmt.Printf("200")
		}
	})
	//子协程匿名函数-注意此时需要手动捕获异常
	go testErr(func() {
		if param.Name == "err1" {
			fmt.Printf("300")
			//【**注意】此时无法拦截，会退出程序
			panic(errors.New("test1"))
			fmt.Printf("400")
		}
	})
	go testErr(func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		if param.Name == "err2" {
			fmt.Printf("500")
			//【**注意】此时会被子线程的defer拦截
			panic(errors.New("test2"))
			fmt.Printf("600")
		}
	})
	testErr(func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		if param.Name == "err3" {
			fmt.Printf("700")
			//【*注意】此时抛出异常会被【子】线程recover拦截到，因此主线程没有拦截
			panic(errors.New("test3"))
			fmt.Printf("800")
		}
	})
	return
}

func e(src io.ReadCloser) {
	defer src.Close()
	b, _ := ioutil.ReadAll(src)
	fmt.Println(len(b))
}

// @Tags 下载
// @Summary 压缩一张图片
// @Produce  json
// @Param title body DemoParam true "param"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Router /download [post]
func download(c *gin.Context) {
	object := "2023/09/20/31e495f9-ecf2-4961-9fb5-788fc9bc95a5/2022航天领航者计划招商宣讲.mp4"
	//minio.DownloadGin(c, object, false)
	f := minio.Download(object)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	c.JSONP(200, len(b))
}

//回调函数
func testErr(f func()) {
	f()
}

func TestWeb(t *testing.T) {
	//1、使用包装快速定义router
	web.GET[DemoParam, DemoResult]("/query", server)
	web.DELETE[DemoParam, DemoResult]("/path/:name", server)
	web.POST[DemoParam, DemoResult]("/json", server)
	web.PUT[DemoParam, DemoResult]("/form", server)
	//2、使用table自定义router
	web.Controller(func(g *gin.RouterGroup) {
		g.GET("/download", download)
		g.POST("/upload2", func(c *gin.Context) {
			web.AsyncFormFile(c, func(filename string, size int64, file io.Reader) {
				object := minio.Upload(filename, file)
				web.Success(c, object)
			})
		})
		g.POST("/upload3", func(c *gin.Context) {
			web.AsyncFormFile(c, func(filename string, size int64, file io.Reader) {
				object := minio.UploadForward(filename, size, file)
				web.Success(c, object)
			})
		})
	})
	app.Start(func(r *gin.Engine) {
		//此处初始化swagger文档
		//swagger.Init(r, doc)
	}, "config/app.yml")
}
