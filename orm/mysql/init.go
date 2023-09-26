package mysql

import (
	"fmt"
	"github.com/QinL233/go-plus/yaml"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var driver *gorm.DB

func Init() {
	config := yaml.Config.Orm.Mysql
	if config.Host == "" || config.Port == 0 || config.Database == "" || config.Username == "" || config.Password == "" {
		log.Println("mysql config is empty !")
		return
	}
	logMode := logger.Warn
	if yaml.Config.Web.Mode != gin.ReleaseMode {
		logMode = logger.Info
	}
	var err error
	//createDatabaseIfNotExist=true&useUnicode=true&characterEncoding=utf8&allowPublicKeyRetrieval=true&useSSL=false&serverTimezone=Asia/Shanghai&allowMultiQueries=true&useAffectedRows=true
	driver, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logMode), //日志级别
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //取消表明被加s
			},
			DisableForeignKeyConstraintWhenMigrating: true, //取消外键约束
			SkipDefaultTransaction:                   true, //禁用默认事务可以提升性能
		})

	if err != nil {
		log.Fatalf("connect mysql err: %v", err)
	}

	db, err := driver.DB()

	if err != nil {
		log.Fatalf("connect mysql err: %v", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	log.Println("mysql connect success!")
}

//从连接池获取连接
func Driver() *gorm.DB {
	return driver
}
