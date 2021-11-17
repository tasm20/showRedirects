package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func showRedirect(domain string, bot string) (result string) {
	if bot == "" {
		bot = "curlBot"
	}

	if !strings.HasPrefix(domain, "http") {
		domain = "http://" + domain
	}

	for {
		domain = strings.TrimSuffix(domain, "/")

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 5 * time.Second,
		}

		req, err := http.NewRequest("GET", domain, nil)
		if err != nil {
			fmt.Println(err)
			break
		}

		req.Header.Set("User-Agent", bot)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			break
		}

		result += string(resp.Request.URL.String()) + " "
		result += strconv.Itoa(resp.StatusCode)

		if resp.StatusCode == 301 {
			result += " -> "
			domain = resp.Header.Get("Location")
		} else if resp.StatusCode == 302 {
			result += " -> "
			domain = domain + resp.Header.Get("Location")

			if strings.HasPrefix(resp.Header.Get("Location"), "http") {
				domain = resp.Header.Get("Location")
			}

		} else {
			break
		}
	}

	return result
}
