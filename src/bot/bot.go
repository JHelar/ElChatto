package bot

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
)

const BOT_TOKEN = "293073695:AAGMj6GL_q2vx1kYQMgUkYsv35TQ8WFYfSo"
const API_URL = "https://api.telegram.org/bot"

type Bot struct {
	Result chan Result
	Results chan []Result
	Quit chan bool
}

type Response struct {
	Ok bool `json:"ok"`
	Results []Result `json:"result"`
}

type Result struct {
	Update_id int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Message_id int `json:"message_id"`
	From User `json:"from"`
	Chat Chat `json:"chat"`
	Date int `json:"date"`
	Text string `json:"text"`
	Entities Entity `json:"entities,omitempty"`
}

type User struct {
	Id int `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
}

type Chat struct {
	Id int `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Type string `json:"type"`
}

type Entity struct {
	Type string `json:"type"`
	Offset int `json:"offset"`
	Length int `json:"length"`
}

func getUpdates() (Response, bool) {
	resp, err := http.Get(fmt.Sprintf("%s%s/getUpdates", API_URL, BOT_TOKEN))
	defer resp.Body.Close()

	if err != nil {
		return Response{}, false
	}

	var response Response
	decoder := json.NewDecoder(resp.Body);

	err = decoder.Decode(&response)
	if err != nil {
		return Response{}, false
	}
	return response, true
}

func NewBot() *Bot {
	return &Bot{ Result: make(chan Result), Results: make(chan []Result), Quit: make(chan bool)}
}

func StartListen(bot *Bot) {
	go func(bot *Bot){
		for {
			select{
			case <- bot.Quit:
				return
			default:
				res, ok := getUpdates()

				if ok {
					fmt.Print("Ok")
					bot.Result <- res.Results[len(res.Results)-1]
					bot.Results <- res.Results
				}
			}
			time.Sleep(1 * time.Second)
		}
	}(bot)
}