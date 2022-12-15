package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/shakhboznorbekov/mytasks/golang_crud/actor/api/docs"
	"github.com/shakhboznorbekov/mytasks/golang_crud/actor/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.GET("/user", handlerV1.GetList)
	r.POST("/user", handlerV1.Create)
	r.PUT("/user", handlerV1.Update)
	r.GET("/user/:id", handlerV1.GetById)
	r.DELETE("/user/:id", handlerV1.Delete)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
