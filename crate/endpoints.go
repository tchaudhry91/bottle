package crate

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/tchaudhry91/bottle/bottle"
)

type getRequest struct {
	ID string `json:"id,omitempty"`
}

type getResponse struct {
	Bottle *bottle.Bottle `json:"bottle,omitempty"`
	Err    string         `json:"err,omitempty"`
}

func makePourEndpoint(svc CrateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getRequest)
		b, err := svc.PourBottle(req.ID)
		if err != nil {
			return getResponse{
				Bottle: b,
				Err:    err.Error(),
			}, nil
		}
		return getResponse{
			Bottle: b,
		}, nil
	}
}

func makeDrainEndpoint(svc CrateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getRequest)
		b, err := svc.DrainBottle(req.ID)
		if err != nil {
			return getResponse{
				Bottle: b,
				Err:    err.Error(),
			}, nil
		}
		return getResponse{
			Bottle: b,
		}, nil
	}
}

type storeRequest struct {
	Bottle *bottle.Bottle `json:"bottle,omitempty"`
}

type storeResponse struct {
	Err string `json:"err,omitempty"`
}

func makeStoreEndpoint(svc CrateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(storeRequest)
		err = svc.StoreBottle((req.Bottle))
		if err != nil {
			return storeResponse{
				Err: err.Error(),
			}, nil
		}
		return storeResponse{}, nil
	}
}
