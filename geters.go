package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

//create url
func getUrl(methodName string) string {
	return TelegramBaseUrl + TelegramKey + "/" + methodName
}

//create url
func getUrlWithOffset(methodName string, id int) string {
	return TelegramBaseUrl + TelegramKey + "/" + methodName + "?offset" + strconv.Itoa(id)
}

//get body json
func getContentByUrl(url string) []byte {
	answer, err := http.Get(url)
	checkError(err)
	defer answer.Body.Close()
	content, err := ioutil.ReadAll(answer.Body)
	checkError(err)
	return content
}

//create message

func getMessage(message string, id int) string {
	return MethodSendMessage + "?chat_id=" + strconv.Itoa(id) + "&text=" + message
}
