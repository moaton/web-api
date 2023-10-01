package revenue

import "time"

type Revenue struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Discription string    `json:"description"`
	Amount      int64     `json:"amount"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (r *Revenue) CreateRevenues() ([]Revenue, error) {
	return []Revenue{}, nil
}
