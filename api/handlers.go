package api

import (
	"context"
	"encoding/json"
	"errors"

	"coginfra/protos"
	"coginfra/utils"
)

func (s *Server) AddDocument(
	ctx context.Context,
	in *protos.AddDocumentRequest,
) (*protos.AddDocumentResponse, error) {
	isValidApiKey, err := utils.ValidateAPIKey(ctx, s.storage)
	if err != nil {
		return nil, err
	}

	if !isValidApiKey {
		return nil, errors.New("not a valid API Key")
	}

	err = s.storage.UpdateDocuments(in.Path, in.Doc)

	if err != nil {
		return nil, err
	}

	return &protos.AddDocumentResponse{
		Inserted: true,
	}, nil
}

func (s *Server) ReadDocument(
	ctx context.Context,
	in *protos.ReadDocumentRequest,
) (*protos.ReadDocumentResponse, error) {
	isValidApiKey, err := utils.ValidateAPIKey(ctx, s.storage)
	if err != nil {
		return nil, err
	}

	if !isValidApiKey {
		return nil, errors.New("not a valid API Key")
	}

	documents, err := s.storage.GetDocuments(in.Path)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(documents)
	if err != nil {
		return nil, err
	}

	return &protos.ReadDocumentResponse{
		Documents: string(jsonBytes),
	}, nil
}
