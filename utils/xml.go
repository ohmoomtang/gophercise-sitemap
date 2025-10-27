/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"encoding/xml"
	"fmt"
	"os"
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
	urls = RemoveDuplicates(urls)
	output, err := xml.MarshalIndent(urls, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	if outfile != "" {
		file, err := os.Create(outfile)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(xml.Header + XML_SCHEMA_HEADER + string(output))
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	} else {
		fmt.Println(xml.Header + XML_SCHEMA_HEADER + string(output))
	}
}

func RemoveDuplicates(slice []MyXmlElement) []MyXmlElement {
	encountered := make(map[MyXmlElement]bool)
	result := []MyXmlElement{}

	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}
