package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Soul struct {
	gorm.Model
	Title string `json:"title"`
	Hits  string `json:"hits"`
}

// 使用随机获取的时候offset填-1，limit设置随机获取的条数
func GetSouls(maps map[string]interface{}, offset int, limit int) ([]Soul, error) {
	var (
		souls []Soul
		err   error
	)
	fmt.Println("maps", maps)
	//db.Exec(); // 一般来说用于执行增删改
	// db.Raw(); // 用于查询
	//SELECT * FROM blog.blog_soul WHERE id >= ((SELECT MAX(id) FROM blog.blog_soul)-(SELECT MIN(id) FROM blog.blog_soul)) * RAND() + (SELECT MIN(id) FROM blog.blog_soul) LIMIT 10
	// 随机查询10条记录
	if offset == -1 {
		if limit == 0 {
			limit = 4
		}
		querySql := fmt.Sprintf("SELECT * FROM blog.blog_soul WHERE id >= ((SELECT MAX(id) FROM blog.blog_soul)-(SELECT MIN(id) FROM blog.blog_soul)) * RAND() + (SELECT MIN(id) FROM blog.blog_soul) LIMIT %d", limit)
		fmt.Println("随机获取soul")
		db.Raw(querySql).Scan(&souls)
		return souls, nil
	}

	if limit > 0 && offset > 0 {
		// 此处注意 find(&souls)要放到最后，因为传地址修改souls要在最后一步赋值
		err = db.Model(&Soul{}).Where(maps).Offset(offset).Limit(limit).Find(&souls).Error
	} else {
		err = db.Model(&Soul{}).Where(maps).Find(&souls).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return souls, nil
}
