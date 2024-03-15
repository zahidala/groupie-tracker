package groupietracker

import (
	"sort"
	"strings"
)

func FilterSearchedArtists(artists []Artist, search string) []Artist {
	search = strings.TrimSpace(search)
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
