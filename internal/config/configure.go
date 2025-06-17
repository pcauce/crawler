package config

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	Pages              map[string]int
	MaxPages           int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func (cfg *Config) AddPageVisit(normalizedURL string) (isFirst bool) {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()

	if _, visited := cfg.Pages[normalizedURL]; visited {
		cfg.Pages[normalizedURL]++
		return false
	}
	cfg.Pages[normalizedURL] = 1
	return true
}

func (cfg *Config) IsMaxPagesReached() bool {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()

	return len(cfg.Pages) >= cfg.MaxPages
}

func Configure(rawBaseURL string, maxConcurrency, maxPages int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &Config{
		Pages:              make(map[string]int),
		MaxPages:           maxPages,
		BaseURL:            baseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
	}, nil
}
