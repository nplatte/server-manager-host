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
