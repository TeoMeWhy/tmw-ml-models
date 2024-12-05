package controllers

import (
	"log"
	"time"
	"tmw_models/clients/aurora"
	"tmw_models/clients/databricks"
)

func ChurnController() error {

	for {

		databricksClient, err := databricks.NewDatabricksClient()
		if err != nil {
			return err
		}

		auroraClient, err := aurora.NewAuroraClient()
		if err != nil {
			return err
		}

		data, err := databricksClient.GetChurnScore()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Minute * 5)
			continue
		}

		if err := auroraClient.InsertChurn(data); err != nil {
			log.Println(err)
		}

		databricksClient.Connection.Close()
		time.Sleep(time.Hour * 12)
	}

}

func RetroController() error {

	for {

		databricksClient, err := databricks.NewDatabricksClient()
		if err != nil {
			return err
		}

		auroraClient, err := aurora.NewAuroraClient()
		if err != nil {
			return err
		}

		data, err := databricksClient.GetRetro()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Minute * 5)
			continue
		}

		if err := auroraClient.InsertRetro(data); err != nil {
			log.Println(err)
		}

		databricksClient.Connection.Close()
		time.Sleep(time.Hour * 12)
	}
}
