package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type JSONSourceHandler struct {
	dataPath string
}

func NewJSONSourceHandler(p string) JSONSourceHandler {
	return JSONSourceHandler{dataPath: p}
}

func (sh JSONSourceHandler) All() (_ []Expense, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in JSONSourceHandler.All: %w", err)
		}
	}()

	data, err := os.ReadFile(sh.dataPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Expense{}, nil
		}
		return nil, err
	}
	var expenses []Expense
	error := json.Unmarshal(data, &expenses)
	return expenses, error
}

func (sh JSONSourceHandler) Save(e []Expense) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in JSONSourceHandler.Save: %w", err)
		}
	}()
	data, err := json.MarshalIndent(&e, "", " ")
	if err != nil {
		return err
	}
	if os.MkdirAll(filepath.Dir(sh.dataPath), 0o755); err != nil {
		return err
	}

	return os.WriteFile(sh.dataPath, data, 0o644)
}
