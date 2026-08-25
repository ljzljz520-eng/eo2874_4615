package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"traininganalysis/internal/model"
)

func (s *Store) SaveBatch(athletes []model.Athlete, sessions []model.TrainingSession) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, a := range athletes {
			data, e := json.Marshal(a)
			if e != nil {
				return e
			}
			if e = tx.Bucket(buckets[0]).Put([]byte(a.ID), data); e != nil {
				return e
			}
		}
		for _, v := range sessions {
			data, e := json.Marshal(v)
			if e != nil {
				return e
			}
			if e = tx.Bucket(buckets[1]).Put([]byte(v.ID), data); e != nil {
				return e
			}
		}
		return nil
	})
}
