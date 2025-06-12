package main

import (
	"fmt"
	"sort"
)

type urlCount struct {
	url   string
	count int
}

func (cfg *config) printReport() {
	fmt.Println(fmt.Sprintf("=============================\n  REPORT for %s://%s \n=============================", cfg.baseURL.Scheme, cfg.baseURL.Hostname()))

	var sortedURLs []urlCount
	for url, count := range cfg.pages {
		sortedURLs = append(sortedURLs, urlCount{url, count})
	}

	sort.Slice(sortedURLs, func(i, j int) bool {
		return sortedURLs[i].count > sortedURLs[j].count
	})

	for sURL := range sortedURLs {
		fmt.Printf("Found %d internal links to https://%s\n", sortedURLs[sURL].count, sortedURLs[sURL].url)
	}
}
