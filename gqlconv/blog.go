package gqlconv

import (
	"practice_gql/graphql/graph/model"
	"practice_gql/interfaces/database"
	"practice_gql/models"

	"github.com/volatiletech/null"
)

func Blog(in *models.Blog) *model.Blog {
	out := &model.Blog{
		ID:        in.ID,
		Slug:      in.Slug,
		Title:     in.Title,
		Abstruct:  &in.Abstract.String,
		Content:   in.Content,
		UserID:    in.UserID,
		IsPublic:  in.IsPublic,
		CreatedAt: in.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: in.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
	return out
}

func Blogs(in []*models.Blog) []*model.Blog {
	out := make([]*model.Blog, len(in))
	for i, v := range in {
		out[i] = &model.Blog{
			ID:        v.ID,
			Slug:      v.Slug,
			Title:     v.Title,
			Abstruct:  &v.Abstract.String,
			Content:   v.Content,
			UserID:    v.UserID,
			IsPublic:  v.IsPublic,
			CreatedAt: v.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		}
	}
	return out
}

func InsertBlog(in *model.NewBlog) *models.Blog {
	out := &models.Blog{
		Slug:     database.GenRandSlug(8),
		Abstract: null.NewString(*in.Abstruct, true),
		Content:  in.Content,
		UserID:   in.UserID,
		IsPublic: in.IsPublic,
	}
	return out
}
