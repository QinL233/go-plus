package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"log"
)

func Init(r *gin.Engine, doc *swag.Spec) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Printf("start swagger %s : /swagger/index.html", doc.Title)
}
