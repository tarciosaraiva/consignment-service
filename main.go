package main

import (
	"fmt"
	"strings"

	"github.com/micro/go-micro"
	pb "github.com/tarciosaraiva/consignment-service/proto/consignment"
)

const (
	host = "localhost"
	port = "50051"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Address(strings.Join([]string{host, port}, ":")),
		micro.Name("service.consignment"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
