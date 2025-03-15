package structs

type Users struct {
	ID           uint    `json:"id"`
	Email        string  `json:"email" gorm:"unique;not null"`
	Password     string  `json:"password" gorm:"not null"`
	PortfolioKey *string `json:"portfoliokey"` // Pointer to make it nullable
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Portfolio struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id" gorm:"not null"`
	Money        uint   `json:"money"`
	StocksInHand string `json:"stocksinhand"`
}

type Stocks struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CurrPrice uint   `json:"currPrice"`
	AllAmount string `json:"allAmount"`
}

/*
	10 : 1000 amount : price
*/

type BuyingStocks struct {
	Name     string `json:"name"`
	Amount   uint   `json:"amount"`
	MaxPrice uint   `json:"maxPrice"`
}
