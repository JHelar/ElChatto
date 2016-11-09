package main

import (
	"fmt"
	"googlemaps"
	"net/http"
)

func getMap(w http.ResponseWriter, r *http.Request){
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if len(from) <= 0 || len(to) <= 0 {
		from = "Örebro"
		to = "Stockholm"
	}
	directions := googlemaps.DemoDirectionRequest(from, to)
	fmt.Printf("Dir: (%+v)", directions)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := "<img src='" + googlemaps.DemoStaticMapFromPolyLine(directions.Routes[0].Overview_polyline) + "' /><ol>"
	for _,instruction := range googlemaps.GetHtmlInstructions(directions){
		html += "<li>" + instruction + "</li>"
	}
	html += "</ol>"
	fmt.Fprint(w, html)
}

func main() {

	http.HandleFunc("/", getMap);
	http.ListenAndServe("localhost:8080", nil)
}
