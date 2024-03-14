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
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistDetailsHandler)
	fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	artists := pkg.FetchArtists()

	title := "Groupie Tracker - Artists"

	// Create a map to hold the data to pass to the template
	data := map[string]interface{}{
		"Title":   title,
		"Artists": artists,
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	artistID := params.Get("id")

	details := pkg.FetchArtistByID(artistID)
	relations := pkg.FetchRelationsByID(artistID)

	title := "Groupie Tracker" + " - " + details.Name

	displayDetails := pkg.DisplayDetails{
		ArtistDetails: details,
		Concerts:      relations,
	}

	data := map[string]interface{}{
		"Title":          title,
		"DisplayDetails": displayDetails,
	}

	templates.ExecuteTemplate(w, "artist-details.html", data)
}
