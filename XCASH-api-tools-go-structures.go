package xcash

type BlockchainError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type BlockchainGetBlockCount struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Count  int    `json:"count"`
		Status string `json:"status"`
	} `json:"result"`
}

type BlockchainGetBlockHash struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type BlockchainGetBlockTemplate struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlockhashingBlob  string `json:"blockhashing_blob"`
		BlocktemplateBlob string `json:"blocktemplate_blob"`
		Difficulty        int    `json:"difficulty"`
		ExpectedReward    int64  `json:"expected_reward"`
		Height            int    `json:"height"`
		PrevHash          string `json:"prev_hash"`
		ReservedOffset    int    `json:"reserved_offset"`
		Status            string `json:"status"`
		Untrusted         bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetLastBlockHeader struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlockHeader struct {
			BlockSize            int    `json:"block_size"`
			BlockWeight          int    `json:"block_weight"`
			CumulativeDifficulty int64  `json:"cumulative_difficulty"`
			Depth                int    `json:"depth"`
			Difficulty           int    `json:"difficulty"`
			Hash                 string `json:"hash"`
			Height               int    `json:"height"`
			MajorVersion         int    `json:"major_version"`
			MinorVersion         int    `json:"minor_version"`
			Nonce                int    `json:"nonce"`
			NumTxes              int    `json:"num_txes"`
			OrphanStatus         bool   `json:"orphan_status"`
			PowHash              string `json:"pow_hash"`
			PrevHash             string `json:"prev_hash"`
			Reward               int64  `json:"reward"`
			Timestamp            int    `json:"timestamp"`
		} `json:"block_header"`
		Status    string `json:"status"`
		Untrusted bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetBlockHeaderByHash struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlockHeader struct {
			BlockSize            int    `json:"block_size"`
			BlockWeight          int    `json:"block_weight"`
			CumulativeDifficulty int64  `json:"cumulative_difficulty"`
			Depth                int    `json:"depth"`
			Difficulty           int    `json:"difficulty"`
			Hash                 string `json:"hash"`
			Height               int    `json:"height"`
			MajorVersion         int    `json:"major_version"`
			MinorVersion         int    `json:"minor_version"`
			Nonce                int    `json:"nonce"`
			NumTxes              int    `json:"num_txes"`
			OrphanStatus         bool   `json:"orphan_status"`
			PowHash              string `json:"pow_hash"`
			PrevHash             string `json:"prev_hash"`
			Reward               int64  `json:"reward"`
			Timestamp            int    `json:"timestamp"`
		} `json:"block_header"`
		Status    string `json:"status"`
		Untrusted bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetBlockHeaderByHeight struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlockHeader struct {
			BlockSize            int    `json:"block_size"`
			BlockWeight          int    `json:"block_weight"`
			CumulativeDifficulty int64  `json:"cumulative_difficulty"`
			Depth                int    `json:"depth"`
			Difficulty           int    `json:"difficulty"`
			Hash                 string `json:"hash"`
			Height               int    `json:"height"`
			MajorVersion         int    `json:"major_version"`
			MinorVersion         int    `json:"minor_version"`
			Nonce                int    `json:"nonce"`
			NumTxes              int    `json:"num_txes"`
			OrphanStatus         bool   `json:"orphan_status"`
			PowHash              string `json:"pow_hash"`
			PrevHash             string `json:"prev_hash"`
			Reward               int64  `json:"reward"`
			Timestamp            int    `json:"timestamp"`
		} `json:"block_header"`
		Status    string `json:"status"`
		Untrusted bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetBlockHeadersRange struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Headers []struct {
			BlockSize            int    `json:"block_size"`
			BlockWeight          int    `json:"block_weight"`
			CumulativeDifficulty int64  `json:"cumulative_difficulty"`
			Depth                int    `json:"depth"`
			Difficulty           int    `json:"difficulty"`
			Hash                 string `json:"hash"`
			Height               int    `json:"height"`
			MajorVersion         int    `json:"major_version"`
			MinorVersion         int    `json:"minor_version"`
			Nonce                int    `json:"nonce"`
			NumTxes              int    `json:"num_txes"`
			OrphanStatus         bool   `json:"orphan_status"`
			PowHash              string `json:"pow_hash"`
			PrevHash             string `json:"prev_hash"`
			Reward               int64  `json:"reward"`
			Timestamp            int    `json:"timestamp"`
		} `json:"headers"`
		Status    string `json:"status"`
		Untrusted bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetBlock struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Blob        string `json:"blob"`
		BlockHeader struct {
			BlockSize            int    `json:"block_size"`
			BlockWeight          int    `json:"block_weight"`
			CumulativeDifficulty int64  `json:"cumulative_difficulty"`
			Depth                int    `json:"depth"`
			Difficulty           int    `json:"difficulty"`
			Hash                 string `json:"hash"`
			Height               int    `json:"height"`
			MajorVersion         int    `json:"major_version"`
			MinorVersion         int    `json:"minor_version"`
			Nonce                int    `json:"nonce"`
			NumTxes              int    `json:"num_txes"`
			OrphanStatus         bool   `json:"orphan_status"`
			PowHash              string `json:"pow_hash"`
			PrevHash             string `json:"prev_hash"`
			Reward               int64  `json:"reward"`
			Timestamp            int    `json:"timestamp"`
		} `json:"block_header"`
		JSON        string `json:"json"`
		MinerTxHash string `json:"miner_tx_hash"`
		Status      string `json:"status"`
		Untrusted   bool   `json:"untrusted"`
	} `json:"result"`
}

