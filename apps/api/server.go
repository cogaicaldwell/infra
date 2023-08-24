package api

import (
	"coginfra/protos"
	"coginfra/storage"
)

// import main "coginfra
type Server struct {
	listenAddress string
	storage       storage.Storage
	HasStarted    bool
	protos.UnimplementedDocumentServiceServer
}

func NewServer(listenAddress string, store storage.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		storage:       store,
	}
}
