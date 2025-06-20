package db

import (
	"github.com/angus-cheng/pro-settings-api/internal/models"
)

func SavePlayer(p models.Player) error {
	_, err := DB.Exec(`INSERT INTO pro_settings(name, sensitivity)
						VALUES ($1, $2)`, p.Name, p.Sensitivity)
	return err
}