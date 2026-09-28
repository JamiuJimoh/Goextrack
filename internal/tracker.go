package expense

import (
	"fmt"
	"time"
	"uuid"
)

// ExpenseRepository handles data
type ExpenseRepository interface {
	// Next() (Expense, error) // one at a time.
	All() ([]Expense, error)
	Save([]Expense) error
}

// Tracker performs business logic operations on the Expenses data
type Tracker struct {
	expenses []Expense
	repo     ExpenseRepository
}

func NewTracker(repo ExpenseRepository) (*Tracker, error) {
	e, err := repo.All()
	if err != nil {
		return nil, fmt.Errorf("error Tracker.New: %w", err)
	}
	return &Tracker{
		expenses: e,
		repo:     repo,
	}, nil
}

func (t *Tracker) Add(e Expense) error {
	err := Validate(e)
	if err != nil {
		return err
	}
	t.expenses = append(t.expenses, e)
	return nil
}

func (t *Tracker) List() []Expense {
	expenses := make([]Expense, len(t.expenses))
	copy(expenses, t.expenses)
	return expenses
}

func (t *Tracker) Get(id uuid.UUID) (Expense, error) {
	for _, e := range t.expenses {
		if e.ID == id {
			return e, nil
		}
	}
	return Expense{}, fmt.Errorf("in Tracker.Get: expense with id=%s: %w", id, ErrExpenseNotFound)
}

func (t *Tracker) Edit(e Expense) error {
	for i, expense := range t.expenses {
		if e.ID == expense.ID {
			e.UpdatedAt = time.Now()
			t.expenses[i] = e
			return nil
		}
	}
	return fmt.Errorf("in Tracker.Edit: expense with id=%s: %w", e.ID, ErrExpenseNotFound)
}

func (t *Tracker) Delete(id uuid.UUID) error {
	for i, expense := range t.expenses {
		if id == expense.ID {
			t.expenses = append(t.expenses[:i], t.expenses[i+1:]...)
			err := t.Save()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("in Tracker.Delete: expense with id=%s: %w", id, ErrExpenseNotFound)
}

func (t *Tracker) Summarize() Summary {
	count := len(t.expenses)
	var total Money
	for _, e := range t.expenses {
		total += e.Amount
	}
	var average float64

	if count > 0 {
		average = float64(total) / float64(count)
	}

	return Summary{
		Count:   count,
		Total:   total,
		Average: average,
	}
}

func (t *Tracker) Save() error {
	return t.repo.Save(t.expenses)
}

type Summary struct {
	Count   int
	Total   Money
	Average float64
}

func (s Summary) TotalToNaira() float64 {
	if s.Total < 0 {
		return 0
	}
	return float64(s.Total) / 100
}
