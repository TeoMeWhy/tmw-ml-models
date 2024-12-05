package controllers

import (
	"fmt"
	"tmw_models/clients/aurora"
	"tmw_models/models"
)

var auroraClient, _ = aurora.NewAuroraClient()

func GetChurnScore(id string) (*float64, error) {
	user, err := models.GetUserChurn(id, auroraClient.Connection)
	if err != nil {
		return nil, err
	}
	return &user.NrProbaRank, nil
}

func GetRetro(id string) (*string, error) {
	user, err := models.GetUserRetro(id, auroraClient.Connection)
	if err != nil {
		return nil, err
	}

	text := `Dias ativos: %d | Pontos Ganhos: %d | Mensagens: %d | !presente: %d | Horas: %.1f`

	if user.RankPontos <= 10 {
		text += " | HIGHLIGHT: Top 10 em pontos!"
	} else if user.RankPontos <= 20 {
		text += " | HIGHLIGHT: Top 20 em pontos!"
	} else if user.RankPontos <= 50 {
		text += " | HIGHLIGHT: Top 50 em pontos!"
	} else if user.RankPontos <= 100 {
		text += " | HIGHLIGHT: Top 100 em pontos!"
	}

	text = fmt.Sprintf(
		text,
		user.QtDias,
		user.QtPontosAcumulados,
		user.QtChatMessages,
		user.QtPresente,
		user.QtTempoTotalHoras,
	)

	return &text, nil
}
