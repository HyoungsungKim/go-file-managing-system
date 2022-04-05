package main

import (
	"log"
	"net"

	"fileServer.com/RAID1Server/src/internal/fileStreamer"
	"fileServer.com/RAID1Server/src/internal/storage"
	"google.golang.org/grpc"

	v1 "fileServer.com/RAID1Server/src/api/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	defer lis.Close()

	upSrv := fileStreamer.NewServer(storage.New("./storage/"))
	rpcSrv := grpc.NewServer()

	v1.RegisterFileStreamServiceServer(rpcSrv, upSrv)
	log.Fatal(rpcSrv.Serve(lis))
}
