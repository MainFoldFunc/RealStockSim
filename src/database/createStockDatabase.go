package database

import "github.com/MainFoldFunc/RealStockSim/src/structs"

func CreateStockDatabase(stock *structs.Stocks) error {
	if err := DB.Create(&stock).Error; err != nil {
		return err
	}

	return nil
}
