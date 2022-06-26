# XCASH API Tools - Go
Building tools for xcash using Go

The tool covers:  
Global API (includes [DPOPS](https://docs.xcash.foundation/api/dpops), [Namespace](https://docs.xcash.foundation/api/namespace) and [Xpayment Twitter](https://docs.xcash.foundation/api/xpayment-twitter))  
Blockchain ([xcash daemon](https://docs.xcash.foundation/applications/rpc-calls/json-rpc-methods) and [xcash wallet](https://docs.xcash.foundation/applications/rpc-calls/xcash-wallet-rpc))

Most methods are included


# How to Install

Run this inside your project directory 
```shell
go get github.com/X-CASH-official/XCASH-api-tools-go
```

Add the import
```go
github.com/X-CASH-official/XCASH-api-tools-go
```

# How to setup

Default Configuration (will use remote daemon and remote wallet)
```go
const XCASH_DAEMON_ADDRESS = "";
const XCASH_WALLET_ADDRESS = "";
```

For example if you want to use a local xcash daemon and xcash wallet  
Default Configuration (will use remote daemon and remote wallet)
```go
const XCASH_DAEMON_ADDRESS = "127.0.0.1:18281";
const XCASH_WALLET_ADDRESS = "127.0.0.1:18285";
```

Initialize the settings  
```go
xcash.Initialize(XCASH_DAEMON_ADDRESS,XCASH_WALLET_ADDRESS);
```

# Testing
Run the test

```bash
go test -v
```

# Example

The full request structure of each method can be found in [here](https://github.com/X-CASH-official/XCASH-api-tools-go/blob/main/XCASH-api-tools-go.go) but they generally follow the request structure of the API

The full response structure of each request can be found in [here](https://github.com/X-CASH-official/XCASH-api-tools-go/blob/main/XCASH-api-tools-go-structures.go)

Global API calls will just return the structure or nil  
All other calls return the structure or nil and an error structure or nil.  

The example below calls one method from the global API, and one method from the Blockchain

```go
package main

import (
"fmt"
"os"
"github.com/X-CASH-official/XCASH-api-tools-go"
)

func main() {
    // initialize
    const XCASH_DAEMON_ADDRESS = "127.0.0.1:18281";
    const XCASH_WALLET_ADDRESS = "127.0.0.1:18285";
    
    xcash.Initialize(XCASH_DAEMON_ADDRESS,XCASH_WALLET_ADDRESS);
    
    blockData,xcashError := xcash.Blockchain_get_block_count()
	if blockData == nil && xcashError == nil {                
		fmt.Println("An error has occurred")
		os.Exit(1)
	}
        if blockData == nil && xcashError != nil {                
		fmt.Println(xcashError.Error.Message)
		os.Exit(1)
	}
    
    fmt.Println(blockData.Result.Count) // prints 993163
    
    s,error := json.Marshal(blockData)
    if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
    
	fmt.Println(string(s))
    
    /* Prints out the following:
    {  
      "id": "0",  
      "jsonrpc": "2.0",  
      "result": {  
        "count": 993163,  
        "status": "OK"
        "untrusted": "false"  
      }  
    }  
    */
    
    
    
    
    APIStats := xcash.API_Blockchain_stats()
	if APIStats == nil {                
		fmt.Println("An error has occurred")
		os.Exit(1)
	}
    
    fmt.Println(APIStats.Height) // prints 810000
    
    s,error = json.Marshal(APIStats)
    if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
    
	fmt.Println(string(s))
    
    /* Prints out the following:
{
  "height": 810000,
  "hash":"c7aa6eb38c47e7f013a5f8042477d1734ff9808fdc8608fb088085d624d2d509",
  "reward": 20000000,
  "size": 20000,
  "version": 13,
  "versionBlockHeight": 1000000,
  "nextVersionBlockHeight": 0,
  "totalPublicTx": 100000,
  "totalPrivateTx": 100000,
  "circulatingSupply": 10000000,
  "generatedSupply": 100000000,
  "totalSupply": 100000000,
  "emissionReward": 1000000000,
  "emissionHeight": 1000000000,
  "emissionTime": 1000000000,
  "inflationHeight": 1000000000,
  "inflationTime": 1000000000
}
    */
}
```
