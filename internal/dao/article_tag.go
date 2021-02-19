package dao

import "github.com/chen2eric/blog-service/internal/model"

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBy string) error {
	articleTag := model.ArticleTag{
		TagID:     tagID,
		ArticleID: articleID,
		Model: &model.Model{
			CreatedBy: createBy,
		},
	}
	return articleTag.Create(d.engine)
}
func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}
