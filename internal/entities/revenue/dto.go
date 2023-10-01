package revenue

import "time"

type CreateRevenueDTO struct {
	Name        string `json:"name"`
	Discription string `json:"description"`
	Amount      int64  `json:"amount"`
	Type        string `json:"type"`
}

type UpdateRevenueDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Discription string    `json:"description"`
	Amount      int64     `json:"amount"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
