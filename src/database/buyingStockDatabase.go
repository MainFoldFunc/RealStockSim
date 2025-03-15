package database

import (
	"errors"
	"fmt"

	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

// BuyStockDatabase processes the stock purchase and updates the portfolio
func BuyStockDatabase(whatToBuy *structs.BuyingStocks, userID uint) error {
	var stock structs.Stocks
	var portfolio structs.Portfolio

	// Check if the stock exists and is within the max price
	if err := DB.Where("name = ? AND curr_price <= ?", whatToBuy.Name, whatToBuy.MaxPrice).First(&stock).Error; err != nil {
		return errors.New("stock not found or price too high")
	}

	// Fetch user's portfolio
	if err := DB.Where("user_id = ?", userID).First(&portfolio).Error; err != nil {
		return errors.New("portfolio not found")
	}

	// Check if user has enough money
	totalCost := whatToBuy.Amount * stock.CurrPrice
	if portfolio.Money < totalCost {
		return errors.New("not enough money to buy stock")
	}

	// Deduct money from the portfolio
	portfolio.Money -= totalCost

	// Update StocksInHand (append new stock if not present)
	if portfolio.StocksInHand == "" {
		portfolio.StocksInHand = fmt.Sprintf("%s:%d", whatToBuy.Name, whatToBuy.Amount)
	} else {
		portfolio.StocksInHand += fmt.Sprintf(",%s:%d", whatToBuy.Name, whatToBuy.Amount)
	}

	// Save the updated portfolio
	if err := DB.Save(&portfolio).Error; err != nil {
		return errors.New("failed to update portfolio")
	}

	return nil
}
