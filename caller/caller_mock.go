package caller

func NewMockCaller() *MockCaller {
	return &MockCaller{}
}

type MockCaller struct {
	returnString string
	returnError  error
}

func (c *MockCaller) AddDevice(IP string) (string, error) {
	return c.returnString, c.returnError
}

func (c *MockCaller) SetReturnString(newValue string) {
	c.returnString = newValue
}

func (c *MockCaller) SetReturnError(newValue error) {
	c.returnError = newValue
}
