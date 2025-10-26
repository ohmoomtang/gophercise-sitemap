/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLink(htmlReader io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	for n := range doc.Descendants() {
		var link Link
		var innerText string
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.Href = a.Val
				}
			}
			for ch := range n.ChildNodes() {
				switch ch.Type {
				case html.TextNode:
					innerText += ch.Data
				case html.ElementNode:
					for descCh := range ch.Descendants() {
						if descCh.Type == html.TextNode {
							innerText += descCh.Data
						}
					}
				default:
				}
			}
			link.Text = strings.TrimSpace(innerText)
		}
		if link.Href != "" && link.Text != "" {
			links = append(links, link)
		}
	}
	return links, nil
}
