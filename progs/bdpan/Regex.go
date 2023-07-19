package bdpan

import (
	"regexp"
	"strings"
)

var RegIsUrl = regexp.MustCompile(`href="(http.+?)"`)
var RegBdPan = regexp.MustCompile(`pan\.baidu\.com\/s\/([^<]+?码[:：]\s?\w{4}|[\w\s-]+)`)

func GetUrl(str string) []string {
	var urls []string
	matches := RegIsUrl.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if strings.Contains(match[1], "javascript") {
			continue
		}
		urls = append(urls, match[1])
	}
	return urls
}
func GetPanLink(str string) []string {
	//log.Println(RegBdPan.FindAllString(str, -1))
	return RegBdPan.FindAllString(str, -1)
}
