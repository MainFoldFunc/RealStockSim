package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

// CreateStockDatabase adds a new stock to the database
func CreateStockDatabase(stock *structs.Stocks) error {
	if err := DB.Create(&stock).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePortfolioWithStock adds the created stock to the user's portfolio
func UpdatePortfolioWithStock(userID uint, stockName string, allAmount string) error {
	var portfolio structs.Portfolio

	// Fetch user's portfolio
	if err := DB.Where("user_id = ?", userID).First(&portfolio).Error; err != nil {
		return errors.New("portfolio not found")
	}

	// Extract the amount of the stock being added
	stockParts := strings.Split(allAmount, ",")
	var totalAmount int

	for _, part := range stockParts {
		var amount, price, ownerID int
		_, err := fmt.Sscanf(part, "%d:%d:%d", &amount, &price, &ownerID)
		if err != nil {
			return errors.New("invalid stock format in allAmount")
		}
		totalAmount += amount
	}

	// Append stock to portfolio's stocksInHand
	if portfolio.StocksInHand == "" {
		portfolio.StocksInHand = fmt.Sprintf("%s:%d", stockName, totalAmount)
	} else {
		portfolio.StocksInHand += fmt.Sprintf(",%s:%d", stockName, totalAmount)
	}

	// Save the updated portfolio
	if err := DB.Save(&portfolio).Error; err != nil {
		return errors.New("failed to update portfolio")
	}

	return nil
}
