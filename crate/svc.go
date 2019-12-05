package crate

import (
	"github.com/tchaudhry91/bottle/bottle"
)

type CrateService interface {
	StoreBottle(*bottle.Bottle) error
	PourBottle(id string) (*bottle.Bottle, error)
	DrainBottle(id string) (*bottle.Bottle, error)
}

type crateService struct {
	store BottleStore
}

func NewCrateService(store BottleStore) *crateService {
	return &crateService{
		store: store,
	}
}

// StoreBottle puts a bottle into the crate
func (svc *crateService) StoreBottle(b *bottle.Bottle) error {
	return svc.store.Put(b)
}

// PourBottle fetches the bottle while preserving a copy in the crate
func (svc *crateService) PourBottle(id string) (b *bottle.Bottle, err error) {
	return svc.store.Get(id)
}

// DrainBottle fetches the bottle from the crate, deleting it in the process
func (svc *crateService) DrainBottle(id string) (b *bottle.Bottle, err error) {
	b, err = svc.store.Get(id)
	if err != nil {
		return
	}
	err = svc.store.Delete(id)
	return b, err
}
