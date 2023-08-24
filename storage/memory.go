package storage

type MemoryStorage struct{}

func (s *MemoryStorage) Get(key string) (string, error) {
	return "", nil
}

func (s *MemoryStorage) Set(key, value string) error {
	return nil
}

func (s *MemoryStorage) Delete(key string) error {
	return nil
}

func (s *MemoryStorage) ValidateApiKey(key string) (bool, error) {
	return false, nil
}

func (s *MemoryStorage) CreateAPIKey() (string, error) {
	return "asdfbshdbfbsdbf", nil
}

func (s *MemoryStorage) Disconnect() error {
	return nil
}
