package model

import "gorm.io/gorm"

type ArticleTag struct {
	*Model
	TagID     uint `json:"tag_id"`
	ArticleID uint `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (at ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ?", at.ArticleID).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (at ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ?", at.TagID).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (at ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?)", articleIDs).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

func (at ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		return err
	}

	return nil
}

func (at ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&at).Where("article_id = ?", at.ArticleID).
		Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (at ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", at.ID).Delete(&at).Error; err != nil {
		return err
	}

	return nil
}

func (at ArticleTag) DeleteOne(db *gorm.DB) error {
	// 超时时间！
	if err := db.Where("article_id = ?", at.ID).Delete(&at).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
