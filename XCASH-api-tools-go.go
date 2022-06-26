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
var XCASH_GLOBAL_API_ADDRESS string = "https://api.xcash.foundation"

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

func DPOPS_delegate_register(delegate string, IP string, key string) (*DPOPSDelegateRegister,*BlockchainError) {
    var d1 DPOPSDelegateRegister
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"delegate_register","params":{"delegate_name":"` + delegate + `","delegate_IP_address":"` + IP + `","delegates_public_key":"` + key +`"}}`)
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

func Blockchain_query_key(key_type string) (*BlockchainQueryKey,*BlockchainError) {
    var d1 BlockchainQueryKey
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"query_key","params":{"key_type":"` + key_type + `"}}`)
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

func Blockchain_make_integrated_address(address string, payment_id string) (*BlockchainQueryKey,*BlockchainError) {
    var d1 BlockchainQueryKey
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"make_integrated_address","params":{"standard_address":"` + address + `","payment_id":"` + payment_id + `"}}`)
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

func Blockchain_split_integrated_address(integrated_address string) (*BlockchainSplitIntegratedAddress,*BlockchainError) {
    var d1 BlockchainSplitIntegratedAddress
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"split_integrated_address","params":{"integrated_address":"` + integrated_address + `"}}`)
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

func Blockchain_rescan_blockchain() (*BlcokchainRescanBlockchain,*BlockchainError) {
    var d1 BlcokchainRescanBlockchain
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"rescan_blockchain"}`)
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

func Blockchain_get_tx_key(tx string) (*BlockchainGetTxKey,*BlockchainError) {
    var d1 BlockchainGetTxKey
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_tx_key","params":{"txid":"` + tx + `"}}`)
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

func Blockchain_check_tx_key(tx string, key string, address string) (*BlockchainCheckTxKey,*BlockchainError) {
    var d1 BlockchainCheckTxKey
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"check_tx_key","params":{"txid":"` + tx + `","tx_key":"` + key + `","address":"` + address + `"}}`)
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

func Blockchain_get_tx_proof(tx string, address string, message string) (*BlockchainGetTxProof,*BlockchainError) {
    var d1 BlockchainGetTxProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_tx_proof","params":{"txid":"` + tx + `","address":"` + address + `","message":"` + message + `"}}`)
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

func Blockchain_check_tx_proof(tx string, address string, message string, signature string) (*BlockchainCheckTxProof,*BlockchainError) {
    var d1 BlockchainCheckTxProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"check_tx_proof","params":{"txid":"` + tx + `","address":"` + address + `","message":"` + message + `","signature":"` + signature + `"}}`)
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

func Blockchain_get_spend_proof(tx string, message string) (*BlockchainGetSpendProof,*BlockchainError) {
    var d1 BlockchainGetSpendProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_spend_proof","params":{"txid":"` + tx + `","message":"` + message + `"}}`)
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

func Blockchain_check_spend_proof(tx string, message string, signature string) (*BlockchainCheckSpendProof,*BlockchainError) {
    var d1 BlockchainCheckSpendProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"check_spend_proof","params":{"txid":"` + tx + `","message":"` + message  +`","signature":"` + signature + `"}}`)
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

func Blockchain_get_reserve_proof(all bool, amount int64, message string) (*BlockchainGetReserveProof,*BlockchainError) {
    var d1 BlockchainGetReserveProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"get_reserve_proof","params":{"all":` + strconv.FormatBool(all) + `,"amount":` + strconv.FormatInt(amount,10) + `,"message":"` + message + `"}}`)
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

func Blockchain_check_reserve_proof(address string, message string, signature string) (*BlockchainCheckReserveProof,*BlockchainError) {
    var d1 BlockchainCheckReserveProof
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"check_reserve_proof","params":{"address":"` + address + `","signature":"` + signature + `","message":"` + message + `"}}`)
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

func Blockchain_sign(message string) (*BlockchainSign,*BlockchainError) {
    var d1 BlockchainSign
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"sign","params":{"data":"` + message + `"}}`)
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

func Blockchain_verify(message string, address string, signature string) (*BlockchainVerify,*BlockchainError) {
    var d1 BlockchainVerify
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"verify","params":{"data":"` + message + `","address":"` + address + `","signature":"` + signature + `"}}`)
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

func Blockchain_rescan_spent() (*BlockchainRescanSpent,*BlockchainError) {
    var d1 BlockchainRescanSpent
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"rescan_spent"}`)
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

