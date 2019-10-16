#   Casperlabs-ee-grpc-go-util  [![Build Status](https://travis-ci.org/hdac-io/casperlabs-ee-grpc-go-util.svg?branch=master)](https://travis-ci.org/hdac-io/casperlabs-ee-grpc-go-util)


This module is for using GRPC and UTIL in GO to utilize Casperlabs Execution-Engine.


## Install
- For use go module, add to environment variable.
```bash
$ export GO111MODULE=on
```
- Download to module in $GOPATH
```bash
$ go get -u github.com/hdac-io/casperlabs-ee-grpc-go-util
```

## Example
- Running to Casperlabs Execution-Engine 
```bash
$ git clone https://github.com/CasperLabs/CasperLabs.git

$ cd execution-engine

$ make build

$ ./target/debug/casperlabs-engine-grpc-server ~/.casperlabs/.casper-node.sock
```
- Install "casperlabs-ee-grpc-go-util" module

- Import "casperlabs-ee-grpc-go-util"
```go
import (
    "github.com/hdac-io/casperlabs-ee-grpc-go-util/util"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/grpc"
)

func main() {
    client := grpc.Connect(`/.casperlabs/.casper-node.sock`)

	mintTokenCode := util.LoadWasmFile("./example/contracts/mint_token.wasm")

	mintResult := grpc.Validate(client, mintCode)
}
```

## Running Example code
- Revise to unix socket path in "./example/hello.go"
- Run
```bash
$ make example
```

## Testing
```bash
$ make test
```

## Develop
- Install protobuf
- Generate proto file
```bash
$ make proto
```
