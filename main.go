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
	http.Handle("GET /static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("GET /", homeHandler)
	http.HandleFunc("GET /artist/{id}", artistDetailsHandler)
	fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func errorHandler(res http.ResponseWriter, data pkg.ErrorPageProps) {
	res.WriteHeader(data.Error.Code)

	err := templates.ExecuteTemplate(res, "error.html", data)

	if err != nil {
		log.Println(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	search := params.Get("q")
	search = strings.TrimSpace(search)

	artists, status := pkg.FetchArtists(search)

	title := "Groupie Tracker - Artists"

	// Create a map to hold the data to pass to the template
	data := map[string]interface{}{
		"Title":   title,
		"Artists": artists,
		"Search":  search,
	}

	if status != 200 {
		errorHandler(w, pkg.ErrorPageProps{
			Error: pkg.Error{
				Message: "API request failed - No artists found.",
				Code:    status,
			},
			Title: "Groupie Tracker - No Artists Found",
		})
		return
	}

	if r.URL.Path != "/" {
		errorHandler(w, pkg.ErrorPageProps{
			Error: pkg.Error{
				Message: "Page not found.",
				Code:    404,
			},
			Title: "Groupie Tracker - Page Not Found",
		})
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		log.Println(err)
	}
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	artistID := r.PathValue("id")

	details, detailStatusCode := pkg.FetchArtistByID(artistID)
	relations, relationStatusCode := pkg.FetchRelationsByID(artistID)
	artistDescription := pkg.FetchArtistDescriptionByName(details.Name)

	title := "Groupie Tracker - " + func() string {
		if details.Name != "" {
			return details.Name
		}
		return "Artist Details"
	}()

	displayDetails := pkg.DisplayDetails{
		ArtistDetails:      details,
		Concerts:           relations,
		RelationStatusCode: relationStatusCode,
		ArtistDescription:  artistDescription,
	}

	data := map[string]interface{}{
		"Title":             title,
		"DisplayDetails":    displayDetails,
		"DetailsPage":       true,
		"Background":        displayDetails.ArtistDetails.Image,
		"AritstDescription": artistDescription,
	}

	if detailStatusCode != 200 {
		errorHandler(w, pkg.ErrorPageProps{
			Error: pkg.Error{
				Message: "API request failed - Artist not found.",
				Code:    detailStatusCode,
			},
			Title: "Groupie Tracker - Artist Not Found",
		})
		return
	}

	err := templates.ExecuteTemplate(w, "artist-details.html", data)

	if err != nil {
		log.Println(err)
	}
}
