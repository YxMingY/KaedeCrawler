package modules

var WorkerCount int
var requests []Request
var ItemHandler func([]interface{})
var visitedUrls = make(map[string]bool)

func Init(workerCount int, itemHandler func([]interface{}), seeds ...Request) {
	WorkerCount = workerCount
	ItemHandler = itemHandler
	for _, r := range seeds {
		requests = append(requests, r)
	}
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
		result := <-publicOut
		addRequest(result.Requests...)
		ItemHandler(result.Items)
	}
}

func createWorker(publicIn chan Request, publicOut chan WorkResult) {
	go func() {
		for {
			req := <-publicIn
			publicOut <- req.Worker(req.Url)
		}
	}()
}

func isVisited(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

// 已去重
func addRequest(req ...Request) {
	for _, r := range req {
		if !isVisited(r.Url) {
			requests = append(requests, r)
		}
	}
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