type BlockchainGetInfo struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AltBlocksCount           int    `json:"alt_blocks_count"`
		BlockSizeLimit           int    `json:"block_size_limit"`
		BlockSizeMedian          int    `json:"block_size_median"`
		BlockWeightLimit         int    `json:"block_weight_limit"`
		BlockWeightMedian        int    `json:"block_weight_median"`
		BootstrapDaemonAddress   string `json:"bootstrap_daemon_address"`
		CumulativeDifficulty     int64  `json:"cumulative_difficulty"`
		DatabaseSize             int    `json:"database_size"`
		Difficulty               int    `json:"difficulty"`
		FreeSpace                int64  `json:"free_space"`
		GreyPeerlistSize         int    `json:"grey_peerlist_size"`
		Height                   int    `json:"height"`
		HeightWithoutBootstrap   int    `json:"height_without_bootstrap"`
		IncomingConnectionsCount int    `json:"incoming_connections_count"`
		Mainnet                  bool   `json:"mainnet"`
		Nettype                  string `json:"nettype"`
		Offline                  bool   `json:"offline"`
		OutgoingConnectionsCount int    `json:"outgoing_connections_count"`
		RPCConnectionsCount      int    `json:"rpc_connections_count"`
		Stagenet                 bool   `json:"stagenet"`
		StartTime                int    `json:"start_time"`
		Status                   string `json:"status"`
		Target                   int    `json:"target"`
		TargetHeight             int    `json:"target_height"`
		Testnet                  bool   `json:"testnet"`
		TopBlockHash             string `json:"top_block_hash"`
		TxCount                  int    `json:"tx_count"`
		TxPoolSize               int    `json:"tx_pool_size"`
		Untrusted                bool   `json:"untrusted"`
		UpdateAvailable          bool   `json:"update_available"`
		WasBootstrapEverUsed     bool   `json:"was_bootstrap_ever_used"`
		WhitePeerlistSize        int    `json:"white_peerlist_size"`
	} `json:"result"`
}

type BlockchainHardForkInfo struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		EarliestHeight int    `json:"earliest_height"`
		Enabled        bool   `json:"enabled"`
		State          int    `json:"state"`
		Status         string `json:"status"`
		Threshold      int    `json:"threshold"`
		Untrusted      bool   `json:"untrusted"`
		Version        int    `json:"version"`
		Votes          int    `json:"votes"`
		Voting         int    `json:"voting"`
		Window         int    `json:"window"`
	} `json:"result"`
}

type BlockchainGetVersion struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status    string `json:"status"`
		Untrusted bool   `json:"untrusted"`
		Version   int    `json:"version"`
	} `json:"result"`
}

type BlockchainGetTx struct {
	Status string `json:"status"`
	Txs    []struct {
		AsHex           string `json:"as_hex"`
		AsJSON          string `json:"as_json"`
		BlockHeight     int    `json:"block_height"`
		BlockTimestamp  int    `json:"block_timestamp"`
		DoubleSpendSeen bool   `json:"double_spend_seen"`
		InPool          bool   `json:"in_pool"`
		OutputIndices   []int  `json:"output_indices"`
		TxHash          string `json:"tx_hash"`
	} `json:"txs"`
	TxsAsHex  []string `json:"txs_as_hex"`
	TxsAsJSON []string `json:"txs_as_json"`
	Untrusted bool     `json:"untrusted"`
}

