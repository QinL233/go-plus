package web

import (
	"github.com/QinL233/go-plus/orm/mysql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

/**
Service - 工具
1、用于快速从context实现做提取param并校验
2、封装service.func(param)(result,err)方法的返回
*/

func Service[P any, R any](c *gin.Context, f func(P) (R, error)) {
	param, err := contextExtractParam[P](c)
	if err != nil {
		log.Println(err)
		Fail(c, 500, err)
		return
	}
	//回调方法
	r, err := f(*param)
	if err != nil {
		Fail(c, 500, err)
		return
	}
	Success(c, r)
}

/**
从controller层获取driver，以统一会话期间的driver以方便使用事务
*/

func DBService[P any, R any](c *gin.Context, f func(db *gorm.DB, param P) (R, error)) {
	param, err := contextExtractParam[P](c)
	if err != nil {
		log.Println(err)
		Fail(c, 500, err)
		return
	}
	//回调方法
	r, err := f(mysql.Driver(), *param)
	if err != nil {
		Fail(c, 500, err)
		return
	}
	Success(c, r)
}

//从context中提取参数到P中
func contextExtractParam[P any](c *gin.Context) (*P, error) {
	var param P
	//1、尝试从header中取得参数【不保证校验】 - 获取`header:"paramName"`
	c.ShouldBindHeader(&param)
	//2、根据参数性质bind参数
	if c.Params != nil && len(c.Params) > 0 {
		//uri请求 - 必须注意service使用`uri:"paramName"`接收
		if err := c.ShouldBindUri(&param); err != nil {
			return nil, errors.New("param is fail")
		}
	} else {
		//query/json/form请求
		if err := c.ShouldBind(&param); err != nil {
			return nil, errors.New("param is fail")
		}
	}
	return &param, nil
}
