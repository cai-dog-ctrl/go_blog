package model

import "blog/pkg/app"

type ArticleSwagger struct {
	List []*Tag
	Pager *app.Pager
}
