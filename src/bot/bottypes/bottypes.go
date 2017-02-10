package bottypes

type Bot struct {
	Result chan Result
	Lastupdate_id int
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
	Entities []Entity `json:"entities,omitempty"`
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

type DataPacket struct {
	Chat_id int `json:"chat_id"`
	Text string `json:"text"`
	Parse_mode string `json:"parse_mode,omitempty"`
	Reply_to_message_id int `json:"reply_to_message,omitempty"`
}
