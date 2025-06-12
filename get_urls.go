package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	docTree, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	var urls []string
	for node := range docTree.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attribute := range node.Attr {
				if attribute.Key == "href" {
					if attribute.Val[0] == '/' {
						urls = append(urls, baseURL.String()+attribute.Val[1:])
						break
					}
					urls = append(urls, attribute.Val)
					break
				}
			}
		}
	}

	return urls, nil
}
