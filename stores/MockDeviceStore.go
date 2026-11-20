package stores

type MockDeviceStore struct {
	devices []Device
}

func NewMockDeviceStore() *MockDeviceStore {
	return &MockDeviceStore{}
}

func (s *MockDeviceStore) GetDevices() []Device {
	// returns a list of devices and their services
	return s.devices
}

func (s *MockDeviceStore) AddDevice(name string, ip string) error {
	new := Device{Name: name, IP: ip, Services: []string{}}
	s.devices = append(s.devices, new)
	return nil
}
