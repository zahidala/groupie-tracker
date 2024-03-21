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

func errorHandler(res http.ResponseWriter, data pkg.ErrorPage) {
	res.WriteHeader(data.Error.Code)
	templates.ExecuteTemplate(res, "error.html", data)
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
		errorHandler(w, pkg.ErrorPage{
			Error: pkg.Error{
				Message: "No artists found.",
				Code:    status,
			},
			Data: data,
		})
		return
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	artistID := r.PathValue("id")

	details, status := pkg.FetchArtistByID(artistID)
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

	if status != 200 {
		errorHandler(w, pkg.ErrorPage{
			Error: pkg.Error{
				Message: "Artist not found.",
				Code:    status,
			},
			Data: data,
		})
		return
	}

	templates.ExecuteTemplate(w, "artist-details.html", data)
}
