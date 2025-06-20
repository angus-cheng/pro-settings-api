package db

import (
	"internal/db"

	"github.com/angus-cheng/pro-settings-api/internal/models"
)

func SavePlayer(p models.Player) error {
	_, err := db.DB.Exec(`INSERT INTO pro_settings(name, sensitivity)
						VALUES ($1, $2)`, p.Name, p.Sensitivity)
	return err
}