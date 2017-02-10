package ElChatto

import (
	"bot/bottypes"
	"bot"
	"googlemaps"
	"fmt"
	"bot/commander"
	"log"
)
const ElChattoTag = "<b>ElChatto:</b> "
var elchatto *bottypes.Bot

func handleDir(msg *bottypes.Message, params []string){
	if(len(params) >= 2) {
		dir := googlemaps.GetDirectionRequest(params[0], params[1])
		mapstr := googlemaps.GetStaticMapFromPolyLine(dir.Routes[0].Overview_polyline)
		bot.SendPhoto(msg, mapstr)
	}
}

func handleError(msg *bottypes.Message){
	msg.Text = fmt.Sprintf("%v Unknown command: '%v'", ElChattoTag, msg.Text)
	bot.SendMessage(msg)
}

func handleResult(msg *bottypes.Message, isCommand bool){
	if isCommand {
		commander.ExecuteCommand(msg, handleError)
	}else{
		msg.Text = fmt.Sprintf("%v %v",ElChattoTag, msg.Text)
		bot.SendMessage(msg)
	}
}

func Start(){
	bot.Listen(elchatto)
	bot.Read(elchatto, handleResult)
	log.Print("ElChatto is listening!")
}

func init(){
	elchatto = bot.NewBot()
	commander.HandleCommand("/dir", handleDir)
}
