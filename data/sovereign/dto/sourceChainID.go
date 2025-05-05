//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=paths=source_relative:. sourceChainID.proto

package dto

// ValidChains defines supported cross chain ids
var ValidChains = map[ChainID]struct{}{
	MVX: {},
	ETH: {},
}

// ValidChainName defines supported cross chain id names
var ValidChainName = map[string]struct{}{
	MVX.String(): {},
	ETH.String(): {},
}

// IsValidCrossChainID returns true if the provided chain id is supported
func IsValidCrossChainID(id ChainID) bool {
	_, ok := ValidChains[id]
	return ok
}

func IsValidCrossChainIDString(id string) bool {
	_, ok := ValidChainName[id]
	return ok
}