func Blockchain_validate_address(address string) (*BlockchainValidateAddress,*BlockchainError) {
    var d1 BlockchainValidateAddress
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"validate_address","params":{"address":"` + address + `"}}`)
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

func Blockchain_create_wallet(filename string, password string) (*BlockchainCreateWallet,*BlockchainError) {
    var d1 BlockchainCreateWallet
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"create_wallet","params":{"filename":"` + filename + `","password":"` + password + `","language":"English"}}`)
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

func Blockchain_open_wallet(filename string, password string) (*BlockchainOpenWallet,*BlockchainError) {
    var d1 BlockchainOpenWallet
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"open_wallet","params":{"filename":"` + filename + `","password":"` + password + `"}}`)
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

func Blockchain_close_wallet() (*BlockchainCloseWallet,*BlockchainError) {
    var d1 BlockchainCloseWallet
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"close_wallet"}`)
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

func Blockchain_change_wallet_password(old_password string, new_password string) (*BlockchainChangeWalletPassword,*BlockchainError) {
    var d1 BlockchainChangeWalletPassword
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"change_wallet_password","params":{"old_password":"` + old_password + `","new_password":"` + new_password + `"}}`)
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

func Namespace_update_remote_data(item string, value string) (*NamespaceUpdateRemoteData,*BlockchainError) {
    var d1 NamespaceUpdateRemoteData
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"update_remote_data","params":{"item":"` + item + `","value":"` + value + `"}}`)
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

func Namespace_remote_data_save_name(name string) (*NamespaceRemoteDataSaveName,*BlockchainError) {
    var d1 NamespaceRemoteDataSaveName
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"remote_data_save_name","params":{"name":"` + name + `"}}`)
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

func Namespace_remote_data_purchase_name(saddress string, saddress_signature string,paddress string, paddress_signature string,tx_hash string) (*NamespaceRemoteDataPurchaseName,*BlockchainError) {
    var d1 NamespaceRemoteDataPurchaseName
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"remote_data_purchase_name","params":{"saddress":"` + saddress + `","saddress_signature":"` + saddress_signature + `","paddress":"` + paddress + `","paddress_signature":"` + paddress_signature + `","tx_hash":"` + tx_hash + `"}}`)
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

func Namespace_remote_data_delegate_set_amount(amount int64) (*NamespaceRemoteDataDelegateSetAmount,*BlockchainError) {
    var d1 NamespaceRemoteDataDelegateSetAmount
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"remote_data_delegate_set_amount","params":{"amount":"` + strconv.FormatInt(amount,10) + `"}}`)
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

func Namespace_remote_data_renewal_start() (*NamespaceRemoteDataRenewalStart,*BlockchainError) {
    var d1 NamespaceRemoteDataRenewalStart
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"remote_data_renewal_start"}`)
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

