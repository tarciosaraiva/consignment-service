package main

import (
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/tarciosaraiva/consignment-service/proto/consignment"
	vp "github.com/tarciosaraiva/vessel-service/proto/vessel"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("service.consignment"),
	)

	srv.Init()

	vc := vp.NewVesselService("vessel.service", srv.Client())

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vc})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
