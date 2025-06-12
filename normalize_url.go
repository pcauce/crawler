package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	urlData, err := url.Parse(strings.Trim(rawURL, "/"))
	if err != nil {
		return "", err
	}

	return strings.ToLower(fmt.Sprintf("%s%s", urlData.Host, urlData.Path)), nil
}
