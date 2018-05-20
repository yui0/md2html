// go get github.com/russross/blackfriday
// go build md2html.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"

	"github.com/russross/blackfriday"
)

func main() {
	flag.Parse()

	md, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

/*	html := blackfriday.MarkdownBasic(md)
	fmt.Println(string(html))*/

	htmlFlags := blackfriday.HTML_TOC
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	html := blackfriday.MarkdownOptions(md, renderer, blackfriday.Options{Extensions: 0})
	fmt.Println(string(html))
}
