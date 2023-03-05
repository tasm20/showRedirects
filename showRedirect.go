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

	redirectsCount := 0

	for {
		if redirectsCount > 6 {
			result += "THERE IS TOO MANY REDIRECTS"
			break
		}

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
		respCode := strconv.Itoa(resp.StatusCode)
		
		if resp.StatusCode > 400 {
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, respCode) // red color
			result += colored
			resp.Body.Close()
			break
		} else if resp.StatusCode == 200 {
			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, respCode) // green color
			result += colored
			resp.Body.Close()
			break
		}

		result += strconv.Itoa(resp.StatusCode)

		if resp.StatusCode == 301 {
			redirectsCount++
			result += " -> "
			checkBot.domain = resp.Header.Get("Location")
			resp.Body.Close()
		} else if resp.StatusCode == 302 {
			result += " -> "
			checkBot.domain = checkBot.domain + resp.Header.Get("Location")

			if strings.HasPrefix(resp.Header.Get("Location"), "http") {
				checkBot.domain = resp.Header.Get("Location")
			}
			resp.Body.Close()
		}
	}

	return result
}
