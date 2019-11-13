package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

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
	err := db.Select("id").Where("name = ?", name).First(&newsTag).Error
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
func GetNewsTags(maps interface{}, pageNum int, pageSize int) ([]NewsTag, error) {
	var (
		tags []NewsTag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
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

/* ---------------------------------------------- news --------------------------------------------------- */

type News struct {
	gorm.Model
	From          string `json:"from"`            // 来源
	Tag           string `json:"tag"`             // 标签
	Title         string `json:"title"`           // 标题
	Url           string `json:"url"`             // url
	Desc          string `json:"desc"`            // 描述
	CoverImageUrl string `json:"cover_image_url"` // 图片
}

func AddNews(data map[string]interface{}) error {
	//from := data["from"]
	//if from == nil {
	//	from = ""
	//}
	news := News{
		From:  data["from"].(string),
		Tag:   data["tag"].(string),
		Title: data["title"].(string),
		Url:   data["url"].(string),
	}

	//err := db.Model(&News{}).Create()
	if err := db.Create(&news).Error; err != nil {
		return err
	}

	return nil
}

func GetNews(maps interface{}, pageNum int, pageSize int) ([]News, error) {
	var (
		news []News
		err  error
	)

	fmt.Println("测试条件", pageNum, pageSize, pageSize > 0 && pageNum > 0)
	if pageSize > 0 && pageNum > 0 {
		// 此处注意 find(&news)要放到最后，因为传地址修改news要在最后一步赋值
		err = db.Model(&News{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&news).Error
	} else {
		err = db.Model(&News{}).Where(maps).Find(&news).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return news, nil
}

func CleanAllNews() (bool, error) {
	if err := db.Unscoped().Delete(&News{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
