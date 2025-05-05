//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf --gogoslick_out=paths=source_relative:. sourceChainID.proto

package dto

// ValidChains defines supported cross chain ids
var ValidChains = map[ChainID]struct{}{
	MVX: {},
	ETH: {},
}

// IsCrossChainID returns true if the provided chain id is supported
func IsCrossChainID(id ChainID) bool {
	_, ok := ValidChains[id]
	return ok
}
