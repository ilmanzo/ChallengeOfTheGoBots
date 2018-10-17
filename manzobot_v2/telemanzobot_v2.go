// telemanzobot project main.go
package main

import (
	"log"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"

	tb "gopkg.in/tucnak/telebot.v2"
)

var bot *tb.Bot

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_API_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	bot = b

	bot.Handle("/reverse", reverseHandler)

	bot.Handle("/kitten", kittenHandler)

	bot.Start()
}

func reverseHandler(m *tb.Message) {
	if len(m.Payload) == 0 {
		bot.Send(m.Sender, "usage: /reverse a string or sentence")
		return
	}
	bot.Send(m.Sender, reverse(m.Payload))
}

func reverse(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		// quietly skip invalid UTF-8
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

func kittenHandler(m *tb.Message) {
	kittens := []string{"kitten_1.jpeg", "kitten_2.jpeg", "kitten_3.jpeg"}
	n := rand.Int() % len(kittens)
	p := &tb.Photo{File: tb.FromDisk(kittens[n])}
	bot.Send(m.Sender, p)
}
