package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Link struct {
	href string
	text string
}

func main() {
	input_file := "ex1.html"
	content, _ := os.ReadFile(input_file)
	reader := strings.NewReader(string(content))
	doc, err := html.Parse(reader)
	if err != nil {
		panic("Error parsing html file")
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			new_link := Link{}
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					new_link.href = attr.Val
				}
			}
			fmt.Println(n.Descendants())
		}
	}
}
