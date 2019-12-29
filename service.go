package main

import (
	"context"

	pb "github.com/tarciosaraiva/consignment-service/proto/consignment"
	vp "github.com/tarciosaraiva/vessel-service/proto/vessel"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo         repository
	vesselClient vp.VesselService
}

// CreateConsignmen - add consignment to store using repo
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Find vessel
	vesselResp, err := s.vesselClient.FindAvailable(context.Background(), &vp.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	req.VesselId = vesselResp.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments - read all consignments from store via repo
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}
