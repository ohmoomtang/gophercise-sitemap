/*
Copyright © 2025 Napawan Srisuksawad <napawan.srisuksawad@gmail.com>
*/
package utils

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

func CheckURL(rawURL string) (bool, error) {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		err := errors.New("URL has invalid format")
		return false, err
	}
	return true, nil
}

func URLtoReader(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
