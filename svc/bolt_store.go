package svc

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/tchaudhry91/bottle/bottle"
	"time"
)

const BOTTLE_BUCKET = "bottles"

// BoltBottleStore is a BoltDB backed bottle store
type BoltBottleStore struct {
	db *bolt.DB
}

// NewBoltBottleStore returns a boltDB backed bottle store
func NewBoltBottleStore(file string) (s *BoltBottleStore, err error) {
	s = &BoltBottleStore{}
	s.db, err = bolt.Open(file, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s *BoltBottleStore) Get(id string) (*bottle.Bottle, error) {
	b := bottle.NewBottle(id)

	err := s.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(BOTTLE_BUCKET))
		contents := buck.Get([]byte(id))
		if contents != nil {
			return errors.New("Bottle ID Not Found")
		}
		return nil
	})
	if err != nil {
		return b, err
	}
	return b, nil
}
