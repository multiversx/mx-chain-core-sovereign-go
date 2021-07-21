module github.com/ElrondNetwork/elrond-go-core

go 1.13

require (
	github.com/ElrondNetwork/elrond-vm-common v1.1.0
	github.com/btcsuite/btcutil v1.0.2
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/mr-tron/base58 v1.2.0
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil v0.0.0-20190901111213-e4ec7b275ada // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
)

replace github.com/gogo/protobuf => github.com/ElrondNetwork/protobuf v1.3.2
