package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := openFile("last_update.txt")
	checkError(err)
	defer file.Close()

	offset64, err := stringToInt(byteToString(file))
	checkError(err)

	offset := int(offset64)

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}

	for {
		content := getContentByUrl(getUrlWithOffset(MethodGetUpdates, offset))
		getUpdate := GetUpdetes{}

		err := json.Unmarshal(content, &getUpdate)
		checkError(err)

		for _, update := range getUpdate.Result {

			if offset == update.UpdateID {

				err = offsetUpdate(offset, file)
				checkError(err)

				if strings.ToLower(update.Message.Text) == "/go" {
					url := getUrl(getMessage(SomeMessage, update.Message.Chat.ID))
					getContentByUrl(url)
					continue
				}
			}

		}
		time.Sleep(1000 * time.Millisecond)
	}

}

const TelegramKey = "1427431817:AAGs2h-Q1o2fF4msSli1IAfX7_FI-jH2tTE"
const TelegramBaseUrl = "https://api.telegram.org/bot"
const MethodGetMe = "getMe"
const MethodGetUpdates = "getUpdates"
const MethodSendMessage = "sendMessage"

var SomeMessage string = "Some type"

//structure for request get me
type GetMe struct {
	Ok     bool        `json:"ok"`
	Result GetMeResult `json:"result"`
}

type GetMeResult struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"user_name"`
}

//structure for send message
type SendMessage struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type Message struct {
	MessageId int                    `json:"message_id"`
	From      GetUpdateResultMessage `json:"from"`
	Chat      GetUpdateResultChat    `json:"chat"`
	Date      int                    `json:"date"`
	Text      string                 `json:"text"`
}

type GetUpdateResultMessage struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type GetUpdateResultChat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type GetUpdetes struct {
	Ok     bool               `json:"ok"`
	Result []GetUpdatesResult `json:"result"`
}

type GetUpdatesResult struct {
	UpdateID int               `json:"update_id"`
	Message  GetUpdatesMessage `json:"message,omitempty"`
}

type GetUpdatesMessage struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

//check error
func checkError(err error) {
	if err != nil {
		fmt.Printf("Error in unmarshal: %s", err.Error())

	}

}

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

//file opening

func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile("last_update.txt", os.O_RDWR, os.ModeExclusive)

	return file, err
}

func byteToString(file *os.File) string {
	data := make([]byte, 64)
	var str string
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str = string(data[:n])
	}

	return str
}

func intToByte(i int64) []byte {
	some := strconv.FormatInt(i, 10)
	return []byte(some)
}

func stringToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 0, 64)
}

func overWrite(file *os.File, sls []byte) error {
	_, err := file.WriteAt(sls, 0)
	return err
}

func offsetUpdate(o int, file *os.File) error {
	o++
	o64 := int64(o)
	return overWrite(file, intToByte(o64))
}
