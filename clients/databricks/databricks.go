package databricks

import (
	"database/sql"
	"os"
	"tmw_models/models"

	_ "github.com/databricks/databricks-sql-go"
	"github.com/joho/godotenv"
)

type DatabricksClient struct {
	Connection *sql.DB
}

func NewDatabricksClient() (*DatabricksClient, error) {
	godotenv.Load(".env")
	dsn := os.Getenv("DATABRICKS_DSN")
	con, err := sql.Open("databricks", dsn)
	if err != nil {
		return nil, err
	}

	client := &DatabricksClient{
		Connection: con,
	}

	return client, nil
}

func (client *DatabricksClient) GetChurnScore() ([]models.UserChurnProba, error) {

	sql := `
		WITH tb_score AS (

		SELECT *
		FROM feature_store.upsell.models
		WHERE descLabel = 1
		AND descModelName = 'feature_store.upsell.churn'
		QUALIFY row_number() OVER (PARTITION BY idCliente ORDER BY dtRef DESC) = 1

		),

		tb_pre_rank AS (

		SELECT *,
				(nrProbLabel - (select min(nrProbLabel) FROM tb_score)) / ((select max(nrProbLabel) FROM tb_score) - (select min(nrProbLabel) FROM tb_score)) AS NrProbNorm,
				row_number() over (order by nrProbLabel ASC) AS NrProbaRank

		FROM tb_score

		)

		SELECT
		DtRef,
		DescModelName,
		NrModelVersion,
		idCliente AS IdUser,
		DescLabel,
		NrProbLabel,
		NrProbNorm,
		NrProbaRank / (select count(*) FROM tb_pre_rank) AS NrProbaRank

		FROM tb_pre_rank`

	rows, err := client.Connection.Query(sql)
	if err != nil {
		return nil, err
	}

	data := []models.UserChurnProba{}
	user := models.UserChurnProba{}
	for rows.Next() {
		rows.Scan(&user.DtRef, &user.DescModelName, &user.NrModelVersion, &user.IdUser, &user.DescLabel, &user.NrProbLabel, &user.NrProbNorm, &user.NrProbaRank)
		data = append(data, user)
	}

	return data, nil
}

func (client *DatabricksClient) GetRetro() ([]models.UserRetro, error) {

	sql := `
		SELECT *
		FROM gold.teomewhy.retro_cliente_2024`

	rows, err := client.Connection.Query(sql)
	if err != nil {
		return nil, err
	}

	data := []models.UserRetro{}
	user := models.UserRetro{}
	for rows.Next() {
		rows.Scan(
			&user.IdUser,
			&user.QtDias,
			&user.QtPontosAcumulados,
			&user.QtPontosGastos,
			&user.QtChatMessages,
			&user.QtPresente,
			&user.DtPrimeiraTransacao,
			&user.QtDiasPrimtransacao,
			&user.QtTempoTotalHoras,
			&user.QtHorasDia,
			&user.RankPontos,
			&user.RankAntigo,
			&user.PctRankPontos,
			&user.PctRankAntigo,
		)
		data = append(data, user)
	}

	return data, nil
}
