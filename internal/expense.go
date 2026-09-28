package expense

import (
	"errors"
	"time"
	"uuid"
)

var (
	ErrExpenseNotFound  = errors.New("expense not found")
	ErrInvalidExpenseID = errors.New("invalid id")
	ErrInvalidCategory  = errors.New("invalid category")
	ErrInvalidAmount    = errors.New("invalid amount")
)

type Money int64
type Expense struct {
	ID          uuid.UUID `json:"id"`
	Category    Category  `json:"category"`
	Amount      Money     `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewExpense(amount Money, category Category, description string) Expense {
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

func (e Expense) AmountToNaira() float64 {
	if e.Amount < 0 {
		return 0
	}
	return float64(e.Amount) / 100
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

func Validate(e Expense) error {
	var errs []error
	if e.ID.Compare(uuid.UUID{}) != 1 {
		errs = append(errs, ErrInvalidExpenseID)
	}
	if e.Amount <= 0 {
		errs = append(errs, ErrInvalidAmount)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
