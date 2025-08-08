package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

/**
 * 并发抓取多个url
 */
func main() {
	urlList := []string{
		"https://www.bejson.com/color/",
		"https://www.baidu.com",
		"https://yiyan.baidu.com/",
		"https://www.bejson.com/color/",
	}

	urlChan := make(chan string, len(urlList))
	var wg sync.WaitGroup

	//启动3个worker
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, urlChan, &wg)
	}

	/**

	通道的发送和接收操作都是原子性的
	每个值被放入通道后，有且只有一个接收者能获取到它
	内部通过互斥锁（mutex）保证线程安全

	当多个goroutine同时尝试接收时：
	通道会唤醒一个等待的接收者
	其他接收者继续保持阻塞状态
	被唤醒的goroutine获取值后，队列指针(recvx)会移动
	*/
	for _, url := range urlList { //这里放了4个url,并非3个worker都把4个url执行一遍. 而是3个worker抢这4个url, 谁抢到谁执行。3个worker共同完成4个url的抓取。
		urlChan <- url
	}
	close(urlChan)

	wg.Wait()
	fmt.Println("所有worker完成")

}

/*
抓取url内容
*/
func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("获取url失败:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Http 状态码错误:%d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败:%w", err)
	}
	return string(bodyBytes), nil
}

/*
*
提取网页中的标题 title
*/
func extractTitle(html string) string {
	re := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return "未找到标题"
}

func worker(id int, urls <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d start\n", id)
	for url := range urls { //这里会阻塞
		fmt.Printf("worker %d 开始抓取: %s\n", id, url)
		html, err := fetchURL(url)
		if err != nil {
			fmt.Printf("worker %d 错误:%v\n", id, err)
			continue
		}
		title := extractTitle(html)
		fmt.Printf("worker %d 提取到的标题: %s\n", id, title)
	}
	fmt.Printf("worker %d finish\n", id)
}
