package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/api/docs"
	"github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.POST("/user", handlerV1.Create)
	r.GET("/user/:id", handlerV1.GetById)
	r.GET("/user", handlerV1.GetList)
	r.PUT("/user", handlerV1.Update)
	r.DELETE("/user/:id", handlerV1.Delete)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
