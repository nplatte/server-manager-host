package stores

type DeviceStore interface {
	GetDevices() []Device
	AddDevice(string, string) error
}

type Device struct {
	Name     string
	IP       string
	Services []string
}

type ServiceStore interface {
	AddService(string) error
	GetServices() []Service
}

type Service struct {
	Name string
}
