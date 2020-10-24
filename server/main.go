package main

import (
	"context"
	"fmt"
	pb "github.com/tomiok/grpc-test-wishlist/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedWishListServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateWishListReq) (*pb.CreateWishListResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}

func (s *server) Add(context.Context, *pb.AddItemReq) (*pb.AddItemResp, error) {
	return nil, nil
}

func (s *server) List(context.Context, *pb.ListWishListReq) (*pb.ListWishListResp, error) {
	return nil, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterWishListServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}
