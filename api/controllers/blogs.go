package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogsController struct {
	blogService services.BlogsService
	logger      infrastructure.Logger
}

func NewBlogsController(
	blogService services.BlogsService,
	logger infrastructure.Logger,
) BlogsController {
	return BlogsController{
		blogService: blogService,
		logger:      logger,
	}
}

func (cc BlogsController) CreateBlog(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBindJSON(&blog); err != nil {
		cc.logger.Zap.Error("Error [CreateBlog] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind blog data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.blogService.CreateBlog(blog); err != nil {
		cc.logger.Zap.Error("Error [CreateBlog] [db CreateBlog]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create blog")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Blog Created Sucessfully")

}

func (cc BlogsController) GetAllBlogs(c *gin.Context) {
	cursor := c.Param("cursor")
	blog, err := cc.blogService.GetAllBlogs(cursor)

	if err != nil {
		cc.logger.Zap.Error("Error [GetAllBlog] [db GetAllBlog]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Get All Blogs")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, blog)
}

func (cc BlogsController) GetOneBlog(c *gin.Context) {
	blogId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	blog, err := cc.blogService.GetOneBlog(blogId)

	if err != nil {
		cc.logger.Zap.Error("Error finding Blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Blog")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, blog)
}

func (cc BlogsController) UpdateBlogs(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blog := models.Blog{}

	if err := c.ShouldBind(&blog); err != nil {
		cc.logger.Zap.Error("Error [UpdateBlog] (ShouldBindJson) : ", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "failed to update blog")
		return
	}

	blog.ID = id

	if err := cc.blogService.UpdateBlogs(blog); err != nil {
		cc.logger.Zap.Error("Error updating blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to update Blog")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Updated")
}

func (cc BlogsController) DeleteBlogs(c *gin.Context) {
	blogId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := cc.blogService.DeleteBlogs(blogId)

	if err != nil {
		cc.logger.Zap.Error("Error Deleting Blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Blog")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Deleted")

}