type DPOPSVote struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		VoteStatus string `json:"vote_status"`
	} `json:"result"`
}

type DPOPSRevote struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type DPOPSVoteStatus struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type DPOPSDelegateRegister struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		DelegateRegisterStatus string `json:"delegate_register_status"`
	} `json:"result"`
}

type DPOPSDelegateUpdate struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		DelegateUpdateStatus string `json:"delegate_update_status"`
	} `json:"result"`
}

type BlockchainGetBalance struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Balance              int64 `json:"balance"`
		MultisigImportNeeded bool  `json:"multisig_import_needed"`
		PerSubaddress        []struct {
			Address           string `json:"address"`
			AddressIndex      int    `json:"address_index"`
			Balance           int64  `json:"balance"`
			Label             string `json:"label"`
			NumUnspentOutputs int    `json:"num_unspent_outputs"`
			UnlockedBalance   int64  `json:"unlocked_balance"`
		} `json:"per_subaddress"`
		UnlockedBalance int64 `json:"unlocked_balance"`
	} `json:"result"`
}

type BlockchainGetAddress struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Address   string `json:"address"`
		Addresses []struct {
			Address      string `json:"address"`
			AddressIndex int    `json:"address_index"`
			Label        string `json:"label"`
			Used         bool   `json:"used"`
		} `json:"addresses"`
	} `json:"result"`
}

type BlockchainGetAddressIndex struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Index struct {
			Major int `json:"major"`
			Minor int `json:"minor"`
		} `json:"index"`
	} `json:"result"`
}

type BlockchainCreateAddress struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Address      string `json:"address"`
		AddressIndex int    `json:"address_index"`
	} `json:"result"`
}

type BlockchainGetAccounts struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		SubaddressAccounts []struct {
			AccountIndex    int    `json:"account_index"`
			Balance         int64  `json:"balance"`
			BaseAddress     string `json:"base_address"`
			Label           string `json:"label"`
			Tag             string `json:"tag"`
			UnlockedBalance int64  `json:"unlocked_balance"`
		} `json:"subaddress_accounts"`
		TotalBalance         int64 `json:"total_balance"`
		TotalUnlockedBalance int64 `json:"total_unlocked_balance"`
	} `json:"result"`
}

type BlockchainCreateAccount struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AccountIndex int    `json:"account_index"`
		Address      string `json:"address"`
	} `json:"result"`
}

type BlockchainGetHeight struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AmountList    []int64  `json:"amount_list"`
		FeeList       []int    `json:"fee_list"`
		MultisigTxset string   `json:"multisig_txset"`
		TxHashList    []string `json:"tx_hash_list"`
		TxKeyList     []string `json:"tx_key_list"`
		UnsignedTxset string   `json:"unsigned_txset"`
	} `json:"result"`
}

type BlockchainTransferSplit struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AmountList    []int64  `json:"amount_list"`
		FeeList       []int    `json:"fee_list"`
		MultisigTxset string   `json:"multisig_txset"`
		TxHashList    []string `json:"tx_hash_list"`
		TxKeyList     []string `json:"tx_key_list"`
		UnsignedTxset string   `json:"unsigned_txset"`
	} `json:"result"`
}

type BlockchainSweepAll struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AmountList    []int64  `json:"amount_list"`
		FeeList       []int    `json:"fee_list"`
		MultisigTxset string   `json:"multisig_txset"`
		TxHashList    []string `json:"tx_hash_list"`
		TxKeyList     []string `json:"tx_key_list"`
		UnsignedTxset string   `json:"unsigned_txset"`
	} `json:"result"`
}

type BlockchainGetPayments struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Payments []struct {
			Address      string `json:"address"`
			Amount       int64  `json:"amount"`
			BlockHeight  int    `json:"block_height"`
			PaymentID    string `json:"payment_id"`
			SubaddrIndex struct {
				Major int `json:"major"`
				Minor int `json:"minor"`
			} `json:"subaddr_index"`
		} `json:"payments"`
	} `json:"result"`
}

type BlockchainIncomingTransfers struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Transfers []struct {
			Amount       int64  `json:"amount"`
			GlobalIndex  int    `json:"global_index"`
			KeyImage     string `json:"key_image"`
			Spent        bool   `json:"spent"`
			SubaddrIndex struct {
				Major int `json:"major"`
				Minor int `json:"minor"`
			} `json:"subaddr_index"`
			TxHash string `json:"tx_hash"`
		} `json:"transfers"`
	} `json:"result"`
}
