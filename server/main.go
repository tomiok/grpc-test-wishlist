package main

import (
	"context"
	pb "github.com/tomiok/grpc-test-wishlist/proto"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedWishListServiceServer
}

func (s *server) Create(context.Context, *pb.CreateWishListReq) (*pb.CreateWishListResp, error) {

	return &pb.CreateWishListResp{
		WishListId: generateID(),
	}, nil
}

func (s *server) Add(context.Context, *pb.AddItemReq) (*pb.AddItemResp, error) {
	return nil, nil
}

func (s *server) List(context.Context, *pb.ListWishListReq) (*pb.ListWishListResp, error) {
	return nil, nil
}

func generateID() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Int())
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		panic("cannot initialize tcp connection" + err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWishListServiceServer(grpcServer, &server{})
	if err = grpcServer.Serve(lis); err != nil {
		panic("server did not start! " + err.Error())
	}
}
