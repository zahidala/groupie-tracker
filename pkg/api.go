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
		log.Println(err)
		return artists, resp.StatusCode
	}
	defer resp.Body.Close()

	
	// Check the response status code
	// this returns for api errors
	if resp.StatusCode != http.StatusOK {
		log.Println("Artists API request failed with status code:", resp.StatusCode)
		return artists, resp.StatusCode
	}
	
	// Decode the JSON response
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Println(err)
		return artists, http.StatusInternalServerError
	}

	return FilterSearchedArtists(artists, search), resp.StatusCode
}

func FetchLocations(artists []Artist) (artistsWLocations []Artist, statusCode int) {
	for _, artist := range artists {
		resp, err := http.Get(artist.LocationsURL)
		if err != nil {
			log.Println(err)
			return artists, resp.StatusCode
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("Artists API request failed with status code:", resp.StatusCode)
			return artists, resp.StatusCode
		}
		
		err = json.NewDecoder(resp.Body).Decode(&artist.Locations)
		if err != nil {
			log.Println(err)
			return artists, http.StatusInternalServerError
		}
	}
	return artist, resp.StatusCode
}

func FetchArtistByID(id string) (fetchedArtist Artist, statusCode int) {
	var artist Artist

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		log.Println(err)
		return artist, resp.StatusCode
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Artist Details (ID: %s) API request failed with status code: %d\n", id, resp.StatusCode)
		return artist, resp.StatusCode
	}

	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		log.Println(err)
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

func FetchRelationsByID(id string) (relation RelationsPage, statusCode int) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Relations (ID: %s) API request failed with status code: %d\n", id, resp.StatusCode)
		return relation, resp.StatusCode
	}

	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		log.Println(err)
	}

	return relation, resp.StatusCode
}
