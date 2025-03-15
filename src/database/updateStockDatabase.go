package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

// UpdateStockDatabase processes the stock update and modifies the amount
func UpdateStockDatabase(stock *structs.UpdateStock) error {
	var dbStock structs.Stocks

	// Fetch the stock based on its ID or Name
	if err := DB.Where("id = ?", stock.ID).First(&dbStock).Error; err != nil {
		return errors.New("stock not found")
	}

	// Parse AllAmount field (format: "amount:price:userID")
	stockAmounts := strings.Split(dbStock.AllAmount, ",")

	var updatedStockAmounts []string
	stockFound := false

	// Iterate over each stock entry in AllAmount
	for _, stockAmount := range stockAmounts {
		var amount, price, userID int
		// Extract the amount, price, and user ID from the format "amount:price:userID"
		_, err := fmt.Sscanf(stockAmount, "%d:%d:%d", &amount, &price, &userID)
		if err != nil {
			return errors.New("invalid stock format in allAmount")
		}

		// Check if the stock price matches the provided price
		if uint(price) == stock.Price {
			// Add the specified amount to the stock
			amount += int(stock.AmountToAdd)
			// Mark stock as found to update
			stockFound = true
		}

		// Add the updated or unchanged stock entry back to the list
		updatedStockAmounts = append(updatedStockAmounts, fmt.Sprintf("%d:%d:%d", amount, price, userID))
	}

	// If the stock was not found for the given price, return an error
	if !stockFound {
		return errors.New("stock with the specified price not found")
	}

	// Update the AllAmount field with the updated stock list
	dbStock.AllAmount = strings.Join(updatedStockAmounts, ",")

	// Save the updated stock back to the database
	if err := DB.Save(&dbStock).Error; err != nil {
		return errors.New("failed to update stock quantity")
	}

	return nil
}
