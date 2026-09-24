package expense

import (
	"errors"
	"testing"
	"time"

	"uuid"
)

type fakeExpenseSource struct {
	expenses []Expense
}

func (f *fakeExpenseSource) All() ([]Expense, error) {
	return f.expenses, nil
}

func (f *fakeExpenseSource) Save([]Expense) error {
	return nil
}

func newTestExpense(id uuid.UUID, amount float64) Expense {
	now := time.Now()

	return Expense{
		ID:          id,
		Category:    Category(food),
		Amount:      amount,
		Description: "Test expense",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func TestNewTracker_Success(t *testing.T) {
	id := uuid.New()

	source := &fakeExpenseSource{
		expenses: []Expense{
			newTestExpense(id, 1500),
		},
	}

	tracker, err := NewTracker(source)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tracker == nil {
		t.Fatal("expected tracker, got nil")
	}

	got := tracker.List()

	if len(got) != 1 {
		t.Fatalf("expected 1 expense, got %d", len(got))
	}

	if got[0].ID != id {
		t.Errorf("expected ID %v, got %v", id, got[0].ID)
	}
}

func TestTracker_Add(t *testing.T) {
	tracker, err := NewTracker(&fakeExpenseSource{})

	if err != nil {
		t.Fatal(err)
	}

	expense := newTestExpense(uuid.New(), 2500)

	tracker.Add(expense)

	got := tracker.List()

	if len(got) != 1 {
		t.Fatalf("expected 1 expense, got %d", len(got))
	}

	if got[0].ID != expense.ID {
		t.Errorf("expected ID %v, got %v", expense.ID, got[0].ID)
	}
}

func TestTracker_List(t *testing.T) {
	expenses := []Expense{
		newTestExpense(uuid.New(), 1000),
		newTestExpense(uuid.New(), 2000),
	}

	tracker, err := NewTracker(&fakeExpenseSource{
		expenses: expenses,
	})

	if err != nil {
		t.Fatal(err)
	}

	got := tracker.List()

	if len(got) != len(expenses) {
		t.Fatalf(
			"expected %d expenses, got %d",
			len(expenses),
			len(got),
		)
	}
}

func TestTracker_Get_Found(t *testing.T) {
	id := uuid.New()
	expected := newTestExpense(id, 3000)

	tracker, err := NewTracker(&fakeExpenseSource{
		expenses: []Expense{expected},
	})

	if err != nil {
		t.Fatal(err)
	}

	got, err := tracker.Get(id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != expected.ID {
		t.Errorf("expected ID %v, got %v", expected.ID, got.ID)
	}

	if got.Amount != expected.Amount {
		t.Errorf("expected amount %.2f, got %.2f", expected.Amount, got.Amount)
	}
}

func TestTracker_Get_NotFound(t *testing.T) {
	tracker, err := NewTracker(&fakeExpenseSource{})

	if err != nil {
		t.Fatal(err)
	}

	_, err = tracker.Get(uuid.New())

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrExpenseNotFound) {
		t.Errorf("expected ErrExpenseNotFound, got %v", err)
	}
}

func TestTracker_Edit_Success(t *testing.T) {
	id := uuid.New()

	original := newTestExpense(id, 1000)
	updated := newTestExpense(id, 5000)

	tracker, err := NewTracker(&fakeExpenseSource{
		expenses: []Expense{original},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = tracker.Edit(updated)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, err := tracker.Get(id)

	if err != nil {
		t.Fatal(err)
	}

	if got.Amount != updated.Amount {
		t.Errorf(
			"expected amount %.2f, got %.2f",
			updated.Amount,
			got.Amount,
		)
	}
}

func TestTracker_Edit_NotFound(t *testing.T) {
	tracker, err := NewTracker(&fakeExpenseSource{})

	if err != nil {
		t.Fatal(err)
	}

	err = tracker.Edit(newTestExpense(uuid.New(), 1000))

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrExpenseNotFound) {
		t.Errorf("expected ErrExpenseNotFound, got %v", err)
	}
}

func TestTracker_Delete_Success(t *testing.T) {
	idToDelete := uuid.New()

	expenses := []Expense{
		newTestExpense(uuid.New(), 1000),
		newTestExpense(idToDelete, 2000),
		newTestExpense(uuid.New(), 3000),
	}

	tracker, err := NewTracker(&fakeExpenseSource{
		expenses: expenses,
	})

	if err != nil {
		t.Fatal(err)
	}

	err = tracker.Delete(expenses[1].ID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got := tracker.List()

	if len(got) != 2 {
		t.Fatalf("expected 2 expenses, got %d", len(got))
	}

	_, err = tracker.Get(idToDelete)

	if !errors.Is(err, ErrExpenseNotFound) {
		t.Errorf("expected deleted expense to be absent, got %v", err)
	}
}

func TestTracker_Delete_NotFound(t *testing.T) {
	tracker, err := NewTracker(&fakeExpenseSource{})

	if err != nil {
		t.Fatal(err)
	}

	err = tracker.Delete(uuid.New())

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrExpenseNotFound) {
		t.Errorf("expected ErrExpenseNotFound, got %v", err)
	}
}

func TestTracker_Summarize(t *testing.T) {
	expenses := []Expense{
		newTestExpense(uuid.New(), 1000),
		newTestExpense(uuid.New(), 2000),
		newTestExpense(uuid.New(), 3000),
	}

	tracker, err := NewTracker(&fakeExpenseSource{
		expenses: expenses,
	})

	if err != nil {
		t.Fatal(err)
	}

	summary := tracker.Summarize()

	if summary.Count != 3 {
		t.Errorf("expected count 3, got %d", summary.Count)
	}

	if summary.Total != 6000 {
		t.Errorf("expected total 6000, got %.2f", summary.Total)
	}

	if summary.Average != 2000.0 {
		t.Errorf("expected average 2000, got %.2f", summary.Average)
	}
}
