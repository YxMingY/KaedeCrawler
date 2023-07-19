package main

import (
	"KaedeCrawler/modules"
	"KaedeCrawler/progs/bdpan"
	"fmt"
	"net/url"
	"time"
)

const BaiduSearch = "https://m.baidu.com/from=844b/pu=sz%401321_1001/s?word="

func main() {
	fmt.Printf("Please input what you want to get:")
	var kw string
	fmt.Scanln(&kw)
	fmt.Printf("Start fetching %q\n", kw)
	time.Sleep(time.Second)

	seed := modules.Request{
		Url:    BaiduSearch + url.QueryEscape(kw+" 百度云"),
		Worker: bdpan.Work,
	}

	modules.Init(10, bdpan.HandleResult, seed)
	modules.Run()
}
