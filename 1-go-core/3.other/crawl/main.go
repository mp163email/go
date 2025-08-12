package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

/*
类似python的页面抓取
*/
func main() {
	resp, err := http.Get("http://www.81.cn/xx_207779/16401057.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("请求失败")
	}

	//使用goquery可以解析抓取以后的html元素
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("提取所有网页标题：")
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%d: %s\n", i+1, s.Text())
	})

	fmt.Println("\n提取所有网页链接")
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Printf("%d: %s  --> %s\n", i+1, s.Text(), href)
		}
	})
}
