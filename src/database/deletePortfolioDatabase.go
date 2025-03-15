package database

import "github.com/MainFoldFunc/RealStockSim/src/structs"

func DeletePortfolioDatabase(userID uint) error {
	if err := DB.Where("user_id = ?", userID).Delete(&structs.Portfolio{}).Error; err != nil {
		return err
	}
	return nil
}
