package utils

import (
	"encoding/xml"
	"fmt"
	"os"
)

func WriteToXML(links []Link, outfile string) {
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	for _, link := range links {
		if err := enc.Encode(link); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}
