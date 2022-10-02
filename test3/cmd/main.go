package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	
)

func main() {
	database()
	bot_url := get_url()
	offset := 0
	for {
		updates, err := getUpdates(bot_url, offset)
		if err != nil {
			log.Println("error:", err.Error())
		}
		for _, update := range updates {
			mes, _ := handCommande(&update.Message)
			respond(bot_url, update, mes)
			offset = update.UpdateId + 1 //только ноывй сообщ
		}
		fmt.Println(updates)
	}
}

