package groupietracker

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchArtists(search string) (fetchedArtists []Artist, statusCode int) {
	// Send GET request to the API endpoint
	var artists []Artist

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
		return artists, resp.StatusCode
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d", resp.StatusCode)
		return artists, resp.StatusCode
	}

	// Decode the JSON response
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
		return artists, http.StatusInternalServerError
	}

	return FilterSearchedArtists(artists, search), resp.StatusCode
}

func FetchArtistByID(id string) (fetchedArtist Artist, statusCode int) {
	var artist Artist

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		log.Fatal(err)
		return artist, resp.StatusCode
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d", resp.StatusCode)
		return artist, resp.StatusCode
	}

	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		log.Fatal(err)
		return artist, http.StatusInternalServerError
	}

	// Since the API fails to return a 404 status code when the artist is not found,
	// we need to check if the ID is 0 to determine if the artist was not found
	// as the API returns ID of 0 for non-existent artists along with empty strings for other fields

	if artist.ID == 0 {
		return artist, http.StatusNotFound
	}

	return artist, resp.StatusCode
}

func FetchRelationsByID(id string) RelationsPage {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d", resp.StatusCode)
	}

	var relation RelationsPage
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		log.Fatal(err)
	}

	return relation
}

// Might use later
/*
func FetchArtistByID(id string) Artist {
	// Fetch from []Artists
	artists := FetchArtists()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	artistByID := artists[idInt-1]
	return artistByID
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

func FetchLocationsByArtistID(id string) LocationsPage {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d", resp.StatusCode)
	}

	var location LocationsPage
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		log.Fatal(err)
	}

	return location
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
*/
