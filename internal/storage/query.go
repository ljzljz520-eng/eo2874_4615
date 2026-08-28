package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"traininganalysis/internal/model"
)

func (s *Store) ListAthletes() ([]model.Athlete, error) {
	out := []model.Athlete{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[0]).ForEach(func(_, v []byte) error {
			var a model.Athlete
			if e := json.Unmarshal(v, &a); e != nil {
				return e
			}
			out = append(out, a)
			return nil
		})
	})
	return out, e
}
func (s *Store) DeleteAthlete(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
