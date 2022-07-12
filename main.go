package main

import (
	"log"
	"myapp/webpage"
	"net/http"
)

func main() {

	http.HandleFunc("/newgame", webpage.NewGame)
	http.HandleFunc("/updatescores", webpage.UpdateScores)
	http.HandleFunc("/play", webpage.PlayRound)
	http.HandleFunc("/", webpage.HomePage)

	log.Println("starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
