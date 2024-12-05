package handlers

import (
	"log"
	"net/http"
	"tmw_models/controllers"

	"github.com/gin-gonic/gin"
)

func GETUserRetro(c *gin.Context) {

	idUser := c.Param("id")
	text, err := controllers.GetRetro(idUser)
	if err != nil {

		log.Println(text)
		log.Println(err)

		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(*text)
	c.JSON(http.StatusOK, gin.H{"uuid": idUser, "msg": text})
}
