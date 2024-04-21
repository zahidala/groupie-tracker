package groupietracker

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FilterSearchedArtists(artists []Artist, search string) []Artist {
	sort.Slice(artists, func(i, j int) bool { return artists[i].Name < artists[j].Name })

	var filteredArtists []Artist

	if search == "" {
		return artists
	}

	search = strings.ToLower(search)

	for _, artist := range artists {
		artistAdded := false // to avoid duplicates

		if strings.Contains(strings.ToLower(artist.Name), search) || strings.Contains(strconv.Itoa(artist.Year), search) || strings.Contains(artist.FirstAlbum, search) {
			filteredArtists = append(filteredArtists, artist)
			continue
		}

		for _, member := range artist.Members {
			if member == search || strings.Contains(strings.ToLower(member), search) {
				filteredArtists = append(filteredArtists, artist)
				artistAdded = true
				break
			}
		}

		if artistAdded {
			continue
		}

		relationsMap := FetchRelationsByID(strconv.Itoa(artist.ID)).Relations // main - line 52
		// check if key (location) matches search
		for key, _ := range relationsMap {
			if key == search || strings.Contains(strings.ToLower(key), search) {
				filteredArtists = append(filteredArtists, artist)
				artistAdded = true
				break
			}
		}
	}

	return filteredArtists
}

func FetchArtistDescriptionByName(name string) string {
	file, _ := os.ReadFile("./constants/artist-descriptions.json")

	var descriptions []ArtistDescription

	err := json.Unmarshal(file, &descriptions)
	if err != nil {
		return "No description available."
	}

	for _, artist := range descriptions {
		if artist.Name == name {
			return artist.Description
		}
	}

	return "No description available."
}

/* ChatGPT
func FilterSearchedArtists(artists []Artist, search string) []Artist {
	// Convert search string to lowercase for case-insensitive comparisons
	search = strings.ToLower(search)

	var filteredArtists []Artist

	// Iterate through each artist
	for _, artist := range artists {
		// Check if any criteria matches the search
		if artistMatchesSearchCriteria(artist, search) {
			// Add artist to filteredArtists if not already added
			if !isArtistAlreadyAdded(filteredArtists, artist) {
				filteredArtists = append(filteredArtists, artist)
			}
		}
	}

	return filteredArtists
}

// Helper function to check if an artist matches the search criteria
func artistMatchesSearchCriteria(artist Artist, search string) bool {
	return strings.Contains(strings.ToLower(artist.Name), search) ||
		strings.Contains(strconv.Itoa(artist.Year), search) ||
		strings.Contains(artist.FirstAlbum, search) ||
		artistHasMatchingMember(artist, search) ||
		artistPerformedAtLocation(artist, search)
}

// Helper function to check if an artist is already in the filteredArtists slice
func isArtistAlreadyAdded(filteredArtists []Artist, artist Artist) bool {
	for _, a := range filteredArtists {
		if a.ID == artist.ID {
			return true
		}
	}
	return false
}

// Other helper functions for band members and locations matching...
*/
