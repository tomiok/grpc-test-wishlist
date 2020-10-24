package main

import (
	"context"
	"fmt"
	pb "github.com/tomiok/grpc-test-wishlist/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)


func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		panic("cannot initialize grpc client: " + err.Error())
	}

	client := pb.NewWishListServiceClient(conn)

	res, err := client.Create(context.Background(), &pb.CreateWishListReq{
		WishList:             &pb.WishList{
			Id:                   "1",
			Name:                 "my wish list",
		},
	})

	if err != nil {
		panic("cannot call server: " + err.Error())
	}

	fmt.Println(res.WishListId)
}
