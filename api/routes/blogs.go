package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/infrastructure"
)

type BlogsRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	blogcontroller controllers.BlogsController
}

func NewBlogsRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	blogcontroller controllers.BlogsController,
) BlogsRoutes {
	return BlogsRoutes{
		logger:         logger,
		router:         router,
		blogcontroller: blogcontroller,
	}
}

func (i BlogsRoutes) Setup() {
	i.logger.Zap.Info("Setting up blogs routes")
	blogs := i.router.Gin.Group("/blog")
	{
		blogs.POST("/create", i.blogcontroller.CreateBlog)
		blogs.POST("/get-all/:cursor", i.blogcontroller.GetAllBlogs)
		blogs.POST("/get-one/:id", i.blogcontroller.GetOneBlog)
		blogs.POST("/update", i.blogcontroller.UpdateBlogs)
		blogs.POST("/delete/:id", i.blogcontroller.DeleteBlogs)

	}
}
