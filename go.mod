module github.com/multiversx/mx-chain-core-go

go 1.23.0

require (
	github.com/btcsuite/btcd/btcutil v1.1.6
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.4
	github.com/mr-tron/base58 v1.2.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.38.0
	google.golang.org/grpc v1.72.1
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250519155744-55703ea1f237 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gogo/protobuf => github.com/multiversx/protobuf v1.3.2
