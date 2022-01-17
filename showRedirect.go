package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func showRedirect(checkBot Bot) string {
	result := "=========== " + checkBot.domain + " " + checkBot.botName + " BOT\n"

	if !strings.HasPrefix(checkBot.domain, "http") {
		checkBot.domain = "http://" + checkBot.domain
	}

	for {
		checkBot.domain = strings.TrimSuffix(checkBot.domain, "/")

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 15 * time.Second,
		}

		req, err := http.NewRequest("GET", checkBot.domain, nil)
		if err != nil {
			fmt.Println(err)
			break
		}

		req.Header.Set("User-Agent", checkBot.bot)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			break
		}

		result += string(resp.Request.URL.String()) + " "
		result += strconv.Itoa(resp.StatusCode)

		if resp.StatusCode == 301 {
			result += " -> "
			checkBot.domain = resp.Header.Get("Location")
		} else if resp.StatusCode == 302 {
			result += " -> "
			checkBot.domain = checkBot.domain + resp.Header.Get("Location")

			if strings.HasPrefix(resp.Header.Get("Location"), "http") {
				checkBot.domain = resp.Header.Get("Location")
			}

		} else {
			break
		}
	}
	return result
}
