build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/chauhanr/shipcon/user-service proto/user/user.proto
	docker build -t user-service .