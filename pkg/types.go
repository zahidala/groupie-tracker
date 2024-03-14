package groupietracker

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
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type DisplayDetails struct {
	ArtistDetails Artist
	Concerts RelationsPage
}