package main

import (
	"fmt"
	pkg "groupietracker/pkg"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	params := r.URL.Query()

	search := params.Get("q")
	search = strings.TrimSpace(search)

	artists := pkg.FetchArtists(search)

	title := "Groupie Tracker - Artists"

	// Create a map to hold the data to pass to the template
	data := map[string]interface{}{
		"Title":   title,
		"Artists": artists,
		"Search":  search,
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	artistID := params.Get("id")

	details := pkg.FetchArtistByID(artistID)
	relations := pkg.FetchRelationsByID(artistID)
	artistDescription := pkg.FetchArtistDescriptionByName(details.Name)

	title := "Groupie Tracker - " + func() string {
		if details.Name != "" {
			return details.Name
		}
		return "Artist Details"
	}()

	displayDetails := pkg.DisplayDetails{
		ArtistDetails:     details,
		Concerts:          relations,
		ArtistDescription: artistDescription,
	}

	data := map[string]interface{}{
		"Title":             title,
		"DisplayDetails":    displayDetails,
		"DetailsPage":       true,
		"Background":        displayDetails.ArtistDetails.Image,
		"AritstDescription": artistDescription,
	}

	templates.ExecuteTemplate(w, "artist-details.html", data)
}
