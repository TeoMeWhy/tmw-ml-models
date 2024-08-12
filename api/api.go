package main

import (
	"api/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var con, _ = db.Connect()

func GETUserChurnScore(c *gin.Context) {

	idUser := c.Param("id")

	score, err := db.GetChurnUserScore(idUser, con)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if score == -1. {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuario nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": idUser, "prob": fmt.Sprintf("%f", score)})
}

func main() {

	r := gin.Default()
	r.GET("/churn_score/:id", GETUserChurnScore)
	r.Run("0.0.0.0:8082")

}
