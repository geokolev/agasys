package utils

import (
	"errors"

	"github.com/jonsch318/royalafg/services/poker/models"
)

func SearchByID(list []models.Player, id string) (*models.Player, int, error) {
	for i, player := range list {
		if player.ID == id {
			return &player, i, nil
		}
	}
	return nil, -1, errors.New("player not found")
}
