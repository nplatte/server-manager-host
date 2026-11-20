package stores

type MockServiceStore struct {
	services []Service
}

func NewMockServiceStore() *MockServiceStore {
	return &MockServiceStore{}
}

func (s *MockServiceStore) AddService(name string) error {
	s.services = append(s.services, Service{name})
	return nil
}

func (s *MockServiceStore) GetServices() []Service {
	return s.services
}
