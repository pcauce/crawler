package pages

import (
	"fmt"
	"github.com/pcauce/crawler/internal/config"
	"github.com/pcauce/crawler/urls"
	"net/url"
)

func CrawlPage(cfg *config.Config, rawCurrentURL string) {
	cfg.ConcurrencyControl <- struct{}{}
	defer func() {
		<-cfg.ConcurrencyControl
		cfg.Wg.Done()
	}()
	if cfg.IsMaxPagesReached() {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.BaseURL.Hostname() {
		return
	}

	normalizedURL, err := urls.Normalize(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	isFirst := cfg.AddPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := urls.GetFromHTML(htmlBody, cfg.BaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.Wg.Add(1)
		go CrawlPage(cfg, nextURL)
	}
}
