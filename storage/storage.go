package storage

import "coginfra/types"

type Storage interface {
	// Get retrieves the value associated with the given key.
	GetDocuments(path string) (types.Documents, error)
	// Set stores the given value and associates it with the given key.
	UpdateDocuments(path, doc string) error
	// Delete removes the value associated with the given key.
	DeleteDocuments(key string) error
	// ValidateApiKey checks if the given API key is valid.
	ValidateApiKey(key string) (bool, error)

	CreateAPIKey() (string, error)

	Disconnect() error
}
