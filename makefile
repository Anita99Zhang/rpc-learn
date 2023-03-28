generate:
	protoc -I=./proto/ --go_out=./proto/ ./proto/helloworld.proto