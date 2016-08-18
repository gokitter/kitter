# 1. Install Go mobile go get golang.org/x/mobile/cmd/gomobile
# 2. gomobile init

get:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

protocols:
	protoc -I/usr/local/include -I. \
 -I${GOPATH}/src \
 --go_out=plugins=grpc:. \
 kitter/kitter.proto

server: protocols
	go run ./server/server.go

mac_client: protocols
	go build -o mac_client ./cli/cli.go

ios_client: protocols
	gomobile bind -v -target=ios -o ./apps/ios/ErrorKitteh/frameworks/mobilesdk.framework ./client

android_client: protocols
	gomobile bind -v -target=android -o ../android/mobilesdk/mobilesdk.aar ./client