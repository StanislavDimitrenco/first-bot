package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func getDailyVerse(url string) []string {

	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		//some error
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	linkAll := doc.Find(".todayItems").Find(".pub-es20").Find(".scalableui")
	title, _ := linkAll.Find("header").Find("h2").Html()
	script, _ := linkAll.Find(".themeScrp").Html()
	text, _ := linkAll.Find(".bodyTxt").Find(".sb").Html()

	srt := []string{title, script, text}

	return srt

}
