package gqlconv

import (
	"practice_gql/graphql/graph/model"
	"practice_gql/models"
)

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
