package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TelegramKey = "1427431817:AAGs2h-Q1o2fF4msSli1IAfX7_FI-jH2tTE"
const TelegramBaseUrl = "https://api.telegram.org/bot"
const MethodGetMe = "getMe"
//const METHOD_GET_UPDATES = "getUpdates"
//const METHOD_SED_MESSAGE = "sendMessage"




//structure for request get me
type GetMe struct {
	Ok bool `json:"ok"`
	Result GetMeResult `json:"result"`
}

type GetMeResult struct {
	ID int64 `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName string `json:"user_name"`
}




//check error
func checkError(err error) {
	if err != nil {
		fmt.Printf("Error in unmarshal: %s", err.Error())
	}

}

//create url
func getUrl(methodName string) string  {
	return TelegramBaseUrl + TelegramKey + "/" + methodName
}

//get body json
func getContentByUrl(url string) []byte  {
	answer, err := http.Get(url)
	checkError(err)
	defer answer.Body.Close()
	content, err := ioutil.ReadAll(answer.Body)
	checkError(err)
	return content
}

func main() {
	content := getContentByUrl(getUrl(MethodGetMe))
	getMe := GetMe{}

	err := json.Unmarshal(content, &getMe)
	checkError(err)

	fmt.Println(getMe.Result.FirstName)

}




