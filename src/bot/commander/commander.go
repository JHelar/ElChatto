package commander

import (
	"regexp"
	"strings"
	"bot/bottypes"
	"log"
)

var commandRegex *regexp.Regexp
var commandMap map[string]func(*bottypes.Message, []string)

func parseCommand(commandText string)(string, []string){
	commandText = strings.TrimSpace(commandText)
	command := commandRegex.FindString(commandText)
	command = strings.TrimSpace(command)

	paramstr := strings.Replace(commandText, command, "", -1)
	paramstr = strings.TrimSpace(paramstr)

	params := strings.Split(paramstr, ";")

	return command, params
}

func HandleCommand(command string, handler func(*bottypes.Message, []string)){
	if _, ok := commandMap[command]; ok{
		return
	}else{
		commandMap[command] = handler
	}
}

func ExecuteCommand(msg *bottypes.Message, errorHandler func(*bottypes.Message)){
	command, params := parseCommand(msg.Text)
	if val, ok := commandMap[command]; ok {
		val(msg, params)
	}else{
		if errorHandler != nil {
			errorHandler(msg)
		}else{
			log.Printf("Command %v does not exist!", command)
		}
	}
}

func init(){
	commandRegex = regexp.MustCompile(`^/[a-z]+\s`)
	commandMap = make(map[string]func(*bottypes.Message, []string))
}

