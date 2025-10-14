package testscommon

import (
	"github.com/multiversx/mx-chain-core-go/data/sovereign"
)

// HeaderSubscriberMock -
type HeaderSubscriberMock struct {
	AddHeaderCalled func(headerHash []byte, header sovereign.IncomingHeaderHandler) error
}

// AddHeader -
func (mock *HeaderSubscriberMock) AddHeader(headerHash []byte, header sovereign.IncomingHeaderHandler) error {
	if mock.AddHeaderCalled != nil {
		return mock.AddHeaderCalled(headerHash, header)
	}

	return nil
}

// IsInterfaceNil -
func (mock *HeaderSubscriberMock) IsInterfaceNil() bool {
	return mock == nil
}
