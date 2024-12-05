package main

import (
	"flag"
	"log"
	"tmw_models/clients/aurora"
	"tmw_models/controllers"
	"tmw_models/handlers"
	"tmw_models/models"

	"github.com/gin-gonic/gin"
)

func main() {

	migration := flag.Bool("migrations", false, "Realizar migrations do banco de dados")
	flag.Parse()

	if *migration {
		log.Println("Executando migrations...")

		auroraClient, err := aurora.NewAuroraClient()
		if err != nil {
			log.Println(err)
		}

		auroraClient.Connection.AutoMigrate(&models.UserChurnProba{}, &models.UserRetro{})
		log.Println("ok")
		return
	}

	go controllers.ChurnController()
	go controllers.RetroController()

	r := gin.Default()
	r.GET("/churn_score/:id", handlers.GETUserChurnScore)
	r.GET("/retro_2024/:id", handlers.GETUserRetro)
	r.Run("0.0.0.0:8080")
}
