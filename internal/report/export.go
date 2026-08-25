package report

import (
	"encoding/csv"
	"os"
	"traininganalysis/internal/model"
)

func ExportCSV(path string, athletes []model.Athlete) error {
	f, e := os.Create(path)
	if e != nil {
		return e
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	if e = w.Write([]string{"id", "name", "age_group", "position"}); e != nil {
		return e
	}
	for _, a := range athletes {
		if e = w.Write([]string{a.ID, a.Name, a.AgeGroup, a.Position}); e != nil {
			return e
		}
	}
	return w.Error()
}
func ImportCSV(path string) ([]model.Athlete, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	rows, e := csv.NewReader(f).ReadAll()
	if e != nil {
		return nil, e
	}
	out := []model.Athlete{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) >= 4 {
			out = append(out, model.Athlete{ID: row[0], Name: row[1], AgeGroup: row[2], Position: row[3], Active: true})
		}
	}
	return out, nil
}
