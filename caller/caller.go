package caller

/*

The caller is responsible for taking the requests from the handlers that forming them into gRPC requests.
It then communicates them back to the calling methods and makes sure issues are hadled appropiately.

*/

type Caller interface {
	AddDevice(string) error
}
