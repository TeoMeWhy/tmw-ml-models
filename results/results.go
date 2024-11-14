package results

import (
	"database/sql"
	"log"
	"time"
	"tmw_models/db"
	"tmw_models/models"

	"gorm.io/gorm"
)

func GetResults(con *sql.DB) ([]models.UserChurnProba, error) {

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

	rows, err := con.Query(sql)
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

func InsertResults(users []models.UserChurnProba, con *gorm.DB) error {

	for _, u := range users {
		res := con.Save(u)
		if res.Error != nil {
			return res.Error
		}
		res.Commit()
	}

	return nil

}

func AutoResults() error {

	for {

		conMySQL, err := db.ConnectMySQL()
		if err != nil {
			log.Println("Erro ao conectar o MySQL:", err)
			return err
		}

		conDatabricks, err := db.ConnectDatabricks()
		if err != nil {
			log.Println("Erro ao conectar o databricks:", err)
			return err
		}

		data, err := GetResults(conDatabricks)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Minute * 5)
			continue
		}

		if err := InsertResults(data, conMySQL); err != nil {
			log.Println(err)
		}

		conDatabricks.Close()
		time.Sleep(time.Hour * 12)
	}

}
