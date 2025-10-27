/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"io"
	"net/url"
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

func CleansingLinks(links []Link, inputUrl string) ([]Link, error) {
	var newLinks []Link
	//Grab FQDN with protocol scheme from input URL first
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return nil, err
	}

	fqdn := parsedURL.Scheme + "://" + parsedURL.Host

	//Add FQDN to the Link if not existing
	for _, link := range links {
		validatedLink, err := url.Parse(link.Href)
		if err != nil {
			return nil, err
		}
		if validatedLink.Scheme == "" && validatedLink.Host == "" {
			link.Href = fqdn + link.Href
		}
		//Filter out not similar domain link and mailto link
		if strings.Contains(link.Href, parsedURL.Host) {
			if validatedLink.Scheme != "mailto" {
				newLinks = append(newLinks, Link{link.Href, link.Text})
			}
		}
	}

	return newLinks, nil
}
