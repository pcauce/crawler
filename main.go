package main

import (
	"fmt"
	"github.com/pcauce/crawler/internal/config"
	"github.com/pcauce/crawler/pages"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
	}
	maxPages, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Printf("Error - MaxPages: %v", err)
	}

	cfg, err := config.Configure(rawBaseURL, int(maxConcurrency), int(maxPages))
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.Wg.Add(1)
	go pages.CrawlPage(cfg, rawBaseURL)
	cfg.Wg.Wait()

	printReport(cfg)
}
