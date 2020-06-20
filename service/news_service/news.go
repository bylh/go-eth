package news_service

import (
	"bytes"
	"fmt"
	"go-eth/models"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"
)

/**
部分热榜标题需要转码
*/
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

/**
执行每个分类数据
*/
func ExecGetData(spider Spider, ch chan FetchData) FetchData {
	// 函数传入值为spider
	reflectValue := reflect.ValueOf(spider)
	// 方法类型
	dataType := reflectValue.MethodByName("Get" + spider.DataType)
	// 调用方法
	data := dataType.Call(nil)
	originData := data[0].Interface().([]map[string]interface{})
	start := time.Now()
	group.Done()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, spider.DataType)
	fetchData := FetchData{
		Type: spider,
		Data: originData,
	}
	if len(fetchData.Data) > 0 {
		//fmt.Println(reflect.TypeOf(data.Data[0]["url"]), data.Data[0]["url"])
		for i, item := range fetchData.Data {
			url := item["url"]
			title := item["title"]
			cover_img_url := item["cover_img_url"]
			avatar_img_url := item["avatar_img_url"]
			if url == nil {
				url = ""
			}
			if title == nil {
				title = ""
			}
			if cover_img_url == nil {
				cover_img_url = ""
			}
			if avatar_img_url == nil {
				avatar_img_url = ""
			}
			err := AddNews(map[string]interface{}{
				"from":           fetchData.Type.DataType,
				"tag":            fetchData.Type.DataType,
				"url":            url,
				"title":          title,
				"cover_img_url":  cover_img_url,
				"avatar_img_url": avatar_img_url,
			})
			if err == nil {
				fmt.Println(fetchData.Type.DataType, i, "插入成功")
			}
		}
	}
	ch <- fetchData
	return fetchData
}

var group sync.WaitGroup

// 抓取所有消息
func FetchNews() []FetchData {
	// 抓取前清空消息
	_, err := CleanAllNews()
	if err != nil {
		log.Fatal("清除消息表出错")
	}
	allData := []string{
		"V2EX",
		"ZhiHu",
		"WeiBo",
		"TieBa",
		"DouBan",
		"TianYa",
		"HuPu",
		"GitHub",
		"BaiDu",
		"36Kr",
		"QDaily",
		"GuoKr",
		"HuXiu",
		"ZHDaily",
		"Segmentfault",
		"WYNews",
		"WaterAndWood",
		"HacPai",
		"KD",
		"NGA",
		"WeiXin",
		"Mop",
		"Chiphell",
		"JianDan",
		"ChouTi",
		"ITHome",
	}
	fmt.Println("开始抓取" + strconv.Itoa(len(allData)) + "种数据类型")

	// 阻塞主线程，直到所有的goroutine完成（goroutine个数为传入参数）
	group.Add(len(allData))

	ch := make(chan FetchData)
	var spider Spider
	for _, value := range allData {
		fmt.Println("开始抓取" + value)
		spider = Spider{DataType: value}
		go ExecGetData(spider, ch)
	}
	group.Wait()
	i := 0
	dataArr := make([]FetchData, len(allData))
	for range allData {
		data := <-ch
		if !ExistNewsTagByName(data.Type.DataType) {
			err := AddNewsTag(map[string]interface{}{
				"name":  data.Type.DataType,
				"from":  data.Type.DataType,
				"title": data.Type.DataType,
			})
			if err != nil {
				fmt.Println("添加news tag 失败", data.Type.DataType)
			}
		}
		dataArr[i] = data
		i++
	}
	fmt.Print("完成抓取")
	return dataArr
}

func ExistNewsTagByName(name string) bool {
	exist, err := models.ExistNewsTagByName(name)
	if err != nil {
		return false
	}
	return exist
}

func AddNewsTag(tag map[string]interface{}) error {
	return models.AddNewsTag(tag)
}
func GetNewsTags(maps map[string]interface{}, pageNum int, pageSize int) ([]models.NewsTag, error) {
	newsTags, err := models.GetNewsTags(maps, pageNum, pageSize)

	if err != nil {
		return nil, err
	}
	return newsTags, nil
}

/**
添加消息
*/
func AddNews(data map[string]interface{}) error {
	return models.AddNews(data)
}

/**
获取消息
*/
func GetNews(maps map[string]interface{}, pageNum int, pageSize int) ([]models.News, error) {
	news, err := models.GetNews(maps, pageNum, pageSize)

	if err != nil {
		return nil, err
	}
	return news, nil
}

/**
清空所有消息
*/
func CleanAllNews() (bool, error) {
	return models.CleanAllNews()
}

// 大写的函数是默认导出的，此函数用于定时器，爬取消息
func Setup() {
	FetchNews()
	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for range ticker.C {
			fmt.Println("开始抓取消息数据")
			FetchNews()
		}
	}()
}
