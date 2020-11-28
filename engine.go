package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

func engine(pref string, offset int, file *os.File) {
	for {
		content := getContentByUrl(getUrlWithOffset(pref, offset))
		getUpdate := GetUpdates{}

		err := json.Unmarshal(content, &getUpdate)
		checkError(err)

		for _, update := range getUpdate.Result {

			if offset == update.UpdateID {

				if strings.ToLower(update.Message.Text) == "/start" {
					url := getUrl(getMessage("Напишите 'Стих на сегодня, и получите стих на сегодня'", update.Message.Chat.ID))
					getContentByUrl(url)
				}

				if strings.Contains(strings.ToLower(update.Message.Text), "сегодня") {
					dt := time.Now()

					sendDailyText(dt, update, JWLink)

				}

				if strings.Contains(strings.ToLower(update.Message.Text), "завтра") {
					dt := time.Now().Add(24 * time.Hour)

					sendDailyText(dt, update, JWLink)
				}

				offset, err = offsetUpdate(offset, file)
				checkError(err)
			}

		}
	}
}
