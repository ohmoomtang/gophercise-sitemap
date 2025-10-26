/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"encoding/xml"
	"fmt"
)

const XML_SCHEMA_HEADER = `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
`

type MyXmlElement struct {
	XMLName xml.Name `xml:"url"`
	URL     string   `xml:"loc"`
}

func WriteToXML(links []Link, outfile string) {
	var urls []MyXmlElement
	for _, link := range links {
		xml := MyXmlElement{URL: link.Href}
		urls = append(urls, xml)
	}
	output, err := xml.MarshalIndent(urls, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(xml.Header + XML_SCHEMA_HEADER + string(output))
}
