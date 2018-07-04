# first stage will build the application in a docker images
# We will use the golang docker image which contains all the correct build tools and libs.
# the `as builder` used below this give a container name that can be used later.
#FROM golang:1.9.0 as builder

# now we set the work dir or our build image with gpath
#WORKDIR /go/src/github.com/chauhanr/shipcon/consignment-cli

#COPY . .

# first will pull all the dependency that we need to use.
#RUN go get -u github.com/golang/dep/cmd/dep

# create a dep project and run ensure which will ensure all the dep are pulled.
#RUN dep init && dep ensure

# now build the binaries
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM debian:latest

#RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

CMD ["./consignment-cli"]