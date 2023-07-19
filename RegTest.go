package main

import (
	"KaedeCrawler/progs/bdpan"
	"fmt"
	"os"
)

func main() {
	bs, _ := os.ReadFile("a.txt")
	str := string(bs)
	fmt.Println(len(bdpan.GetPanLink(str)))
	for _, s := range bdpan.GetPanLink(str) {
		fmt.Printf("%q\n\n\n\n\n\n\n", s)
	}
	//fmt.Printf("%q", bdpan.GetPanLink(str))
}
