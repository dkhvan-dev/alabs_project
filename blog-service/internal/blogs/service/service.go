package service

import blogs "blogs-api/internal/blogs/model"

type BlogService interface {
	Create(input blogs.BlogCreate)
}
