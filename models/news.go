package models

import "github.com/jinzhu/gorm"

type NewsTag struct {
	gorm.Model
	From          string `json:"from"` // 反引号声明元信息
	Name          string `json:"name"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
}

// 查找是否存在
func ExistNewsTagByName(name string) (bool, error) {
	var newsTag NewsTag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&newsTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if newsTag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddTag Add a NewsTag
func AddNewsTag(name string) error {
	tag := NewsTag{
		Name: name,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetNewsTags(pageNum int, pageSize int, maps interface{}) ([]NewsTag, error) {
	var (
		tags []NewsTag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetNewsTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&NewsTag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistTagByID determines whether a NewsTag exists based on the ID
func ExistNewsTagByID(id int) (bool, error) {
	var tag NewsTag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteTag delete a tag
func DeleteNewsTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&NewsTag{}).Error; err != nil {
		return err
	}

	return nil
}

// EditTag modify a single tag
func EditNewsTag(id int, data interface{}) error {
	if err := db.Model(&NewsTag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllTag clear all tag
func CleanAllNewsTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&NewsTag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
