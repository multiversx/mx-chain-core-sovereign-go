module github.com/multiversx/mx-chain-core-go

go 1.23.0

require (
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gammazero/deque v0.2.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.4
	github.com/hashicorp/go-set v0.1.13
	github.com/mr-tron/base58 v1.2.0
	github.com/multiversx/mx-chain-vm-common-go v1.6.0
	github.com/multiversx/mx-chain-vm-go v1.5.43
	github.com/nikolaydubina/fpdecimal v0.16.0
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	golang.org/x/crypto v0.33.0
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/multiversx/mx-chain-logger-go v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gogo/protobuf => github.com/multiversx/protobuf v1.3.2
