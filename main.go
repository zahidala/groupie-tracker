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
	"fmt"
	pkg "groupietracker/pkg"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	artists := pkg.FetchArtists()
	locations := pkg.FetchLocations()
	dates := pkg.FetchDates()
	relations := pkg.FetchRelations()

	fmt.Println("time: ", time.Since(start))

	fmt.Println(artists)
	fmt.Println(locations[0].Locations)
	fmt.Println(dates[0])
	fmt.Println(relations[0].Relations)

	// server
	http.HandleFunc("/", homeHandler)
	fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Groupie Tracker")
}

// // print relations
// fmt.Println("Relations:")
// for key, values := range relations[0].Relations {
// 	fmt.Println(key, ":")
// 	for _, value := range values {
// 		fmt.Println("-", value)
// 	}
// }
