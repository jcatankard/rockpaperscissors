package webpage

import (
	"encoding/json"
	"log"
	"myapp/rps"
	"net/http"
	"text/template"
)

// HomePage renders the home page on launch
func HomePage(w http.ResponseWriter, r *http.Request) {
	renderTempalate(w, "webpage/index.html")
}

// NewGame starts a new game
func NewGame(w http.ResponseWriter, r *http.Request) {
	rps.GAME = rps.Game{}
	log.Println("new game started successfully")
}

// UpdateScores writes back the up to date scores
func UpdateScores(w http.ResponseWriter, r *http.Request) {

	out, err := json.MarshalIndent(rps.GAME, "", "    ")
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		log.Println("score updated successfully")
	}
}

// PlayRound takes a player's turn as a response then writes back the results of the round
func PlayRound(w http.ResponseWriter, r *http.Request) {

	playerChoice := r.URL.Query().Get("c")
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		log.Println("round played successfully")
	}
}

// rederTemplate renders a given html file
func renderTempalate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
