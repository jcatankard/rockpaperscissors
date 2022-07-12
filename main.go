package main

import (
	"log"
	"myapp/webpage"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/newgame", webpage.NewGame)
	http.HandleFunc("/updatescores", webpage.UpdateScores)
	http.HandleFunc("/play", webpage.PlayRound)
	http.HandleFunc("/", webpage.HomePage)

	port := os.Getenv("PORT")
	// port = ":8080"
	log.Printf("starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}
