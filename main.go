package main

import (
	"myapp/webpage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	http.HandleFunc("/newgame", webpage.NewGame)
	http.HandleFunc("/updatescores", webpage.UpdateScores)
	http.HandleFunc("/play", webpage.PlayRound)
	http.HandleFunc("/", webpage.HomePage)

	// port := ":8080"
	// log.Printf("starting web server on port %s\n", port)
	// http.ListenAndServe(port, nil)

	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
