package storage

import "context"

type MockStorage struct{}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m *MockStorage) Upload(
	ctx context.Context,
	key string,
	data []byte,
) error {
	return nil
}
