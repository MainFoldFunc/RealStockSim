package database

import (
	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

func GetUserByEmail(email string) (*structs.Users, error) {
	var user structs.Users

	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
