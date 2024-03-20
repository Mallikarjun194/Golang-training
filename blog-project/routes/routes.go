package routes

import (
	"blog-project/config"
	"blog-project/controller"
	"blog-project/model"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	addRouters(router)
	return router
}

func addRouters(router *gin.Engine) {
	var blogMap = make(map[string]model.Blog) // in-memory
	blogC := controller.BlogController{
		BlogMap: blogMap,
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
