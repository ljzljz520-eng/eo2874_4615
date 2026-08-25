package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
	"traininganalysis/internal/model"
)

var buckets = [][]byte{[]byte("athletes"), []byte("sessions"), []byte("drills"), []byte("plans")}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return &Store{db: db}, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put[T any](s *Store, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](s *Store, b []byte, key string, out *T) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveAthlete(v model.Athlete) error { return put(s, buckets[0], v.ID, v) }
func (s *Store) LoadAthlete(id string) (model.Athlete, error) {
	var v model.Athlete
	e := get(s, buckets[0], id, &v)
	return v, e
}
func (s *Store) SaveSession(v model.TrainingSession) error { return put(s, buckets[1], v.ID, v) }
func (s *Store) SaveDrill(v model.DrillResult) error       { return put(s, buckets[2], v.ID, v) }
func (s *Store) SavePlan(v model.TeamPlan) error           { return put(s, buckets[3], v.ID, v) }
