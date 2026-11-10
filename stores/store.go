package stores

type DeviceStore interface {
	GetDevices() []Device
	AddDevice(string, string) error
}

type Device struct {
	name     string
	ip       string
	services []string
}
