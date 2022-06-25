package xcash

import (
"fmt"
"strings"
"bytes"
"strconv"
"encoding/json"
"io/ioutil"
"net/http"
"time"
"errors"
)

var XCASH_DAEMON_ADDRESS string = "us1.xcash.foundation:18281";
var XCASH_WALLET_ADDRESS string = "us1.xcash.foundation:18289";

func send_http_data(url string,data string) (string, error) {
  // create the http request
  req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
  if err != nil {
    return "",errors.New("Could not send the request")
  }

  // set the request headers  
  req.Header.Set("Content-Type", "application/json")
 
  // set the http client settings
  client := &http.Client{}
  client.Timeout = time.Second * 5

  // send the request
  resp, err := client.Do(req)
  if err != nil {
    return "",errors.New("")
  }

  // close the connection
  defer resp.Body.Close()
  
  // get the response body
  body, _ := ioutil.ReadAll(resp.Body)
  
  fmt.Println(string(body))

  return string(body),nil
}

func send_http_data_get(url string) (string, error) {
  resp, err := http.Get(url)
    if err != nil {
    return "",errors.New("Could not send the request")
  }

    // close the connection
  defer resp.Body.Close()
  
  // get the response body
  body, _ := ioutil.ReadAll(resp.Body)
  
  fmt.Println(string(body))

  return string(body),nil
}

func Initialize(s1 string,s2 string) {
  XCASH_DAEMON_ADDRESS = s1
  XCASH_WALLET_ADDRESS = s2
  return
}

func Blockchain_get_block_count() (*BlockchainGetBlockCount,*BlockchainError) {
    var d1 BlockchainGetBlockCount
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block_count"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block_hash(height int) (*BlockchainGetBlockHash,*BlockchainError) {
    var d1 BlockchainGetBlockHash
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"on_get_block_hash","params":[` + strconv.Itoa(height) + `]}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block_template(address string, reserve_bytes int) (*BlockchainGetBlockTemplate,*BlockchainError) {
    var d1 BlockchainGetBlockTemplate
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block_template","params":{"wallet_address":"` + address + `","reserve_size":` + strconv.Itoa(reserve_bytes) + `}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_last_block_header() (*BlockchainGetLastBlockHeader,*BlockchainError) {
    var d1 BlockchainGetLastBlockHeader
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_last_block_header"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block_header_by_hash(height string) (*BlockchainGetBlockHeaderByHash,*BlockchainError) {
    var d1 BlockchainGetBlockHeaderByHash
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block_header_by_hash","params":{"hash":"` + height + `"}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block_header_by_height(height int) (*BlockchainGetBlockHeaderByHeight,*BlockchainError) {
    var d1 BlockchainGetBlockHeaderByHeight
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block_header_by_height","params":{"height":` + strconv.Itoa(height) + `}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block_headers_range(startHeight int,endHeight int) (*BlockchainGetBlockHeadersRange,*BlockchainError) {
    var d1 BlockchainGetBlockHeadersRange
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block_headers_range","params":{"start_height":` + strconv.Itoa(startHeight) + `,"end_height":` + strconv.Itoa(endHeight) + `}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_block(height int) (*BlockchainGetBlock,*BlockchainError) {
    var d1 BlockchainGetBlock
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_block","params":{"height":` + strconv.Itoa(height) + `}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_info() (*BlockchainGetInfo,*BlockchainError) {
    var d1 BlockchainGetInfo
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_info"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_hard_fork_info() (*BlockchainHardForkInfo,*BlockchainError) {
    var d1 BlockchainHardForkInfo
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"hard_fork_info"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_version() (*BlockchainGetVersion,*BlockchainError) {
    var d1 BlockchainGetVersion
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_version"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_tx(tx string) (*BlockchainGetTx,*BlockchainError) {
    var d1 BlockchainGetTx
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_DAEMON_ADDRESS + "/get_transactions",`{"txs_hashes":["` + tx + `"],"decode_as_json":true}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func DPOPS_vote(delegate string, amount int64) (*DPOPSVote,*BlockchainError) {
    var d1 DPOPSVote
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"vote","params":{"delegate_data":"` + delegate + `","amount":"` + strconv.FormatInt(amount,10) + `"}}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func DPOPS_revote(amount int64) (*DPOPSRevote,*BlockchainError) {
    var d1 DPOPSRevote
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"vote","params":{"amount":"` + strconv.FormatInt(amount,10) + `"}}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func DPOPS_vote_status() (*DPOPSVoteStatus,*BlockchainError) {
    var d1 DPOPSVoteStatus
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"vote_status"}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func DPOPS_delegate_register(delegate string, IP string) (*DPOPSDelegateRegister,*BlockchainError) {
    var d1 DPOPSDelegateRegister
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"delegate_register","params":{"delegate_name":"` + delegate + `","delegate_IP_address":"` + IP +`"}}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func DPOPS_delegate_update(item string, value string) (*DPOPSDelegateUpdate,*BlockchainError) {
    var d1 DPOPSDelegateUpdate
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"delegate_update","params":{"item":"` + item + `","value":"` + value + `"}}`)
    if !strings.Contains(data_send, "\"status\": \"OK\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_height() (*BlockchainGetHeight,*BlockchainError) {
    var d1 BlockchainGetHeight
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_height"}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_transfer_split(address []string, amount []int64) (*BlockchainTransferSplit,*BlockchainError) {
    var d1 BlockchainTransferSplit
    var e1 BlockchainError
    
    data := `{"jsonrpc":"2.0","id":"0","method":"transfer_split","params":{"destinations":[`
    for i, _ := range address {
        data += `{"amount":` + strconv.FormatInt(amount[i],10) + `,"address":"` + address[i] + `"},`
    }
    data = data[:len(data)-1]
    data += `],"ring_size":21,"get_tx_keys": true}}`
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",data)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_sweep_all(address string) (*BlockchainSweepAll,*BlockchainError) {
    var d1 BlockchainSweepAll
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"sweep_all","params":{"address":"` + address + `","ring_size":21}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_get_payments(payment_id string) (*BlockchainGetPayments,*BlockchainError) {
    var d1 BlockchainGetPayments
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_payments","params":{"payment_id":"` + payment_id + `"}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}

func Blockchain_incoming_transfers(transfer_type string) (*BlockchainIncomingTransfers,*BlockchainError) {
    var d1 BlockchainIncomingTransfers
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"incoming_transfers","params":{"transfer_type":"` + transfer_type + `"}}`)
    if !strings.Contains(data_send, "\"result\"") || error != nil {
      if err := json.Unmarshal([]byte(data_send), &e1); err != nil {
        return nil,nil
    }
    return nil,&e1
  }
  if err := json.Unmarshal([]byte(data_send), &d1); err != nil {
    return nil,nil
  }
  return &d1,&e1
}
