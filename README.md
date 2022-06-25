# XCASH API Tools - Go
Building tools for xcash using Go

The tool covers:  
Global API (includes [DPOPS](https://docs.xcash.foundation/api/dpops), [Namespace](https://docs.xcash.foundation/api/namespace) and [Xpayment Twitter](https://docs.xcash.foundation/api/xpayment-twitter))  
Blockchain ([xcash daemon](https://docs.xcash.foundation/applications/rpc-calls/json-rpc-methods) and [xcash wallet](https://docs.xcash.foundation/applications/rpc-calls/xcash-wallet-rpc))

Most methods are inclduing


# How to Install

Run this inside your project directory 
```shell
go get github.com/X-CASH-official/XCASH-api-tools-go
```

Add the import
```shell
github.com/X-CASH-official/XCASH-api-tools-go
```

# How to setup

Default Configuration (will use remote daemon and remote wallet)
```shell
const XCASH_DAEMON_ADDRESS = "";
const XCASH_WALLET_ADDRESS = "";
```

For example if you want to use a local xcash daemon and xcash wallet  
Default Configuration (will use remote daemon and remote wallet)
```shell
const XCASH_DAEMON_ADDRESS = "127.0.0.1:18281";
const XCASH_WALLET_ADDRESS = "127.0.0.1:18285";
```

# Testing
Run the test

```bash
go test -v
```

# Example

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
}
```
