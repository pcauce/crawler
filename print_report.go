package main

import (
	"fmt"
	"github.com/pcauce/crawler/internal/config"
	"sort"
)

type urlCount struct {
	url   string
	count int
}

func printReport(cfg *config.Config) {
	fmt.Println(fmt.Sprintf("=============================\n  REPORT for %s://%s \n=============================", cfg.BaseURL.Scheme, cfg.BaseURL.Hostname()))

	var sortedURLs []urlCount
	for url, count := range cfg.Pages {
		sortedURLs = append(sortedURLs, urlCount{url, count})
	}

	sort.Slice(sortedURLs, func(i, j int) bool {
		return sortedURLs[i].count > sortedURLs[j].count
	})

	for sURL := range sortedURLs {
		fmt.Printf("Found %d internal links to https://%s\n", sortedURLs[sURL].count, sortedURLs[sURL].url)
	}
}
