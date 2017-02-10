package bot

import (
	"bot/bottypes"
	"net/http"
	"encoding/json"
	"time"
	"log"
	"fmt"
	"bytes"
)

const BOT_TOKEN = "293073695:AAGMj6GL_q2vx1kYQMgUkYsv35TQ8WFYfSo"
const API_URL = "https://api.telegram.org/bot"


func getUpdates() (bottypes.Response, bool) {
	resp, err := http.Get(fmt.Sprintf("%s%s/getUpdates", API_URL, BOT_TOKEN))
	defer resp.Body.Close()

	if err != nil {
		log.Print(err)
		return bottypes.Response{}, false
	}

	var response bottypes.Response
	decoder := json.NewDecoder(resp.Body);

	err = decoder.Decode(&response)

	if err != nil {
		log.Print(err)
		return bottypes.Response{}, false
	}
	return response, true
}

func SendMessage(message *bottypes.Message){
	data := bottypes.DataPacket{Chat_id: message.Chat.Id, Text: message.Text, Reply_to_message_id: message.Message_id, Parse_mode: "HTML"}

	jsonresp, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Print(err)
		return
	}
	if err != nil {
		log.Print(err)
		return
	}else{
		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/sendMessage", API_URL, BOT_TOKEN), bytes.NewBuffer(jsonresp))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		if err != nil {
			log.Print(err)
			return
		}
	}
}

func SendPhoto(message *bottypes.Message, imgstring string){
	data := struct {
		Chat_id int `json:"chat_id"`
		Photo string  `json:"photo"`
	}{Chat_id:message.Chat.Id, Photo:imgstring}
	jsonresp, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Print(err)
		return
	}else{
		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/sendPhoto", API_URL, BOT_TOKEN), bytes.NewBuffer(jsonresp))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		if err != nil {
			log.Print(err)
			return
		}
	}
}

func NewBot() *bottypes.Bot {
	return &bottypes.Bot{ Result: make(chan bottypes.Result), Lastupdate_id:0}
}

func Listen(bot *bottypes.Bot) {
	go func(bot *bottypes.Bot){
		for {
			res, ok := getUpdates()

			if ok && len(res.Results) > 0{
				result := res.Results[len(res.Results) - 1]
				if(result.Update_id != bot.Lastupdate_id){
					bot.Lastupdate_id = result.Update_id
					bot.Result <- res.Results[len(res.Results)-1]
				}

			}
			time.Sleep(1 * time.Second)
		}
	}(bot)
}

func Read(bot *bottypes.Bot, resultHandler func(*bottypes.Message, bool)){
	go func(bot *bottypes.Bot) {
		for {
			select {
			case res := <- bot.Result:
				log.Printf("Bot Read: %+v", res.Message.Text)
				resultHandler(&res.Message, len(res.Message.Entities) > 0)

			default:
				//log.Print("ElChatto no new message.")
			}
		}
	}(bot)
}