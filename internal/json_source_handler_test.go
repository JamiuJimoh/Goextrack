package expense

import (
	"os"
	"testing"

	"uuid"
)

const testDataPath = "testdata/expenses.json"

func TestJSONSourceHandler_Save(t *testing.T) {
	expenses := []Expense{
		newTestExpense(uuid.New(), 1000),
		newTestExpense(uuid.New(), 3000),
	}

	handler := &JSONSourceHandler{
		dataPath: testDataPath,
	}

	err := handler.Save(expenses)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	data, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Fatalf("failed to read saved file: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("expected saved file to contain data")
	}
}

func TestJSONSourceHandler_All(t *testing.T) {
	handler := &JSONSourceHandler{
		dataPath: testDataPath,
	}

	expenses, err := handler.All()
	if err != nil {
		t.Fatalf("All failed: %v", err)
	}

	t.Logf("Loaded expenses: %+v", expenses)
}
