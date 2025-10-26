/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

type Link struct {
	Href string `xml:"url>loc"`
	Text string `xml:"-"`
}
