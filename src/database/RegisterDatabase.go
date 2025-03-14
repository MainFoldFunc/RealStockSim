package database

import (
	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

func RegisterDatabase(user *structs.Users) error {
	err := DB.Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}
