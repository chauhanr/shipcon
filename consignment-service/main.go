package main


import (

	pb "shipcon/consignment-service/proto/consignment"
	vesselProto "shipcon/vessel-service/proto/vessel"
		"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"fmt"
	"log"
)


type IRepository interface{
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConsignmentRepository struct{
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment{
	return repo.consignments
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error){
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

type service struct{
	repo IRepository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error{
	//consignment, err := s.repo.Create(req)
	vesselRes, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	log.Printf("Found vessels %s\n", vesselRes.Vessel.Name)
	if err != nil{
		return err
	}

	req.VesselId = vesselRes.Vessel.Id
	consignment, err := s.repo.Create(req)
	if err != nil{
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response)  error{
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main(){
	repo := &ConsignmentRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(),&service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}

