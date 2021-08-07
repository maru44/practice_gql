package database

import (
	"context"
	"database/sql"
	"practice_gql/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/lib/pq"
)

func InsertBlog(db *sql.DB) {
	// db, _ := ConnectSql()
	blog := models.Blog{
		Slug:     GenRandSlug(8),
		Title:    "サンプル1",
		Abstract: null.StringFrom("概要1"),
		Content:  "内容\n本文",
		UserID:   GenRandSlug(12),
	}
	blog.Insert(context.Background(), db, boil.Infer())
}

func BlogAll(db *sql.DB) ([]*models.Blog, error) {
	// db, _ := ConnectSql()
	blogs, err := models.Blogs().All(context.Background(), db)
	return blogs, err
}
