package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func get_url() string {
	botToken := ""
	//https://api.telegram.org/bot.../getMe
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	return botUrl
}

// update
func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset)) //"?offset=" + strconv.Itoa(offset) только ноывй сообщ
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //закрыть тело ответа после обработки, те когда выйдем из ф-ции
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse) //распарсим
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

// answer
func respond(bot_url string, update Update, text string) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = "bot: " + "\nyou id:" + strconv.Itoa(update.Message.Chat.ChatId) + "\n" + update.Message.Text + "\n" + text //тот же текст что и пришло

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(bot_url+"/sendMessage", "application/json", bytes.NewBuffer(buf)) //send json . Специальный ридер
	if err != nil {
		return err
	}

	return nil

}

