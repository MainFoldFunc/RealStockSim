package database

import (
	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

func CreatePortfolioDatabase(portfolio *structs.Portfolio) error {
	if err := DB.Create(&portfolio).Error; err != nil {
		return err
	}
	return nil
}
