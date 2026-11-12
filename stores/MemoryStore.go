package stores

type MemoryDeviceStore struct {
	devices []Device
}

func NewMemoryStore() *MemoryDeviceStore {
	return &MemoryDeviceStore{}
}

func (s *MemoryDeviceStore) GetDevices() []Device {
	// returns a list of devices and their services
	return s.devices
}

func (s *MemoryDeviceStore) AddDevice(name string, ip string) error {
	new := Device{Name: name, IP: ip, Services: []string{}}
	s.devices = append(s.devices, new)
	return nil
}
