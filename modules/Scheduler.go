package modules

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/**
程序结束条件：1.已获得30个item 2.已运行30秒
满足其中之一即可
结束后获得的item及对应来源网址将保存在文件中（此部分由work自行实现）
*/

var MaxItemCount = 30
var FetchedItemCount = 0
var DieTimeSecond = 30
var deadLine <-chan time.Time
var doneChannels []chan struct{}

var WorkerCount int
var requests []Request
var ItemHandler func([]interface{})
var visitedUrls = make(map[string][]interface{})

func Init(workerCount int, itemHandler func([]interface{}), seeds ...Request) {
	WorkerCount = workerCount
	ItemHandler = itemHandler
	for _, r := range seeds {
		requests = append(requests, r)
	}
	deadLine = time.After(time.Duration(DieTimeSecond) * time.Second)
}

func Run() {
	var publicIn = make(chan Request)
	var publicOut = make(chan WorkResult)
	for i := 0; i < WorkerCount; i++ {
		createWorker(publicIn, publicOut)
	}
	for {
		for {
			if req, ok := PopRequest(); ok {
				//防止in与out相互等待
				go func() {
					publicIn <- req
				}()
			} else {
				break
			}
		}
		select {
		case result := <-publicOut:
			HandleResult(&result)
			FetchedItemCount += len(result.Items)
			if FetchedItemCount > MaxItemCount {
				log.Println("Item max")
				CallItADay()
				return
			}
		case <-deadLine:
			log.Println("Time out")
			CallItADay()
			return
		}
	}
}

func HandleResult(result *WorkResult) {
	visitedUrls[result.OriginUrl] = result.Items
	//if len(result.Items) > 0 {
	//	fmt.Println(result.OriginUrl)
	//	fmt.Println(result.Items)
	//}
	for _, r := range result.Requests {
		if !isVisited(r.Url) {
			addRequest(r)
		}
	}
	ItemHandler(result.Items)
}

func createWorker(publicIn chan Request, publicOut chan WorkResult) {
	done := make(chan struct{})
	doneChannels = append(doneChannels, done)
	go func() {
		var result WorkResult
		var hasValue = false
		for {
			var ctrlOutChan chan WorkResult
			if hasValue {
				ctrlOutChan = publicOut
			}
			select {
			case req := <-publicIn:
				result = req.Worker(req.Url)
				hasValue = true
			case ctrlOutChan <- result:
				hasValue = false
			case <-done:
				done <- struct{}{}
				return
			}
		}
	}()
}

func isVisited(url string) bool {
	if _, ok := visitedUrls[url]; ok {
		return true
	}
	visitedUrls[url] = make([]interface{}, 0)
	return false
}

func addRequest(req Request) {
	requests = append(requests, req)
}

func PopRequest() (Request, bool) {
	var r Request
	if len(requests) > 0 {
		r = requests[0]
		requests = requests[1:]
	} else {
		return r, false
	}
	return r, true
}

func CallItADay() {
	for i, done := range doneChannels {
		log.Printf("Killing worker %d\n", i)
		done <- struct{}{}
		<-done
	}
	formatTimeStr := time.Now().Format(time.RFC3339)
	formatTimeStr = strings.ReplaceAll(formatTimeStr, ":", ".")
	s := ""
	f, err := os.Create(formatTimeStr + ".txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for k, v := range visitedUrls {
		s = s + k + "\n"
		for _, i := range v {
			s = s + "    >>>>" + Strval(i) + "\n"
			//fmt.Printf("    " + Strval(i) + "\n")
		}
	}
	f.WriteString(s)
	fmt.Printf("Fetched data has been saved at file %s\n", formatTimeStr+".txt")
}
