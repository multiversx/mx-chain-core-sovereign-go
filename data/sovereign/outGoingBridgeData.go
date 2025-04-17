//go:generate protoc -I=. --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src --go_out=. --go_out=paths=source_relative:. outGoingBridgeData.proto
package sovereign
