package main

import (
	"googlemaps"
	"net/http"
	"html/template"
	"log"
	"ElChatto"
)

func getIndex(w http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Panic(err)
	}

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if len(from) <= 0 || len(to) <= 0 {
		from = "Örebro"
		to = "Stockholm"
	}

	directions := googlemaps.GetDirectionRequest(from, to)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Image string
		Descriptions []string
	}{
		Image: googlemaps.GetStaticMapFromPolyLine(directions.Routes[0].Overview_polyline),
		Descriptions: googlemaps.GetHtmlInstructions(directions),
	}

	tpl.Execute(w, data)
}

func main() {

	log.Print("Starting ElChatto!")
	ElChatto.Start()


	http.HandleFunc("/", getIndex);
	http.ListenAndServe("localhost:8080", nil)
}
