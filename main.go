package main

import (
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/tarciosaraiva/consignment-service/proto/consignment"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("service.consignment"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
