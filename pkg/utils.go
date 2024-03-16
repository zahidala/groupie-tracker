package groupietracker

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

func FilterSearchedArtists(artists []Artist, search string) []Artist {
	sort.Slice(artists, func(i, j int) bool { return artists[i].Name < artists[j].Name })

	var filteredArtists []Artist

	if search == "" {
		return artists
	}

	for _, artist := range artists {
		if artist.Name == search || strings.Contains(strings.ToLower(artist.Name), strings.ToLower(search)) {
			filteredArtists = append(filteredArtists, artist)
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
