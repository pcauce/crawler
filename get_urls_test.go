package main

import (
	"net/url"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		HTMLBody    string
		expected    []string
		errIntended bool
	}{
		{
			name:    "absolute and relative URLs",
			baseURL: "https://blog.boot.dev",
			HTMLBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected:    []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			errIntended: false,
		},
		{
			name:    "malformed HTML",
			baseURL: "https://blog.boot.dev",
			HTMLBody: `
				<html>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</html>
				`,
			expected:    []string{},
			errIntended: true,
		},
		{
			name:    "no links",
			baseURL: "https://blog.boot.dev",
			HTMLBody: `
				<html>
					<body>
						<p>Boot.dev Blog</p>
					</body>
				</html>
				`,
			expected:    []string{},
			errIntended: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			tc.baseURL = baseURL.String()
			actual, err := getURLsFromHTML(tc.HTMLBody, baseURL)
			switch tc.errIntended {
			case true:

			case false:
				if err != nil {
					t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
					return
				}

				for i, url := range actual {
					if url != tc.expected[i] {
						t.Errorf("Test %v - %s FAIL: expected URL list: %v, actual: %v", i, tc.name, tc.expected, actual)
					}
				}
			}
		})
	}
}
