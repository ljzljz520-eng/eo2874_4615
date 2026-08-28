package main

import (
	"fmt"
	"traininganalysis/internal/analytics"
	"traininganalysis/internal/model"
	"traininganalysis/internal/storage"
	"traininganalysis/internal/training"
)

func main() {
	db, err := storage.Open("training.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	roster := training.NewRoster()
	roster.Add(model.Athlete{ID: "a1", Name: "Lin", AgeGroup: "U16", Position: "Guard"})
	roster.Add(model.Athlete{ID: "a2", Name: "Kai", AgeGroup: "U18", Position: "Center"})
	roster.Filter("U16", 1, 20)
	s := analytics.Summarize(roster.Current())
	fmt.Println(s)
}
