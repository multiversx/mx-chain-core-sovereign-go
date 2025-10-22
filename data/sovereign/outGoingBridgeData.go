// Run this in terminal

//go:generate protoc -I=. --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go --plugin=protoc-gen-go-grpc=$(go env GOPATH)/bin/protoc-gen-go-grpc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative outGoingBridgeData.proto
package sovereign
