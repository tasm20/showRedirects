package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ver string = "2.7.4"
)

// Bot overwrite
type Bot struct {
	botName string
	bot     string
	domain  string
}

var count int32

func main() {
	start := time.Now()

	version := flag.Bool("v", false, "version")
	filename := flag.String("f", "", "file name")
	slowCheck := flag.Bool("s", false, "slow check")

	flag.String("without flag", "", "domains separted by space")

	flag.Parse()

	bots := map[string]string{
		"USER":   "curl/7.74.0",
		"YANDEX": "Mozilla/5.0 (Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots))",
		"GOOGLE": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}

	if len(os.Args[1:]) < 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Println(ver)
		os.Exit(0)
	}

	domains := domainList(filename)

	if !*slowCheck {
		var wg sync.WaitGroup

		for _, domain := range domains {

			wg.Add(1)
			go func(domain string) {
				atomic.AddInt32(&count, 1)

				defer wg.Done()
				result := "\n"

				for botname, bot := range bots {
					checkBot := Bot{
						botName: botname,
						bot:     bot,
						domain:  domain,
					}

					result += showRedirect(checkBot)
					result += "\n"
				}

				fmt.Println(result)

			}(domain)
		}
		wg.Wait()
	} else {
		for _, domain := range domains {
			atomic.AddInt32(&count, 1)

			result := "\n"

			for botname, bot := range bots {
				checkBot := Bot{
					botName: botname,
					bot:     bot,
					domain:  domain,
				}

				result += showRedirect(checkBot)
				result += "\n"
			}

			fmt.Println(result)
		}
	}

	duration := time.Since(start)
	fmt.Printf("\nWere checked %d domains for %v\n", count, duration)
}
