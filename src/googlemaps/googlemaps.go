package googlemaps

import (
	"net/http"
	"encoding/json"
	"log"
)

const STATIC_MAP_API_KEY = "AIzaSyD55li1OuTm-bRAzfO4Mo3AsdNKHywfp1s"
const DIRECTION_API_KEY = "AIzaSyBp9yzE23JEZdVBUjgReVfUellKeZg54S4"
const STATIC_MAP_URL = "https://maps.googleapis.com/maps/api/staticmap?"
const DIRECTION_URL = "https://maps.googleapis.com/maps/api/directions/json?"

type Geocoded_waypoint struct {
	status string `json:"geocoder_status"`
	id string `json:"place_id"`
	types []string `json:"types"`
}

type Direction struct {
	Geocoded_waypoints []Geocoded_waypoint `json:"geocoded_waypoints"`
	Routes []Route `json:"routes"`
	Status string `json:"status"`
}

type Route struct {
	Summary string `json:"summary"`
	Legs []Leg `json:"legs"`
	Copyrights string `json:"copyrights"`
	Overview_polyline Polyline `json:"overview_polyline"`
	Warnings []string `json:"warnings,omitempty"`
	Waypoint_order []int `json:"waypoint_order,omitempty"`
	Bounds Bound `json:"bounds"`
}

type Bound struct {
	Northeast Point `json:"northeast"`
	Southwest Point	`json:"southwest"`
}

type Point struct {
	lat float32 `json:"lat"`
	lng float32 `json:"lng"`
}

type Leg struct {

	Steps []Step `json:"steps"`
	Distance TextValue `json:"distance"`
	Duration TextValue `json:"duration"`
	End_address string `json:"end_address"`
	End_location Point `json:"end_location"`
	Start_address string `json:"start_address"`
	Start_location Point `json:"start_location"`
	Traffic_speed_entry []string `json:"traffic_speed_entry,omitempty"`
	Via_waypoint []string `json:"via_waypoint,omitempty"`
}

type TextValue struct {
	Text string `json:"text"`
	Value int `json:"value"`
}

type Step struct {
	Travel_mode string `json:"travel_mode"`
	Start_location Point `json:"start_location"`
	End_location Point `json:"end_location"`
	Polyline Polyline `json:"polyline"`
	Duration TextValue `json:"duration"`
	Html_instructions string `json:"html_instructions"`
	Distance TextValue `json:"distance"`
	Maneuver string `json:"maneuver,omitempty"`
}

type Polyline struct {
	Points string `json:"points"`
}

func DemoDirectionRequest(fromAddress, toAddress string) *Direction {
	resp, err := http.Get("https://maps.googleapis.com/maps/api/directions/json?origin="+ fromAddress +"&destination="+ toAddress +"&key=AIzaSyD55li1OuTm-bRAzfO4Mo3AsdNKHywfp1s")
	if err != nil {
		log.Panic(err)
	}
	decoder := json.NewDecoder(resp.Body)

	var direction Direction
	err = decoder.Decode(&direction)
	if err != nil {
		log.Panic(err)
	}
	return &direction

}

func DemoStaticMapFromPolyLine(polyline Polyline) string {
	return "https://maps.googleapis.com/maps/api/staticmap?size=400x400&path=weight:3%7Ccolor:blue%7Cenc:" + polyline.Points + "&key=AIzaSyD55li1OuTm-bRAzfO4Mo3AsdNKHywfp1s"
}

func GetHtmlInstructions(direction *Direction) []string {
	instructions := make([]string, 0)
	for _,leg := range direction.Routes[0].Legs {
		for _,step := range leg.Steps {
			instructions = append(instructions, step.Html_instructions)
		}
	}
	return instructions
}