package caller

import "errors"

/*

The caller is responsible for taking the requests from the handlers that forming them into gRPC requests.
It then communicates them back to the calling methods and makes sure issues are hadled appropiately.

*/

var (
	ErrUnreacable        = errors.New("Client Device Unreachable")
	ErrRefused           = errors.New("Client Refused Connection")
	ErrServiceNotFound   = errors.New("Service Not Found")
	ErrServiceRestricted = errors.New("Service Restricted")
)

type Caller interface {
	AddDevice(string) (string, error)
	AddService(string) (string, error)
}
