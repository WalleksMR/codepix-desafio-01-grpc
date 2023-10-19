package main

import (
	"github.com/walleksmr/codepix-desafio-01-grpc/application/grpc"
	db "github.com/walleksmr/codepix-desafio-01-grpc/infrastructure/database"
)

func main() {
	database := db.ConnectDB()
	port := 50051
	grpc.StartGrpcServer(database, port)
}
