package handlers

import (
	"net/http"
	"tmw_models/controllers"

	"github.com/gin-gonic/gin"
)

func GETUserChurnScore(c *gin.Context) {

	idUser := c.Param("id")

	probaRank, err := controllers.GetChurnScore(idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": idUser, "prob": probaRank})
}
