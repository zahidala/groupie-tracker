// package main

// import "net/http"

// func main() {

// 	// Create a new request multiplexer
// 	// Take incoming requests and dispatch them to the matching handlers
// 	mux := http.NewServeMux()

// 	// Register the routes and handlers
// 	mux.Handle("/", &homeHandler{})

// 	// Run the server
// 	http.ListenAndServe(":8080", mux)
// }

// type homeHandler struct{}

// func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This is my home page"))
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Year         int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}

type LocationsPage struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesPage struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationsPage struct {
	ID        int      `json:"id"`
	Relations []string `json:"datesLocations"`
}

func main() {
	// Send GET request to the API endpoint
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}

	// Process the fetched data
	for ind, artist := range artists {
		if ind > 3 { //////////////////// del
			break
		}
		fmt.Println("ID:", artist.ID)
		fmt.Println("Artist:", artist.Name)
		fmt.Println("Image:", artist.Image)
		fmt.Println("Year:", artist.Year)
		fmt.Println("First album:", artist.FirstAlbum)
		fmt.Println("Members:")
		for _, member := range artist.Members {
			fmt.Println("-", member)
		}
		// LOCATIONS
		// get locations from link
		loc, err := http.Get(artist.LocationsURL)
		if err != nil {
			log.Fatal(err)
		}
		defer loc.Body.Close()
		// check response status code
		if loc.StatusCode != http.StatusOK {
			log.Fatalf("API locations request failed with status code: %d", loc.StatusCode)
		}
		// decode JSON locations
		var locationsPg LocationsPage
		err = json.NewDecoder(loc.Body).Decode(&locationsPg)
		if err != nil {
			log.Fatal(err)
		}
		// print locations
		fmt.Println("Locations:")
		for _, location := range locationsPg.Locations {
			fmt.Println("-", location)
		}
		// DATES
		// get dates from link
		datesJSON, err := http.Get(artist.DatesURL)
		if err != nil {
			log.Fatal(err)
		}
		defer datesJSON.Body.Close()
		// check response status code
		if datesJSON.StatusCode != http.StatusOK {
			log.Fatalf("API dates request failed with status code: %d", datesJSON.StatusCode)
		}
		// decode JSON dates
		var datesPg DatesPage
		err = json.NewDecoder(datesJSON.Body).Decode(&datesPg)
		if err != nil {
			log.Fatal(err)
		}
		// print dates
		fmt.Println("Dates:")
		for _, date := range datesPg.Dates {
			fmt.Println("-", date)
		}
		// RELATIONS
		// get relations from URL
		rel, err := http.Get(artist.RelationsURL)
		if err != nil {
			log.Fatal(err)
		}
		defer rel.Body.Close()
		// check response status code
		if rel.StatusCode != http.StatusOK {
			log.Fatalf("API relations request failed with status: %d", rel.StatusCode)
		}
		// decode json relations
		var relationsPg RelationsPage
		err = json.NewDecoder(rel.Body).Decode(&relationsPg)
		if err != nil {
			log.Fatal(err)
		}
		// print relations
		fmt.Println("Relations:")
		for _, relation := range relationsPg.Relations {
			fmt.Println("-", relation)
		}

		fmt.Println()
	}
	// fmt.Println(artists[0])
}
