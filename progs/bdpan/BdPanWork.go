package bdpan

import (
	"KaedeCrawler/modules"
	"log"
)

var PanLinks []string

/**
  Work() contains three modules:
     1. Download from the url
     2. Parse the string downloaded
     3. handle the data parsed (output,savefile etc.)
  My design is to avoid handling data in engine
  in order to make the program more extendable
*/

func Work(url string) modules.WorkResult {
	//log.Printf("Fetching: %s\n", url)
	body := modules.AntiAntiCrawlerDownload(url)
	urls := GetUrl(body)
	links := GetPanLink(body)
	var result modules.WorkResult
	//var reqs []engine.Request
	//fmt.Println(body)
	for _, u := range urls {
		//log.Printf("Fetched url: %s\n", u)
		if u == "" {
			continue
		}
		result.Requests = append(result.Requests, modules.Request{
			Url:    u,
			Worker: Work,
		})
	}
	for _, link := range links {
		if link == "" {
			continue
		}
		//log.Printf("Got Link: %q\n", link)
		result.Items = append(result.Items, link)
	}

	return result
}

func HandleResult(items []interface{}) {
	for _, i := range items {
		s := i.(string)
		m := ConvertStrSlice2Map(PanLinks)
		if InMap(m, s) {
			continue
		}
		PanLinks = append(PanLinks, s)
		log.Printf("Got Link: %s\n", s)
	}
}

func ConvertStrSlice2Map(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

func InMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}

func AllDone() {
	log.Printf("Program ended. %d links has been founded.\n", len(PanLinks))
}
