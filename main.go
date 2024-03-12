package main

import (
	"fmt"
	pkg "groupietracker/pkg"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// artists := pkg.FetchArtists()
	// locations := pkg.FetchLocations()
	// dates := pkg.FetchDates()
	// relations := pkg.FetchRelations()

	// fmt.Println(artists[0].Name)
	// fmt.Println(locations[0].Locations)
	// fmt.Println(dates[0])
	// fmt.Println(relations[0].Relations)

	// server
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistDetailsHandler)
	fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	artists := pkg.FetchArtists()

	templates.ExecuteTemplate(w, "index.html", artists)
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	artistID := params.Get("id")

	locations := pkg.FetchLocationsByArtistID(artistID)

	templates.ExecuteTemplate(w, "artist-details.html", locations)
}
