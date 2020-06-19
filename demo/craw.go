package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"time"
)

type UrlItem struct {
	name string
	url  string
}

var toFetchList = [...]UrlItem{
	{name: "segmentfault", url: "https://segmentfault.com/hottest"},
	{name: "v2ex", url: "https://www.v2ex.com/?tab=hot"},
	{name: "zhihu", url: "https://www.zhihu.com/hot"},
	{name: "zhdaily", url: "http://daily.zhihu.com/"},
	{name: "douban", url: "https://www.douban.com/group/explore"},
	{name: "tianya", url: "http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1"},
	{name: "github", url: "https://github.com/trending"},
	{name: "36kr", url: "https://36kr.com/"},
	{name: "JianDan", url: "http://jandan.net/"},
	{name: "ChouTi", url: "https://www.v2ex.com/?tab=hot"},
}

type NewsItem struct {
	name   string
	title  string
	url    string
	from   string
	tag    string
	desc   string
	imgUrl string
}

func main() {
	fmt.Println("main: 开始抓取")
	totalTime := 0.00

	/* ------------------------- goroutine ------------------------------- */
	start := time.Now()
	ch := make(chan UrlItem)
	for i, v := range toFetchList {
		fmt.Printf("异步抓取%d", i)
		//go func(value string) {
		//	fetch(v.url, ch)
		//}(v.url)
		go fetch(v, ch)
	}
	// 方式一： 遍历信道，但是到达最大数量需要终止，否则没有值，信道一直处于等待状态
	//i := 0
	//for c := range ch {
	//	fmt.Printf("ch: %d %s\n", i, c)
	//	i ++
	//	if i >= len(toFetchList) {
	//		break
	//	}
	//}
	// 方式二： 遍历信道赋值的数组，确保都赋值的情况下长度一致
	for i, _ := range toFetchList {
		data := <-ch
		fmt.Printf("i: %d data: %s\n", i, data)
	}
	seconds := time.Since(start).Seconds()
	totalTime += seconds
	fmt.Printf("抓取完毕，共用时 %.2fs \n", totalTime)
}

func fetch(urlItem UrlItem, ch chan UrlItem) {
	start := time.Now()
	fmt.Printf("fetch start %s\n", urlItem.name)
	timeout := 5 * time.Second //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", urlItem.url, Body)
	if err != nil {
		fmt.Println("设置请求失败")
		return
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("http get error", err)
	}

	//函数结束后关闭相关链接
	defer res.Body.Close()

	if err != nil {
		fmt.Println("read error", err)
		return
	}
	seconds := time.Since(start).Seconds()
	fmt.Printf("fetch finish用时： %.2f\n", seconds)
	switch urlItem.name {
	case "segmentfault":
		handleSegment(res)
		break
	}
	ch <- urlItem
}

func handleSegment(res *http.Response) {
	// 只能有一个reader, 因为读取会就触发Body io.ReadCloser
	//body, err := ioutil.ReadAll(res.Body)
	//err = ioutil.WriteFile("segment.html", body, 0644)
	//if err != nil {
	//	fmt.Printf("写入文件出错");
	//}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("解析body失败")
		return
	}
	document.Find(".news-list .news__item-info").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("a:nth-child(2)").First()
		url, boolUrl := s.Attr("href")
		text := s.Find("h4").Text()
		if len(text) != 0 {
			if boolUrl {
				fmt.Printf("url %s text %s \n", url, text)
			}
		}
	})

	/* 为什么这一段放前面document就查询不到元素了？？？？？
	原因是使用ioread之后body的发生了改变？ 导致document, err := goquery.NewDocumentFromReader(res.Body)失效？
	大概是因为defer res.Body.Close()的关闭时机吧，使用reader后就触发关闭了，所以后面的就是空了
	*/
	//body, err := ioutil.ReadAll(res.Body)
	//err = ioutil.WriteFile("segment.html", body, 0644)
	//if err != nil {
	//	fmt.Printf("写入文件出错");
	//}
}
