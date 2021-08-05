package database

import (
	"context"
	"fmt"
	"practice_gql/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/lib/pq"
)

func InsertBlog() {
	db := ConnectSql()
	blog := models.Blog{
		Slug:     GenRandSlug(8),
		Title:    "サンプル1",
		Abstract: null.StringFrom("概要1"),
		Content:  "内容\n本文",
		UserID:   GenRandSlug(12),
	}
	blog.Insert(context.Background(), db, boil.Infer())
}

func BlogAll() {
	db := ConnectSql()
	blogs, _ := models.Blogs().All(context.Background(), db)
	for _, b := range blogs {
		fmt.Println(b)
	}
}
