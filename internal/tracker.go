package expense

import (
	"fmt"
	"uuid"
)

// ExpenseSourceHandler handles data source
type ExpenseSourceHandler interface {
	// Next() (Expense, error) // one at a time.
	All() ([]Expense, error)
	Save([]Expense) error
}

// Tracker performs business logic operations on the Expenses data
type Tracker struct {
	expenses      []Expense
	sourceHandler ExpenseSourceHandler
}

func NewTracker(handler ExpenseSourceHandler) (*Tracker, error) {
	e, err := handler.All()
	if err != nil {
		return nil, fmt.Errorf("error Tracker.New: %w", err)
	}
	return &Tracker{
		expenses:      e,
		sourceHandler: handler,
	}, nil
}

func (t *Tracker) Add(e Expense) {
	t.expenses = append(t.expenses, e)
}
func (t *Tracker) AddAll(e []Expense) {
	t.expenses = append(t.expenses, e...)
}

func (t *Tracker) List() []Expense {
	return t.expenses
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
			return nil
		}
	}
	return fmt.Errorf("in Tracker.Delete: expense with id=%s: %w", id, ErrExpenseNotFound)
}

func (t *Tracker) Summarize() Summary {
	count := len(t.expenses)
	var total float64
	for _, e := range t.expenses {
		total += e.Amount
	}

	return Summary{
		Count:   count,
		Total:   total,
		Average: float64(total) / float64(count),
	}
}

func (t *Tracker) Save() error {
	return t.sourceHandler.Save(t.expenses)
}

// func (t *Tracker) getByID(id uuid.UUID) (_ Expense, err error) {
// 	defer func() {
// 		if err != nil {
// 			err = fmt.Errorf("in Tracker.getByID: expense with id=%s: %w", id, err)
// 		}
// 	}()
//
// 	// first look for e in currSessExpenses and return if found
// 	for _, cse := range t.expenses {
// 		if cse.id == id {
// 			return cse, nil
// 		}
// 	}
//
// }

type Summary struct {
	Count   int
	Total   float64
	Average float64
}