func Namespace_remote_data_renewal_end(tx string) (*NamespaceRemoteDataRenewalEnd,*BlockchainError) {
    var d1 NamespaceRemoteDataRenewalEnd
    var e1 BlockchainError
    
	data_send,error := send_http_data("http://" + XCASH_WALLET_ADDRESS + "/json_rpc",`{"jsonrpc":"2.0","id":"0","method":"remote_data_renewal_end","params":{"tx_hash":"` + tx + `"}}`)
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

func API_Blockchain_stats() (*APIBlockchainStats) {
    var d1 APIBlockchainStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/stats/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_block_data(height int) (*APIBlockchainBlockData) {
    var d1 APIBlockchainBlockData
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/blocks/" + strconv.Itoa(height) + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_tx_data(tx string) (*APIBlockchainTxData) {
    var d1 APIBlockchainTxData
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/tx/" + tx + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_prove_tx(tx string, address string, key string) (*APIBlockchainProveTx) {
    var d1 APIBlockchainProveTx
    
    data_send,error := send_http_data(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/tx/prove/",`{"tx":"` + tx + `","address":"` + address + `","key": "` + key + `"}`)
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_prove_balance(address string, signature string) (*APIBlockchainProveBalance) {
    var d1 APIBlockchainProveBalance
    
    data_send,error := send_http_data(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/address/prove/",`{"address":"` + address + `","signature": "` + signature + `"}`)
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_tx_history(tx_type string, address string) (*[]*APIBlockchainTxData) {
    d1:=[]*APIBlockchainTxData{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/address/history/" + tx_type + "/" + address + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_validate_address(address string) (*APIBlockchainValidateAddress) {
    var d1 APIBlockchainValidateAddress
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/address/validate/" + address + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Blockchain_create_integrated_address(address string, payment_id string) (*APIBlockchainCreateIntegratedAddress) {
    var d1 APIBlockchainCreateIntegratedAddress
    
    data_send,error := send_http_data(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/blockchain/unauthorized/address/createIntegrated/",`{"address":"` + address + `","paymentId": "` + payment_id + `"}`)
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_stats() (*APIDPOPSStats) {
    var d1 APIDPOPSStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xpayment-twitter/twitter/unauthorized/stats/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_registered_delegates() (*[]*APIDPOPSDelegates) {
    d1:=[]*APIDPOPSDelegates{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/registered/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_online_delegates() (*[]*APIDPOPSDelegates) {
    d1:=[]*APIDPOPSDelegates{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/online/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_active_delegates() (*[]*APIDPOPSDelegates) {
    d1:=[]*APIDPOPSDelegates{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/active/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_delegate(delegate string) (*APIDPOPSDelegate) {
    var d1 APIDPOPSDelegate
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/" + delegate + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_delegate_round_stats(delegate string) (*APIDPOPSDelegateRoundStats) {
    var d1 APIDPOPSDelegateRoundStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/rounds/" + delegate + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_delegates_votes(delegate string, start int, limit int) (*[]*APIDPOPSDelegatesVotes) {
    d1:=[]*APIDPOPSDelegatesVotes{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/delegates/votes/" + delegate + "/" + strconv.Itoa(start) + "/" + strconv.Itoa(limit))
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_vote_detail(address string) (*APIDPOPSVoteDetails) {
    var d1 APIDPOPSVoteDetails
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/votes/" + address + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_round_detail(height int) (*APIDPOPSRoundDetails) {
    var d1 APIDPOPSRoundDetails
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/rounds/" + strconv.Itoa(height) + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_DPOPS_last_block_producer() (*APIDPOPSLastBlockProducer) {
    var d1 APIDPOPSLastBlockProducer
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/dpops/unauthorized/lastBlockProducer/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_stats() (*APINamespaceStats) {
    var d1 APINamespaceStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/stats/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_registered_delegates() (*[]*APINamespaceRegisteredDelegates) {
    d1:=[]*APINamespaceRegisteredDelegates{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/delegates/registered/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_delegates(delegate string) (*APINamespaceDelegates) {
    var d1 APINamespaceDelegates
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/delegates/" + delegate + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_names(name string) (*APINamespaceName) {
    var d1 APINamespaceName
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/names/" + name + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_name_status(name string) (*APINamespaceNameStatus) {
    var d1 APINamespaceNameStatus
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/names/status/" + name + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_address_status(address string) (*APINamespaceAddressStatus) {
    var d1 APINamespaceAddressStatus
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/addresses/status/" + address + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_name_to_address(name string) (*APINamespaceNameToAddress) {
    var d1 APINamespaceNameToAddress
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/names/convert/" + name + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Namespace_address_to_name(address string) (*APINamespaceAddressToName) {
    var d1 APINamespaceAddressToName
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xcash/namespace/unauthorized/addresses/convert/" + address + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Xpayment_Twitter_stats() (*APIXpaymentTwitterStats) {
    var d1 APIXpaymentTwitterStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xpayment-twitter/twitter/unauthorized/stats/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Xpayment_Twitter_stats_per_day(start int, limit int) (*[]*APIXpaymentTwitterStatsPerDay) {
    d1:=[]*APIXpaymentTwitterStatsPerDay{}
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xpayment-twitter/twitter/unauthorized/statsPerDay/" + strconv.Itoa(start) + "/" + strconv.Itoa(limit) + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Xpayment_Twitter_top_stats(amount int) (*APIXpaymentTwitterTopStats) {
    var d1 APIXpaymentTwitterTopStats
    
    data_send,error := send_http_data_get(XCASH_GLOBAL_API_ADDRESS + "/v1/xpayment-twitter/twitter/unauthorized/topStats/" + strconv.Itoa(amount) + "/")
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}

func API_Xpayment_Twitter_recent_tips(amount int, sort string, sort_type string) (*[]*APIXpaymentTwitterRecentTips) {
    d1:=[]*APIXpaymentTwitterRecentTips{}
    
    data_send,error := send_http_data(XCASH_GLOBAL_API_ADDRESS + "/v1/xpayment-twitter/twitter/unauthorized/recentTips/" + strconv.Itoa(amount) + "/",`{"sort":"` + sort + `","type": "` + sort_type + `"}`)
    if strings.Contains(data_send, "\"Error\"") || error != nil {
      return nil
    }
    if error = json.Unmarshal([]byte(data_send), &d1); error != nil {
      return nil
    }
    return &d1
}
