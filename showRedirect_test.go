package main

import "fmt"

func ExampleshowRedirect() {
	checkBot := Bot{
		botName: "USER",
		bot:     "curl/7.74.0",
		domain:  "apple.com",
	}

	fmt.Println(showRedirect(checkBot))
	// Output:
	// =========== apple.com USER BOT
	// http://apple.com 301 -> https://www.apple.com 200
}
