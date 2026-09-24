package main

import (
	"log"

	"github.com/JamiuJimoh/Goextrack"
	"github.com/JamiuJimoh/Goextrack/internal"
)

func main() {
	handler := expense.NewJSONSourceHandler("data/expense.json")
	tracker, err := expense.NewTracker(handler)
	if err != nil {
		log.Fatal(err)
	}

	app := repl.New(tracker)
	log.Fatal(app.Run())
}
