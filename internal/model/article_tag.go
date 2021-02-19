package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ArticleID, 0).Delete(&a).Error
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	return db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Limit(1).Delete(&a).Error
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Limit(1).Update(values).Error
}

func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}
