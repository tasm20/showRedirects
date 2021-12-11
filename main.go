package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ver string = "2.3"
)

func main() {
	version := flag.Bool("v", false, "версия")
	filename := flag.String("f", "", "имя файла")

	flag.String("без ключа", "", "домен[ы] через пробел")

	flag.Parse()

	bots := map[string]string{
		"User":   "curl",
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

	for _, domain := range domains {
		for name, bot := range bots {
			fmt.Println("===========", domain, name, "BOT")
			checkedDomain := showRedirect(domain, bot)
			fmt.Println(checkedDomain)
		}

		fmt.Println()
	}
}
