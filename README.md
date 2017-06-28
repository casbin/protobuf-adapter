Protobuf Adapter [![Godoc](https://godoc.org/github.com/casbin/protobuf-adapter?status.svg)](https://godoc.org/github.com/casbin/protobuf-adapter)
====

Protobuf Adapter is the [Google Protocol Buffers](https://developers.google.com/protocol-buffers/) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from Protocol Buffers or save policy to it.

## Installation

    go get github.com/casbin/protobuf-adapter

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/mysql-adapter"
)

func main() {
	// Initialize a Protobuf adapter and use it in a Casbin enforcer:
	b := []byte{} // b stores Casbin policy in Protocol Buffers.
	a := protobufadapter.NewProtobufAdapter(&b) // Use b as the data source. 
	e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	// Load the policy from Protocol Buffers bytes b.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to Protocol Buffers bytes b.
	e.SavePolicy()
}
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
