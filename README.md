shows redirects of domain for several bots  
-v - shows version  
-f filename - gets list of domains from file  


example
=====================
```
=========== amazone.com YANDEX BOT
http://amazone.com 301 -> http://www.amazon.fr 301 -> https://www.amazon.fr 200
=========== amazone.com GOOGLE BOT
http://amazone.com 301 -> http://www.amazon.fr 301 -> https://www.amazon.fr 200
=========== amazone.com USER BOT
http://amazone.com 301 -> http://www.amazon.fr 301 -> https://www.amazon.fr 200

=========== github.com USER BOT
http://github.com 301 -> https://github.com 200
=========== github.com YANDEX BOT
http://github.com 301 -> https://github.com 200
=========== github.com GOOGLE BOT
http://github.com 301 -> https://github.com 200