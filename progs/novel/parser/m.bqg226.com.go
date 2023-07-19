package parser

import (
	"awesomeProject/crawler/engine"
	"log"
	"os"
	"regexp"
	"strings"
)

const mainUrl = "https://m.bqg226.com"

func Parse(body string, lastData [2]string) engine.ParserResult {
	secName := lastData[0]
	bufContent := lastData[1]
	NewContent := getContent(body)
	nextUrl := getNextUrl(body)
	nextFileName := getNextFileName(nextUrl)
	if nextFileName == "" {
		log.Println("All done！")
		bufContent = bufContent + NewContent
		os.WriteFile("novel/"+string(secName)+".txt", []byte(bufContent), 0660)
		return engine.ParserResult{}
	} else {
		nextSecName := getNextSecName(nextFileName)
		if secName != nextSecName {
			bufContent = bufContent + NewContent
			os.WriteFile("novel/"+string(secName)+".txt", []byte(bufContent), 0660)
			bufContent = ""
			//log.Println("下载完成！")
		} else {
			bufContent = bufContent + NewContent
		}
		newRequest := engine.Request{
			Url: mainUrl + nextUrl,
			Parser: func(body string) engine.ParserResult {
				return Parse(body, [2]string{nextSecName, bufContent})
			},
		}
		return engine.ParserResult{
			Requests: []engine.Request{newRequest},
			Items:    []interface{}{"New Page" + nextFileName},
		}
	}
}

func getContent(s string) string {
	r1, _ := regexp.Compile(`<div id="chaptercontent" class="Readarea ReadAjax_content">([\s\S]+?)<p class="noshow">`)
	f1 := r1.FindStringSubmatch(s)
	content := string(f1[1])
	content = strings.ReplaceAll(content, "<br />", "\n")
	content = strings.ReplaceAll(content, "请收藏：https://m.bqg226.com", "")
	content = strings.TrimSpace(content)
	return content
}

func getNextUrl(s string) string {
	r2, _ := regexp.Compile(`<a href="(.+?)" id="pb_next" class="Readpage_down js_page_down">`)
	f2 := r2.FindStringSubmatch(s)
	next := f2[1]
	return next
}

func getNextFileName(s string) string {
	r3, _ := regexp.Compile(`/([0-9_]+?)\.html`)
	f3 := r3.FindStringSubmatch(s)
	if len(f3) <= 1 {
		return ""
	}
	fileName := f3[1]
	return fileName
}

func getNextSecName(s string) string {
	var secName string
	if i := strings.Index(s, "_"); i >= 0 {
		secName = s[:i]
	} else {
		secName = s
	}
	return secName
}
