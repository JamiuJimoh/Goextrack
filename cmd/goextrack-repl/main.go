package main

import (
	"log"
	"os"
	"path/filepath"

	repl "github.com/JamiuJimoh/Goextrack"
	expense "github.com/JamiuJimoh/Goextrack/internal"
)

func main() {
	dataPath, err := expenseFilePath()
	if err != nil {
		log.Fatal(err)
	}
	handler := expense.NewJSONSourceHandler(dataPath)
	tracker, err := expense.NewTracker(handler)
	if err != nil {
		log.Fatal(err)
	}

	app := repl.New(tracker)
	log.Fatal(app.Run())
}

func expenseFilePath() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(
		baseDir,
		"goextrack",
		"data",
		"expense.json",
	), nil
}
