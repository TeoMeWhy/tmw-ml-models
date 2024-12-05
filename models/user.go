package models

import (
	"time"

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

type UserRetro struct {
	IdUser              string    `json:"id_user" gorm:"primaryKey"`
	QtDias              int       `json:"qt_dias"`
	QtPontosAcumulados  int       `json:"qt_pontos_acumulados"`
	QtPontosGastos      int       `json:"qt_pontos_gastos"`
	QtChatMessages      int       `json:"qt_chatmessages"`
	QtPresente          int       `json:"qt_presente"`
	DtPrimeiraTransacao time.Time `json:"dt_primeira_transacao"`
	QtDiasPrimtransacao int       `json:"qt_dias_prim_transacao"`
	QtTempoTotalHoras   float64   `json:"qt_tempo_total_horas"`
	QtHorasDia          float64   `json:"qt_horas_dia"`
	RankPontos          int       `json:"rank_pontos"`
	RankAntigo          int       `json:"rank_antigo"`
	PctRankPontos       float64   `json:"pct_rank_pontos"`
	PctRankAntigo       float64   `json:"pct_rank_antigo"`
}

func GetUserChurn(idUser string, con *gorm.DB) (*UserChurnProba, error) {
	user := &UserChurnProba{}
	res := con.First(&user, "id_user = ?", idUser)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func GetUserRetro(idUser string, con *gorm.DB) (*UserRetro, error) {
	user := &UserRetro{}
	res := con.First(&user, "id_user = ?", idUser)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
