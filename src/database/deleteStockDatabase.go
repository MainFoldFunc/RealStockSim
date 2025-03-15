package database

import "github.com/MainFoldFunc/RealStockSim/src/structs"

func DeleteStockDatabase(stockID uint) error {
	if err := DB.Where("id = ?", stockID).Delete(&structs.Stocks{}).Error; err != nil {
		return err
	}
	return nil
}
