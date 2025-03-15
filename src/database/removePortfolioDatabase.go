package database

import "github.com/MainFoldFunc/RealStockSim/src/structs"

func RemoveUserDatabase(userID uint) error {
	if err := DB.Where("id = ?", userID).Delete(&structs.Users{}).Error; err != nil {
		return err
	}

	return nil
}
