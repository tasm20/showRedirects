package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ver string = "2.2"
)

var filename *string = flag.String("f", "", "имя файла")

func init() {
	version := flag.Bool("v", false, "версия")

	flag.String("без ключа", "", "домен[ы] через пробел")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Println(ver)
		os.Exit(0)
	}
}

func main() {
	bots := map[string]string{
		"YANDEX": "Mozilla/5.0 (Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots))",
		"GOOGLE": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}

	domains := domainList(filename)

	for _, domain := range domains {
		fmt.Println("===========", domain, "User bot")
		fmt.Println(showRedirect(domain, ""))

		for name, bot := range bots {
			fmt.Println("===========", domain, name, "BOT")
			fmt.Println(showRedirect(domain, bot))
		}

		fmt.Println()
	}
}
