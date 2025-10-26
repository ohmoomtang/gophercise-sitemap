/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"encoding/xml"
	"fmt"
	"os"
)

type MyXmlElement struct {
	XMLName xml.Name `xml:"url"`
	URL     string   `xml:"loc"`
}

func WriteToXML(links []Link, outfile string) {
	enc := xml.NewEncoder(os.Stdout)
	defer enc.Close()
	enc.Indent("  ", "    ")
	for _, link := range links {
		xml := MyXmlElement{URL: link.Href}
		err := enc.Encode(xml)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}
