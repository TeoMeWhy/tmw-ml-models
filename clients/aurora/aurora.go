package aurora

import (
	"fmt"
	"os"
	"tmw_models/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuroraClient struct {
	Connection *gorm.DB
}

func NewAuroraClient() (*AuroraClient, error) {

	godotenv.Load(".env")
	HOST_DB := os.Getenv("HOST_DB")
	PORT_DB := os.Getenv("PORT_DB")
	USER_DB := os.Getenv("USER_DB")
	PASSWORD_DB := os.Getenv("PASSWORD_DB")

	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, USER_DB, PASSWORD_DB, HOST_DB, PORT_DB, "ml_models")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	client := &AuroraClient{
		Connection: db,
	}

	return client, err
}

func (client *AuroraClient) InsertChurn(users []models.UserChurnProba) error {
	for _, u := range users {
		res := client.Connection.Save(u)
		if res.Error != nil {
			return res.Error
		}
		res.Commit()
	}
	return nil
}

func (client *AuroraClient) InsertRetro(users []models.UserRetro) error {
	for _, u := range users {
		res := client.Connection.Save(u)
		if res.Error != nil {
			return res.Error
		}
		res.Commit()
	}
	return nil
}
