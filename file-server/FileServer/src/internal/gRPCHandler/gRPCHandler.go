package gRPCHandler

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	v1 "fileServer.com/FileServer/src/api/v1"
	"google.golang.org/grpc"
)

/*
	acknowledge code of gRPC response. These codes follow http stats code.
*/
const (
	ackStatusOK         = 200
	ackStatusBadRequest = 404
)

type Client struct {
	client v1.FileStreamServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: v1.NewFileStreamServiceClient(conn),
	}
}

/*
	SendFileInfo sends dirID and fileName to server
	Assume URI of stored file is ./storage/ZGlyMS9kaXIyL2RpcjM=/helloworld.txt
	dirID is `ZGlyMS9kaXIyL2RpcjM=`, fileName is `ZGlyMS9kaXIyL2RpcjM=`
*/
func (c Client) SendFileInfo(ctx context.Context, dirId string, fileName string) (*v1.Ack, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(30*time.Second))
	defer cancel()

	// Send file info before sending file stream
	infoStream, err := c.client.SendFileInfo(ctx)
	if err == nil {
		if infoStreamErr := infoStream.Send(&v1.FileInfo{
			FileName: fileName,
			FilePath: dirId + "/",
		}); infoStreamErr != nil {
			log.Println(infoStreamErr)

			return &v1.Ack{
				AckStatusCode:    ackStatusBadRequest,
				AckStatusMessage: "Fail to send file info",
			}, infoStreamErr
		}
	} else {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "Fail to excute gRPC `SendFileInfo`.",
		}, err
	}

	res, err := infoStream.CloseAndRecv()
	return res, err
}

/*
	StreamFIle sends file from file server to RAID1 server
*/
func (c Client) StreamFile(ctx context.Context, root string, dirId string, fileName string) (*v1.Ack, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(30*time.Second))
	defer cancel()

	fileStream, err := c.client.StreamFile(ctx)
	if err != nil {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "gRPC error",
		}, err
	}

	log.Println("File path is", dirId+"/"+fileName)
	fil, err := os.Open(root + "/" + dirId + "/" + fileName)
	if err != nil {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "File does not exists",
		}, err
	}

	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)

			return &v1.Ack{
				AckStatusCode:    ackStatusBadRequest,
				AckStatusMessage: "File read error",
			}, err
		}

		if err := fileStream.Send(&v1.FileStreamRequest{
			ChunkData: buf[:num],
		}); err != nil {
			log.Println(err)

			return &v1.Ack{
				AckStatusCode:    ackStatusBadRequest,
				AckStatusMessage: "File to excute gRPC `FileStreamRequest`.",
			}, err
		}
	}

	res, err := fileStream.CloseAndRecv()
	return res, err
}

func (c Client) DeleteFile(ctx context.Context, root string, dirId string, fileName string) (*v1.Ack, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(30*time.Second))
	defer cancel()

	deleteFileSignal, err := c.client.DeleteFile(ctx)
	if err != nil {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "gRPC error",
		}, err
	}

	log.Println("File path is", dirId+"/"+fileName)
	_, err = os.Open(root + "/" + dirId + "/" + fileName)
	if err != nil {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "File does not exists",
		}, err
	}

	if err := deleteFileSignal.Send(&v1.DeleteFileSignal{
		DeleteFileSignal: true,
	}); err != nil {
		log.Println(err)

		return &v1.Ack{
			AckStatusCode:    ackStatusBadRequest,
			AckStatusMessage: "Fail to excute gRCP `DeleteFileSignal`",
		}, err
	}

	res, err := deleteFileSignal.CloseAndRecv()

	return res, err
}
