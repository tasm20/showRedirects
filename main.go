package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

const (
	ver string = "2.4"
)

type Bot struct {
	botName string
	bot     string
	domain  string
	result  chan string
}

func main() {
	version := flag.Bool("v", false, "version")
	filename := flag.String("f", "", "file name")

	flag.String("wothout flag", "", "domains separted by space")

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

	var wg sync.WaitGroup

	for _, domain := range domains {

		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			result := "\n"
			checkedDomain := make(chan string)

			for botname, bot := range bots {
				checkBot := Bot{
					botName: botname,
					bot:     bot,
					domain:  domain,
					result:  checkedDomain,
				}

				go showRedirect(checkBot)
				result += <-checkBot.result + "\n"
			}

			fmt.Println(result)

		}(domain)
	}
	wg.Wait()
}
