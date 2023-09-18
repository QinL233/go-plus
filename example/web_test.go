package example

import (
	"fmt"
	"github.com/QinL233/go-plus"
	"github.com/QinL233/go-plus/orm/mysql"
	"github.com/QinL233/go-plus/web"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mime/multipart"
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
2、参数从方法参数传递，输出使用指针：func(...*param) (*result,err)【统一】
3、service.go中定义的param和result尽量不要循环引用
4、函数的实现保证参数校验、原子性、健壮性【规范】
5、在抽象类中定义全局以保证多个service传递中实现事务、全局变量的传递性【db传递】

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
	File  *multipart.FileHeader `form:"file"`
	Token string                `header:"token" binding:"required"`
}

type DemoResult struct {
	web.Response
	Password string `json:"password"`
	Ids      []int  `json:"ids"`
}

//2、定义service接口并实现
func get(param DemoParam) (DemoResult, error) {
	fmt.Printf("%v", param)
	r := DemoResult{
		Password: "123456",
	}
	mysql.Driver().Raw("select id from template limit 10").Find(&r.Ids)
	return r, nil
}

func get2(db *gorm.DB, param DemoParam) (result []DemoResult, err error) {
	fmt.Printf("%v", param)
	r := DemoResult{}
	r.Password = "123456"
	if err = db.Raw("select id from template limit 10").Find(&r.Ids).Error; err != nil {
		return
	}
	testErr(func() {
		//err = errors.New("test")
		result = make([]DemoResult, 0)
	})
	return
}

func testErr(f func()) {
	f()
}

//3、定义controller
func DemoController(c *gin.Context) {
	//web.Service[DemoParam, DemoResult](c, get)
	web.DBService[DemoParam, []DemoResult](c, get2)
}

func TestWeb(t *testing.T) {
	web.Controller(func(g *gin.RouterGroup) {
		g.GET("/query", DemoController)
		g.GET("/path/:name", DemoController)
		g.POST("/json", DemoController)
		g.POST("/form", DemoController)
	})
	app.Start("config/app.yml")
}
