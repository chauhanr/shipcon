package main

import (
	pb "github.com/chauhanr/shipcon/user-service/proto/user"
	"log"
	"github.com/micro/go-micro"
)

func main(){
	// create a db connection and close it at the end
	db, err := CreateConnection()
	defer db.Close()
	if err != nil{
		log.Fatalf("Could not connect to the database %v", err)
	}
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("lastest"),
	)
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
}