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
	User         string `json:"User" gorm:"unique;not null"`
	Money        uint   `json:"money" gorm:"unique;not null"`
	StocksInHand string `json:"stocksinhand" gorm:"unique;not null"`
}
