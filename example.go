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
        "status": "OK",
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
