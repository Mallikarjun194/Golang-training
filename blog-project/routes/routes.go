package routes

import (
	"blog-project/config"
	"blog-project/controller"
	"blog-project/database"
	"github.com/gin-gonic/gin"
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
		Database: dbObj,
	}

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
