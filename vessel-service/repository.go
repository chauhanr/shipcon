package main

import (
	pb "github.com/chauhanr/shipcon/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

const(
	dbName = "shipcon"
	vesselCollection = "vessels"
)

type Repository interface{
	FindAvailable(*pb.Specification)(*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}


type VesselRepository struct{
	session *mgo.Session
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error){
	var vessel *pb.Vessel
	// querying the consignment service
	err := repo.collection().Find(bson.M{
		"capacity": bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil{
		return nil, err
	}
	return nil, errors.New("Vessel not found")
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) collection() *mgo.Collection{
	return repo.session.DB(dbName).C(vesselCollection)
}
