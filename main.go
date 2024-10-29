package main

import (
	"flag"
	"log"
	"tmw_models/db"
	"tmw_models/handlers"
	"tmw_models/models"
	"tmw_models/results"

	"github.com/gin-gonic/gin"
)

func main() {

	migration := flag.Bool("migrations", false, "Realizar migrations do banco de dados")
	flag.Parse()

	if *migration {
		log.Println("Executando migrations...")

		conMySQL, err := db.ConnectMySQL()
		if err != nil {
			log.Println(err)
		}

		conMySQL.AutoMigrate(&models.UserChurnProba{})

		log.Println("ok")
	}

	go results.AutoResults()

	r := gin.Default()
	r.GET("/churn_score/:id", handlers.GETUserChurnScore)
	r.Run("0.0.0.0:8080")
}
