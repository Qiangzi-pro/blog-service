package main

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"testing"
)

func TestModel(t *testing.T) {
	err := global.DBEngine.AutoMigrate(&model.Tag{}, &model.Article{}, &model.ArticleTag{})
	t.Log(err)
}
