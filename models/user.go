package models

import (
	"gorm.io/gorm"
)

type UserChurnProba struct {
	DtRef          string  `json:"dt_reference"`
	DescModelName  string  `json:"model_name"`
	NrModelVersion int     `json:"model_version"`
	IdUser         string  `json:"id_user" gorm:"primaryKey"`
	DescLabel      string  `json:"model_label"`
	NrProbLabel    float64 `json:"model_label_proba"`
	NrProbNorm     float64 `json:"model_label_proba_norm"`
	NrProbaRank    float64 `json:"model_label_proba_rank"`
}

func GetUser(idUser string, con *gorm.DB) (*UserChurnProba, error) {

	user := &UserChurnProba{}

	res := con.First(&user, "id_user = ?", idUser)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
