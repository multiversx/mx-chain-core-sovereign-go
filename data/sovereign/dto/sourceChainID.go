//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=paths=source_relative:. sourceChainID.proto

package dto

// ValidChains defines supported cross chain ids
var ValidChains = map[ChainID]struct{}{
	MVX: {},
	ETH: {},
}

// ValidChainNames defines supported cross chain id names
var ValidChainNames = map[string]struct{}{
	MVX.String(): {},
	ETH.String(): {},
}

// IsValidCrossChainID returns true if the provided chain id is supported
func IsValidCrossChainID(id ChainID) bool {
	_, ok := ValidChains[id]
	return ok
}

// IsValidCrossChainIDString returns true if the provided chain id as string is supported
func IsValidCrossChainIDString(id string) bool {
	_, ok := ValidChainNames[id]
	return ok
}
