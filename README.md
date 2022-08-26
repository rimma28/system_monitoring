# system-monitoring

A small monitoring daemon that sends info about server's health via protobuf

## Supported OS

Currently supports only Linux (Ubuntu)

## Requirements

* one of the supported OS
* docker
* internet access (for pulling docker images)

## Running

To run a service, simply type
`make run`

To run a client, after server startup type
`docker exec -it system-monitor ./client/client`

## Developing

To regenerate GRPC implementation, type
`protoc --go_out=. --go-grpc_out=. protobuf/system-monitor.proto`