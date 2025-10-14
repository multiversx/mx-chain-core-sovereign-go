package sovereign

import (
	"errors"
)

var errNilHeaderSubscriber = errors.New("nil header subscriber provided")

var errNilIncomingHeader = errors.New("nil incoming header provided")
