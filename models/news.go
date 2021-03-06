package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type NewsTag struct {
	gorm.Model
	From        string `json:"from"` // 反引号声明元信息
	Name        string `json:"name"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	CoverImgUrl string `json:"cover_img_url"`
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

func AddNewsTag(data map[string]interface{}) error {
	tag := NewsTag{
		From:  data["from"].(string),
		Title: data["title"].(string),
		Name:  data["name"].(string),
	}

	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetNewsTags(maps map[string]interface{}, pageNum int, pageSize int) ([]NewsTag, error) {
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
	From         string `json:"from"`           // 来源
	Tag          string `json:"tag"`            // 标签
	Title        string `json:"title"`          // 标题
	Url          string `json:"url"`            // url
	Desc         string `json:"desc"`           // 描述
	CoverImgUrl  string `json:"cover_img_url"`  // 图片,
	AvatarImgUrl string `json:"avatar_img_url"` // 头像
	CommentCount int    `json:"comment_count"`  // 评论数
	ViewCount    int    `json:"view_count"`     // 查看次数
	LikeCount    int    `json:"like_count"`     // 点赞数
	HotCount     int    `json:"hot_count"`      // 热度
	PostTime     string `json:"post_time"`      // 发帖时间
	UpdateTime   string `json:"update_time"`    // 最后更新时间或回复时间
}

func AddNews(data map[string]interface{}) error {
	if data["cover_img_url"] == nil {
		data["cover_img_url"] = ""
	}
	if data["avatar_img_url"] == nil {
		data["avatar_img_url"] = ""
	}
	if data["comment_count"] == nil {
		data["comment_count"] = 0
	}
	if data["view_count"] == nil {
		data["view_count"] = 0
	}
	if data["like_count"] == nil {
		data["like_count"] = 0
	}
	if data["hot_count"] == nil {
		data["hot_count"] = 0
	}
	if data["post_time"] == nil {
		data["post_time"] = ""
	}
	if data["update_time"] == nil {
		data["update_time"] = ""
	}
	fmt.Println("插入的数据", data)
	news := News{
		From:         data["from"].(string),
		Tag:          data["tag"].(string),
		Title:        data["title"].(string),
		Url:          data["url"].(string),
		CoverImgUrl:  data["cover_img_url"].(string),
		AvatarImgUrl: data["avatar_img_url"].(string),
		CommentCount: data["comment_count"].(int),
		ViewCount:    data["view_count"].(int),
		LikeCount:    data["like_count"].(int),
		HotCount:     data["hot_count"].(int),
		PostTime:     data["post_time"].(string),
		UpdateTime:   data["update_time"].(string),
	}
	fmt.Println("向数据库添加：", news)
	//err := db.Model(&News{}).Create()
	if err := db.Create(&news).Error; err != nil {
		fmt.Println("插入数据库出错：", err)
		return err
	}

	return nil
}

func GetNews(maps map[string]interface{}, pageNum int, pageSize int) ([]News, error) {
	var (
		news []News
		err  error
	)
	fmt.Println("tag", maps["tag"])
	fmt.Println("测试条件", pageNum, pageSize, pageSize > 0 && pageNum > 0, maps["tag"])
	//db.Exec(); // 一般来说用于执行增删改
	// db.Raw(); // 用于查询
	//SELECT * FROM blog.blog_news WHERE id >= ((SELECT MAX(id) FROM blog.blog_news)-(SELECT MIN(id) FROM blog.blog_news)) * RAND() + (SELECT MIN(id) FROM blog.blog_news) LIMIT 10
	// 随机查询10条记录
	if maps["tag"] == "" {
		if pageSize == 0 {
			pageSize = 10
		}
		querySql := fmt.Sprintf("SELECT * FROM blog.blog_news WHERE id >= ((SELECT MAX(id) FROM blog.blog_news)-(SELECT MIN(id) FROM blog.blog_news)) * RAND() + (SELECT MIN(id) FROM blog.blog_news) LIMIT %d", pageSize)
		fmt.Println("随机获取数据")
		db.Raw(querySql).Scan(&news)
		return news, nil
	}

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
