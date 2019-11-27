package svc

import (
	"github.com/tchaudhry91/bottle/bottle"
)

// Store is an interface for a bottle database
type BottleStore interface {
	Get(id string) (b *bottle.Bottle, err error)
	Put(b *bottle.Bottle) error
}
