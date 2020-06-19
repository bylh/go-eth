package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"regexp"
	"time"
)

func main() {
	//getV2EX()
	getSegmentfault()
}
func getSegmentfault() []map[string]interface{} {
	url := "https://segmentfault.com/hottest"
	timeout := 5 * time.Second //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println()
		return []map[string]interface{}{}
	}
	document.Find(".news-list .news__item-info").Each(func(i int, selection *goquery.Selection) {
		coverImg, boolCoverImg := selection.Find(".news-img").Attr("style")
		if boolCoverImg {
			// 匹配()中且以http或https开头的内容
			rgx := regexp.MustCompile(`\((http[s]?.*?)\)`)
			//http|ftp|https
			rs := rgx.FindStringSubmatch(coverImg)
			if len(rs) > 1 {
				coverImg = rs[1]
			} else {
				coverImg = ""
			}
		}
		s := selection.Find(".news__item-title")
		text := s.Text()
		a := selection.Find(".news-img+a")
		url, boolUrl := a.Attr("href")
		fmt.Println(text)
		if len(text) != 0 {
			if !boolUrl {
				url = ""
			}
		}
		allData = append(allData, map[string]interface{}{"title": string(text), "url": "https://segmentfault.com" + url, "cover_image_url": coverImg})
	})
	fmt.Println(allData)
	return allData
}
func getV2EX() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
	timeout := 5 * time.Second //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取")
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取失败")
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".cell.item").Each(func(i int, selection *goquery.Selection) {
		imgUrl, boolImgUrl := selection.Find(".avatar").Attr("src")
		if !boolImgUrl {
			imgUrl = ""
		}

		url, boolUrl := selection.Find(".item_title .topic-link").Attr("href")
		text := selection.Find("a").Text()
		if !boolUrl {
			url = ""
		}
		allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com" + url, "cover_image_url": imgUrl})
	})
	fmt.Println(allData)
	return allData
}
