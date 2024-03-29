package crate

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/tchaudhry91/bottle/bottle"
	"time"
)

const BottleBucket = "bottles"

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
		buck := tx.Bucket([]byte(BottleBucket))
		contents := buck.Get([]byte(id))
		if contents == nil {
			return errors.New("Bottle ID Not Found")
		}
		b.Contents = contents
		return nil
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

func (s *BoltBottleStore) Put(b *bottle.Bottle) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(BottleBucket))
		if err != nil {
			return err
		}
		return buck.Put([]byte(b.ID), b.Contents)
	})
}

func (s *BoltBottleStore) Delete(id string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(BottleBucket))
		return buck.Delete([]byte(id))
	})
}

func (s *BoltBottleStore) List() ([]*bottle.Bottle, error) {
	bb := []*bottle.Bottle{}
	err := s.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(BottleBucket))
		if buck == nil {
			return errors.New("No records found")
		}
		err := buck.ForEach(func(k []byte, v []byte) error {
			b := bottle.NewBottle(string(k))
			b.Contents = v
			bb = append(bb, b)
			return nil
		})
		return err
	})
	return bb, err
}
