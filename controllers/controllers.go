package controllers

import (
	"tmw_models/db"
	"tmw_models/models"
)

var conMySQL, _ = db.ConnectMySQL()

func GetChurnScore(id string) (*float64, error) {
	user, err := models.GetUser(id, conMySQL)
	if err != nil {
		return nil, err
	}
	return &user.NrProbaRank, nil
}
