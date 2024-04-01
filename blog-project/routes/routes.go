package routes

import (
	"blog-project/config"
	"blog-project/controller"
	"blog-project/database"
	"blog-project/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	DatabaseObj database.Database
}

func (db *Router) SetupRouter() *gin.Engine {
	router := gin.Default()

	addRouters(router, db.DatabaseObj)
	return router
}

func addRouters(router *gin.Engine, dbObj database.Database) {
	//var blogMap = make(map[string]model.Blog) // in-memory
	blogC := controller.BlogController{
		//BlogMap:  blogMap,
		Database: &dbObj,
	}
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerProject := router.Group(config.Blog)
	{
		routerProject.POST("", blogC.AddBlog)
		routerProject.GET("", blogC.GetBlogs)
		routerProject.GET(config.ID, blogC.GetBlogByID) // localhost:8080/blog/<id>
		routerProject.PATCH(config.ID, blogC.UpdateBlogByID)
		routerProject.DELETE(config.ID, blogC.DeleteBlogByID)

	}

}

// localhost:8080/blog -> POST
// localhost:8080/blog/{id} -> GetByID
// localhost:8080/blog -> GET
