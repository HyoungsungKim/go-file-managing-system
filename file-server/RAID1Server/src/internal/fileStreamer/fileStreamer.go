package fileStreamer

import (
	"bytes"
	"io"

	v1 "fileServer.com/RAID1Server/src/api/v1"
	"fileServer.com/RAID1Server/src/internal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ackStatusOK         = 200
	ackStatusBadRequest = 404
)

type Server struct {
	storage storage.Manager
	v1.UnimplementedFileStreamServiceServer
}

var file *storage.File

func resetFile(f *storage.File) {
	f = &storage.File{
		Name:   "",
		Path:   "",
		Buffer: &bytes.Buffer{},
	}
}

func NewServer(storage storage.Manager) Server {
	return Server{
		storage: storage,
	}
}

func (s Server) SendFileInfo(infoStream v1.FileStreamService_SendFileInfoServer) error {
	fileInfo, err := infoStream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	} else {
		fileName := fileInfo.GetFileName()
		filePath := fileInfo.GetFilePath()

		file = storage.NewFile(fileName, filePath)

		return infoStream.SendAndClose(&v1.Ack{
			AckStatusCode:    ackStatusOK,
			AckStatusMessage: "File info is received",
		})
	}

}

func (s Server) StreamFile(fileStream v1.FileStreamService_StreamFileServer) error {
	for {
		req, err := fileStream.Recv()
		if err == io.EOF {
			if err := s.storage.Store(file); err != nil {
				resetFile(file)
				return status.Error(codes.Internal, err.Error())
			}

			resetFile(file)
			return fileStream.SendAndClose(&v1.Ack{
				AckStatusCode:    ackStatusOK,
				AckStatusMessage: "File stream is succeeded",
			})
		}
		if err != nil {
			resetFile(file)
			return status.Error(codes.Internal, err.Error())
		}
		if err := file.Write(req.GetChunkData()); err != nil {
			resetFile(file)
			return status.Error(codes.Internal, err.Error())
		}
	}

}

func (s Server) DeleteFile(deleteFileSignal v1.FileStreamService_DeleteFileServer) error {
	signal, err := deleteFileSignal.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if signal.GetDeleteFileSignal() {
		if err := s.storage.Delete(file); err != nil {
			resetFile(file)
			return status.Error(codes.Internal, err.Error())
		}
	}

	resetFile(file)
	return deleteFileSignal.SendAndClose(&v1.Ack{
		AckStatusCode:    ackStatusOK,
		AckStatusMessage: "File delete is succeeded",
	})
}
