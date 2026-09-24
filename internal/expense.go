package expense

import (
	"errors"
	"time"
	"uuid"
)

var (
	ErrExpenseNotFound = errors.New("expense not found")
)

type Expense struct {
	ID          uuid.UUID `json:"id"`
	Category    Category  `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewExpense(amount float64, category Category, description string) Expense {
	now := time.Now()
	return Expense{
		ID:          uuid.NewV7(),
		Category:    category,
		Amount:      amount,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type Category string

const (
	food          string = "food"
	transport     string = "transport"
	housing       string = "housing"
	utilities     string = "utilities"
	health        string = "health"
	education     string = "education"
	entertainment string = "entertainment"
	shopping      string = "shopping"
	other         string = "other"
)

func AllCategories() []Category {
	return []Category{
		Category(food),
		Category(transport),
		Category(housing),
		Category(utilities),
		Category(health),
		Category(education),
		Category(entertainment),
		Category(shopping),
		Category(other),
	}
}

func (c Category) String() string {
	return string(c)
}
