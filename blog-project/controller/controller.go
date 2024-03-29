package controller

import (
	"blog-project/database"
	"blog-project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type BlogController struct {
	//BlogMap  map[string]model.Blog
	Database database.Database
}

func (b *BlogController) AddBlog(c *gin.Context) {
	var blog model.Blog

	err := c.ShouldBindJSON(&blog)
	if err != nil {
		logrus.Errorf("bad-request")
		c.JSON(http.StatusBadRequest, model.Error{
			Msg: err.Error(),
		})
		return
	}
	blog.ID = uuid.New().String()
	blog.CreatedAt = time.Now()
	// Adding blog to map
	// TODO add db call
	err = b.Database.Create(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Msg: err.Error()})
		return
	}
	//b.BlogMap[blog.ID] = blog
	logrus.Info("Blog created successfully with the ID", blog.ID)
	c.JSON(http.StatusCreated, blog)

}

func (b *BlogController) GetBlogs(c *gin.Context) {
	var blogs []model.Blog
	//if len(b.BlogMap) < 1 || b.BlogMap == nil {
	//	logrus.Errorf("No blogs are present!!")
	//	c.JSON(http.StatusNotFound, model.Error{Msg: "No blogs are present!!"})
	//	return
	//}
	//for _, blog := range b.BlogMap {
	//	blogs = append(blogs, blog)
	//}
	err := b.Database.FetchAll(&blogs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Msg: err.Error()})
		return
	}
	result := model.Result{Data: blogs, NoOfRecords: len(blogs)}
	if len(blogs) < 1 {
		result.Status = false
	} else {
		result.Status = true
	}
	logrus.Info("Blogs read successfully!!!")
	c.JSON(http.StatusOK, result)

}

func (b *BlogController) GetBlogByID(c *gin.Context) {
	ID := c.Param("id")
	var blog model.Blog
	blog.ID = ID
	//if blog, ok := b.BlogMap[ID]; ok {
	//	logrus.Infof("Blog with ID %s is successfully fetch", ID)
	//	c.JSON(http.StatusOK, blog)
	//	return
	//}
	err := b.Database.FetchByID(&blog)
	if err != nil {
		logrus.Errorf("Blog with ID %s not found!", ID)
		c.JSON(http.StatusNotFound, model.Error{Msg: fmt.Sprintf("Blog with ID %s not found!", ID)})
		return
	}
	logrus.Infof("Blog with ID %s is successfully fetch", ID)
	c.JSON(http.StatusOK, blog)
}

func (b *BlogController) DeleteBlogByID(c *gin.Context) {
	ID := c.Param("id")
	var blog model.Blog
	blog.ID = ID
	//if _, ok := b.BlogMap[ID]; ok {
	//	logrus.Infof("Blog with ID %s is successfully deleted", ID)
	//	delete(b.BlogMap, ID)
	//	c.JSON(http.StatusOK, model.Msg{Message: fmt.Sprintf("Blog with ID %s is successfully deleted", ID)})
	//	return
	//}
	err := b.Database.FetchByID(&blog)
	if err != nil {
		logrus.Errorf("Blog with ID %s not found!", ID)
		c.JSON(http.StatusNotFound, model.Error{Msg: fmt.Sprintf("Blog with ID %s not found!", ID)})
		return
	}
	err = b.Database.Delete(&blog)
	if err != nil {
		logrus.Errorf("Blog with ID %s not found!", ID)
		c.JSON(http.StatusNotFound, model.Error{Msg: fmt.Sprintf("Blog with ID %s not found!", ID)})
		return
	}
	logrus.Infof("Blog with ID %s is successfully deleted", ID)
	c.JSON(http.StatusOK, model.Msg{Message: fmt.Sprintf("Blog with ID %s is successfully deleted", ID)})

}

func (b *BlogController) UpdateBlogByID(c *gin.Context) {
	ID := c.Param("id")
	var blog model.Blog
	blog.ID = ID
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		logrus.Errorf("bad-request")
		c.JSON(http.StatusBadRequest, model.Error{Msg: err.Error()})
		return
	}
	//if _, ok := b.BlogMap[ID]; ok {
	//	logrus.Infof("Blog with ID %s is successfully updated", ID)
	//	blog.CreatedAt = time.Now()
	//	b.BlogMap[ID] = blog
	//	c.JSON(http.StatusOK, model.Msg{Message: fmt.Sprintf("Blog with ID %s is successfully updated!!!", ID)})
	//	return
	//}
	err = b.Database.Update(&blog)
	if err != nil {
		logrus.Errorf("Blog with ID %s not found!", ID)
		c.JSON(http.StatusNotFound, model.Error{Msg: fmt.Sprintf("Blog with ID %s not found!", ID)})
		return
	}
	logrus.Infof("Blog with ID %s is successfully updated", ID)
	c.JSON(http.StatusOK, model.Msg{Message: fmt.Sprintf("Blog with ID %s is successfully updated!!!", ID)})

}
