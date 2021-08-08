package database

import (
	"context"
	"database/sql"
	"practice_gql/models"

	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/lib/pq"
)

func InsertBlog(db *sql.DB, b models.Blog) {
	b.Insert(context.Background(), db, boil.Infer())
}

func BlogAll(db *sql.DB) ([]*models.Blog, error) {
	// db, _ := ConnectSql()
	blogs, err := models.Blogs().All(context.Background(), db)
	return blogs, err
}

func Get(db *sql.DB, id int) (*models.Blog, error) {
	b, err := models.FindBlog(context.Background(), db, id)
	return b, err
}
