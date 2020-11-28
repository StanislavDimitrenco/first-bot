package main

import (
	strip "github.com/grokify/html-strip-tags-go"
	"time"
)

func sendDailyText(t time.Time, update GetUpdatesResult, jwUrl string) {
	frt := t.Format("2006/01/02")
	url := getUrl(getMessage("Мы готовим стих, подождите пару секунд :)", update.Message.Chat.ID))
	getContentByUrl(url)

	dailyVerseMessage := ""

	for _, v := range getDailyVerse(jwUrl + frt) {
		dailyVerseMessage = strip.StripTags(v)
		url := getUrl(getMessage(dailyVerseMessage, update.Message.Chat.ID))
		getContentByUrl(url)
	}
}
