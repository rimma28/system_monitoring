# system-monitoring

A monitoring daemon that sends info about server's health via protobuf

## Supported OS

Currently supports only Linux (Ubuntu)

## Functionality & Features

* Calculating:
  * Load average (Linux only)
  * CPU load (Linux only)
  * Disks load (Linux only)
* Simple client for daemon
* Unit tests

## Running

To run a service, simply type
`make run`

To run a client, after server startup type
`docker exec -it system-monitor ./client/client`

## Developing

To regenerate GRPC implementation, type
`protoc --go_out=. --go-grpc_out=. protobuf/system-monitor.proto`