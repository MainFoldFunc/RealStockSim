package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

// BuyStockDatabase processes the stock purchase and updates the portfolio
func BuyStockDatabase(whatToBuy *structs.BuyingStocks, userID uint) error {
	var stock structs.Stocks
	var portfolio structs.Portfolio

	// Check if the stock exists
	if err := DB.Where("name = ?", whatToBuy.Name).First(&stock).Error; err != nil {
		return errors.New("stock not found")
	}

	// Check if the portfolio exists
	if err := DB.Where("user_id = ?", userID).First(&portfolio).Error; err != nil {
		return errors.New("portfolio not found")
	}

	// Parse the allAmount field to get individual stock-price pairs
	stockAmounts := strings.Split(stock.AllAmount, ",")
	var totalAvailableAmount int
	var totalCost uint

	// Check availability of stock in different price points and calculate total cost
	for _, stockAmount := range stockAmounts {
		// Split stockAmount to get amount and price
		var availableAmount, price int
		_, err := fmt.Sscanf(stockAmount, "%d:%d", &availableAmount, &price)
		if err != nil {
			return errors.New("invalid stock format in allAmount")
		}

		// If the stock price is within the max price, check availability
		if price <= int(whatToBuy.MaxPrice) {
			if availableAmount >= int(whatToBuy.Amount) {
				// We can fulfill this order
				totalAvailableAmount = availableAmount
				totalCost = uint(whatToBuy.Amount * uint(price))
				break
			}
		}
	}

	// If no matching stock price was found or stock was insufficient
	if totalAvailableAmount == 0 {
		return errors.New("not enough stock available or price too high")
	}

	// Check if user has enough money
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

	// Now deduct the purchased stock amount from the available stock
	// Update the stock quantities in the database
	updatedStockAmounts := []string{}
	for _, stockAmount := range stockAmounts {
		var availableAmount, price int
		_, err := fmt.Sscanf(stockAmount, "%d:%d", &availableAmount, &price)
		if err != nil {
			return errors.New("invalid stock format in allAmount")
		}

		// Deduct the purchased amount from the correct price point
		if price <= int(whatToBuy.MaxPrice) {
			if availableAmount >= int(whatToBuy.Amount) {
				availableAmount -= int(whatToBuy.Amount)
				// Update the stock amount for this price point
				updatedStockAmounts = append(updatedStockAmounts, fmt.Sprintf("%d:%d", availableAmount, price))
				break
			}
		}
	}

	// Save the updated stock with new amounts
	stock.AllAmount = strings.Join(updatedStockAmounts, ",")
	if err := DB.Save(&stock).Error; err != nil {
		return errors.New("failed to update stock quantity")
	}

	return nil
}
