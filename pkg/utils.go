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
		// artistAdded := false // to avoid duplicates

		if strings.Contains(strings.ToLower(artist.Name), search) || strings.Contains(strconv.Itoa(artist.Year), search) || strings.Contains(artist.FirstAlbum, search) {
			filteredArtists = append(filteredArtists, artist)
			continue
		}

		for _, member := range artist.Members {
			if member == search || strings.Contains(strings.ToLower(member), search) {
				filteredArtists = append(filteredArtists, artist)
				// artistAdded = true
				break
			}
		}

		// if artistAdded {
		// 	continue
		// }

		// relationsMap := FetchRelationsByID(strconv.Itoa(artist.ID)).Relations // main - line 52
		// // check if key (location) matches search
		// for key, _ := range relationsMap {
		// 	if key == search || strings.Contains(strings.ToLower(key), search) {
		// 		filteredArtists = append(filteredArtists, artist)
		// 		artistAdded = true
		// 		break
		// 	}
		// }
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

// func FixLocations(loc string) string {
// 	loc = strings.ToUpper(string(loc[0]))
// 	loc = strings.ReplaceAll(loc, "-", ", ")
// 	countryIndex := strings.Index(loc, ", ") + 1
// 	loc = strings.ToUpper(string(loc[countryIndex]))
// 	return loc
// }
