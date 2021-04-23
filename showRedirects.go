package main

/*
	проходится по Location заголовкам доменов
*/

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// переменная для тега - получает имя файла со списком доменов
	filename := flag.String("f", "", "имя файла")
	// для справки
	flag.String("без ключа", "", "домен[ы] через пробел")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		// fmt.Printf("Usage\n - showRedirects domainName\nor\n - showRedirects -f filename\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var domains []string
	// юзер агенты ботов
	bots := map[string]string{
		"YANDEX": "Mozilla/5.0 (Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots))",
		"GOOGLE": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}

	// если подкинули файл - читает и собирает из него список доменов
	if *filename != "" {
		file, err := os.Open(*filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		input := bufio.NewScanner(file)

		for input.Scan() {
			if input.Text() == "" {
				break
			}

			domains = append(domains, input.Text())
		}

		file.Close()
		// собирает список доменов из аргуменов командной строки
	} else {
		domains = append(domains, os.Args[1:]...)
	}

	// запускает работу по собранному списку доменов
	for _, domain := range domains {
		fmt.Println("===========", domain, "User bot")
		fmt.Println(showRedirect(domain, ""))

		// запускает проверку с назначенным заголовком юзер агентов ботов
		for name, bot := range bots {
			fmt.Println("===========", domain, name, "BOT")
			fmt.Println(showRedirect(domain, bot))
		}

		// чтоб красивей вывод
		fmt.Println()
	}
}

// проходится по списку, находит заголовки location в цикле
func showRedirect(domain string, bot string) (result string) {
	// если юзер агент не указан - ставим кастомный
	if bot == "" {
		bot = "curlBot"
	}

	// проверяет наличие протокола - без него ошибка запуска скрипта
	if !strings.HasPrefix(domain, "http") {
		domain = "http://" + domain
	}

	// бесконечный цикл запроса Location заголовка - не известно сколько редиректов будет
	for {
		// убирает / с окончания доменного имени - без этого проблемы при редиректах на локаль
		if strings.HasSuffix(domain, "/") {
			domain = strings.TrimSuffix(domain, "/")
		}

		// когда узнаю больше опишу и это действие, а пока создаётся слиент для опроса доменов
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

		// устанавливаем заголовок юзер агента (или кастомный или боты)
		req.Header.Set("User-Agent", bot)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			break
		}

		// продолжаем собирать результат
		result += string(resp.Request.URL.String()) + " "
		result += strconv.Itoa(resp.StatusCode)

		// если 200 - заканчиваем
		// если 302 - обычно это локаль типа "/en" - собираем следующую итерацию
		// если 301 повторяем цикл с новыми данными
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
