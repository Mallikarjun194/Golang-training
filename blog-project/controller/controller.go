package controller

import (
	"blog-project/model"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogMap map[string]model.Blog
}

func (b *BlogController) AddBlog(c *gin.Context) {

}

func (b *BlogController) GetBlogs(c *gin.Context) {

}

func (b *BlogController) GetBlogByID(c *gin.Context) {

}

func (b *BlogController) DeleteBlogByID(c *gin.Context) {

}

func (b *BlogController) UpdateBlogByID(c *gin.Context) {

}
