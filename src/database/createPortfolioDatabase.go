package database

import (
	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

func CreatePortfolioDatabase(portfolio *structs.Portfolio) error {
	err := DB.Create(&portfolio)
	if err != nil {
		return err.Error
	}
	return nil
}
