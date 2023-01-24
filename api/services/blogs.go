package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type BlogsService struct {
	repository repository.BlogsReposiory
}

func NewBlogsService(
	repository repository.BlogsReposiory,
) BlogsService {
	return BlogsService{
		repository: repository,
	}
}

func (c BlogsService) CreateBlog(Blogs models.Blog) error {
	return c.repository.CreateBlog(Blogs)
}

func (c BlogsService) GetAllBlogs(cursor string) ([]models.Blog, error) {
	return c.repository.GetAllBlogs(cursor)
}

func (c BlogsService) GetOneBlog(blogId int64) (Blog models.Blog, err error) {
	return c.repository.GetOneBlog(blogId)
}

func (c BlogsService) UpdateBlogs(Blog models.Blog) error {
	return c.repository.UpdateBlogs(Blog)
}

func (c BlogsService) DeleteBlogs(blogId int64) error {
	return c.repository.DeleteBlogs(blogId)
}
