package groupietracker

import (
	"encoding/json"
	"log"
	"net/http"
)

var artists []Artist

func FetchArtists() []Artist {
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
	// var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}
	return artists
}

func FetchLocations() []LocationsPage {
	// artists := FetchArtists()
	var locations []LocationsPage = make([]LocationsPage, len(artists))
	for ind, artist := range artists {
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
		err = json.NewDecoder(loc.Body).Decode(&locations[ind])
		if err != nil {
			log.Fatal(err)
		}

	}
	return locations
}

func FetchDates() []DatesPage {
	// artists := FetchArtists()
	var dates []DatesPage = make([]DatesPage, len(artists))
	// get dates from link
	for ind, artist := range artists {
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
		err = json.NewDecoder(datesJSON.Body).Decode(&dates[ind])
		if err != nil {
			log.Fatal(err)
		}
	}
	return dates
}

func FetchRelations() []RelationsPage {
	// artists := FetchArtists()
	var relations []RelationsPage = make([]RelationsPage, len(artists))
	for ind, artist := range artists {
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
		err = json.NewDecoder(rel.Body).Decode(&relations[ind])
		if err != nil {
			log.Fatal(err)
		}
	}
	return relations
}
