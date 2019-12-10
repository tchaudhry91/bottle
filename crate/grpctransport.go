package crate

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/tchaudhry91/bottle/bottle"
	"github.com/tchaudhry91/bottle/crate/pb"
)

type grpcServer struct {
	store kitgrpc.Handler
	pour  kitgrpc.Handler
	drain kitgrpc.Handler
	list  kitgrpc.Handler
}

// NewGRPCServer makes a of endpoints available as a gRPC Crate Server
func NewGRPCServer(endpoints Endpoints, logger log.Logger) pb.CrateServer {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	return &grpcServer{
		pour: kitgrpc.NewServer(
			endpoints.Pour,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		drain: kitgrpc.NewServer(
			endpoints.Drain,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		store: kitgrpc.NewServer(
			endpoints.Store,
			decodeGRPCStoreRequest,
			encodeGRPCStoreResponse,
			options...,
		),
		list: kitgrpc.NewServer(
			endpoints.List,
			decodeGRPCEmptyRequest,
			encodeGRPCListResponse,
			options...,
		),
	}
}

func (s *grpcServer) Pour(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := s.pour.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}

func (s *grpcServer) Drain(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := s.drain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}

func (s *grpcServer) Store(ctx context.Context, req *pb.StoreRequest) (*pb.StoreResponse, error) {
	_, resp, err := s.store.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.StoreResponse), nil
}

func (s *grpcServer) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListResponse), nil
}

func decodeGRPCStoreRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.StoreRequest)
	return storeRequest{Bottle: PBToBottle(req.Bottle)}, nil
}

func encodeGRPCStoreResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(storeResponse)
	return &pb.StoreResponse{
		Err: resp.Err,
	}, nil
}

func decodeGRPCGetRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRequest)
	return getRequest{ID: req.Id}, nil
}

func encodeGRPCGetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(getResponse)
	return &pb.GetResponse{
		Bottle: BottleToPB(resp.Bottle),
		Err:    resp.Err,
	}, nil
}

func decodeGRPCEmptyRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return emptyRequest{}, nil
}

func encodeGRPCListResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listResponse)
	return &pb.ListResponse{
		Bottles: resp.Bottles,
		Err:     resp.Err,
	}, nil
}

func BottleToPB(bottle *bottle.Bottle) *pb.Bottle {
	return &pb.Bottle{
		Id:       bottle.ID,
		Contents: bottle.Contents,
	}
}

func PBToBottle(b *pb.Bottle) *bottle.Bottle {
	return &bottle.Bottle{
		ID:       b.Id,
		Contents: b.Contents,
	}
}
