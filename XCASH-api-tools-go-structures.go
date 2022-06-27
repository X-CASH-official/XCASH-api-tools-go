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

type BlockchainQueryKey struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Key string `json:"key"`
	} `json:"result"`
}

type BlockchainMakeIntegratedAddress struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		IntegratedAddress string `json:"integrated_address"`
		PaymentID         string `json:"payment_id"`
	} `json:"result"`
}

type BlockchainSplitIntegratedAddress struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		IsSubaddress    bool   `json:"is_subaddress"`
		PaymentID       string `json:"payment_id"`
		StandardAddress string `json:"standard_address"`
	} `json:"result"`
}

type BlcokchainRescanBlockchain struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type BlockchainGetTxKey struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		TxKey string `json:"tx_key"`
	} `json:"result"`
}

type BlockchainCheckTxKey struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Confirmations int   `json:"confirmations"`
		InPool        bool  `json:"in_pool"`
		Received      int64 `json:"received"`
	} `json:"result"`
}

type BlockchainGetTxProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Signature string `json:"signature"`
	} `json:"result"`
}

type BlockchainCheckTxProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Confirmations int   `json:"confirmations"`
		Good          bool  `json:"good"`
		InPool        bool  `json:"in_pool"`
		Received      int64 `json:"received"`
	} `json:"result"`
}

type BlockchainGetSpendProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Signature string `json:"signature"`
	} `json:"result"`
}

type BlockchainCheckSpendProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Good bool `json:"good"`
	} `json:"result"`
}

type BlockchainGetReserveProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Signature string `json:"signature"`
	} `json:"result"`
}

type BlockchainCheckReserveProof struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Good  bool `json:"good"`
		Spent int  `json:"spent"`
		Total int  `json:"total"`
	} `json:"result"`
}

type BlockchainSign struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Signature string `json:"signature"`
	} `json:"result"`
}

type BlockchainVerify struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Good bool `json:"good"`
	} `json:"result"`
}

type BlockchainRescanSpent struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type BlockchainValidateAddress struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Valid bool `json:"valid"`
	} `json:"result"`
}

type BlockchainCreateWallet struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type BlockchainOpenWallet struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type BlockchainCloseWallet struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type BlockchainChangeWalletPassword struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
	} `json:"result"`
}

type NamespaceUpdateRemoteData struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type NamespaceRemoteDataSaveName struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type NamespaceRemoteDataPurchaseName struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type NamespaceRemoteDataDelegateSetAmount struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type NamespaceRemoteDataRenewalStart struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type NamespaceRemoteDataRenewalEnd struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type APIBlockchainStats struct {
	Height                 int `json:"height"`
	Hash                   string `json:"hash"`
	Reward                 int64 `json:"reward"`
	Size                   int64 `json:"size"`
	Version                int `json:"version"`
	VersionBlockHeight     int `json:"versionBlockHeight"`
	NextVersionBlockHeight int `json:"nextVersionBlockHeight"`
	TotalPublicTx                int `json:"totalPublicTx"`
	TotalPrivateTx                int `json:"totalPrivateTx"`
	CirculatingSupply      int64 `json:"circulatingSupply"`
	GeneratedSupply        int64 `json:"generatedSupply"`
	TotalSupply            int64 `json:"totalSupply"`
	EmissionReward         int64 `json:"emissionReward"`
        EmissionHeight         int `json:"emissionHeight"`
    	EmissionTime           int `json:"emissionTime"`
        InflationHeight        int `json:"inflationHeight"`
    	InflationTime          int `json:"inflationTime"`
}

type APIBlockchainBlockData struct {
	Height       int           `json:"height"`
	Hash         string        `json:"hash"`
	Reward       int64         `json:"reward"`
	Time         int           `json:"time"`
	XcashDPOPS   bool          `json:"xcashDPOPS"`
	DelegateName string        `json:"delegateName"`
	Tx           []string      `json:"tx"`
}

type APIBlockchainTxData struct {
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Type          string `json:"type"`
	Sender        string `json:"sender"`
	Receiver      string `json:"receiver"`
	Amount        int64    `json:"amount"`
}

type APIBlockchainProveTx struct {
	Valid  bool `json:"valid"`
	Amount int64  `json:"amount"`
}

type APIBlockchainProveBalance struct {
	Amount       int64 `json:"amount"`
}

type APIBlockchainTxHistory []struct {
	Tx       string `json:"tx"`
	Key      string `json:"key"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   int    `json:"amount"`
	Height   int    `json:"height"`
	Time     int    `json:"time"`
}

type APIBlockchainValidateAddress struct {
	Valid       bool `json:"valid"`
}

type APIBlockchainCreateIntegratedAddress struct {
	IntegratedAddress string `json:"integratedAddress"`
	PaymentID         string `json:"paymentId"`
}



// XCASH DPOPS

type APIDPOPSStats struct {
	MostTotalRoundsDelegateName                   string `json:"mostTotalRoundsDelegateName"`
	MostTotalRounds                               int    `json:"mostTotalRounds"`
	BestBlockVerifierOnlinePercentageDelegateName string `json:"bestBlockVerifierOnlinePercentageDelegateName"`
	BestBlockVerifierOnlinePercentage             int    `json:"bestBlockVerifierOnlinePercentage"`
	MostBlockProducerTotalRoundsDelegateName      string `json:"mostBlockProducerTotalRoundsDelegateName"`
	MostBlockProducerTotalRounds                  int    `json:"mostBlockProducerTotalRounds"`
	TotalVotes                                    int64  `json:"totalVotes"`
	TotalVoters                                   int    `json:"totalVoters"`
	AverageVote                                   int64  `json:"averageVote"`
	VotePercentage                                int    `json:"votePercentage"`
	RoundNumber                                   int    `json:"roundNumber"`
	TotalRegisteredDelegates                      int    `json:"totalRegisteredDelegates"`
	TotalOnlineDelegates                          int    `json:"totalOnlineDelegates"`
	CurrentBlockVerifiersMaximumAmount            int    `json:"currentBlockVerifiersMaximumAmount"`
	CurrentBlockVerifiersValidAmount              int    `json:"currentBlockVerifiersValidAmount"`
}

type APIDPOPSDelegates struct {
	Votes                    int64  `json:"votes"`
	Voters                   int    `json:"voters"`
	IPAdress                 string `json:"IPAdress"`
	DelegateName             string `json:"delegateName"`
	SharedDelegate           bool   `json:"sharedDelegate"`
	SeedNode                 bool   `json:"seedNode"`
	Online                   bool   `json:"online"`
	Fee                      int    `json:"fee"`
	TotalRounds              int    `json:"totalRounds"`
	TotalBlockProducerRounds int    `json:"totalBlockProducerRounds"`
	OnlinePercentage         int    `json:"onlinePercentage"`
}

type APIDPOPSDelegate struct {
	Votes                    int64  `json:"votes"`
	Voters                   int    `json:"voters"`
	IPAdress                 string `json:"IPAdress"`
	DelegateName             string `json:"delegateName"`
	PublicAddress            string `json:"publicAddress"`
	About                    string `json:"about"`
	Website                  string `json:"website"`
	Team                     string `json:"team"`
	Specifications           string `json:"specifications"`
	SharedDelegate           bool   `json:"sharedDelegate"`
	SeedNode                 bool   `json:"seedNode"`
	Online                   bool   `json:"online"`
	Fee                      int    `json:"fee"`
	TotalRounds              int    `json:"totalRounds"`
	TotalBlockProducerRounds int    `json:"totalBlockProducerRounds"`
	OnlinePercentage         int    `json:"onlinePercentage"`
	Rank                     int    `json:"rank"`
}

type APIDPOPSDelegateRoundStats struct {
	TotalBlocksProduced int `json:"totalBlocksProduced"`
	TotalBlockRewards   int64 `json:"totalBlockRewards"`
	AveragePercentage   int `json:"averagePercentage"`
	AverageTime         int `json:"averageTime"`
	BlocksProduced      []struct {
		BlockHeight int `json:"blockHeight"`
		BlockReward int64 `json:"blockReward"`
		Time        int `json:"time"`
	} `json:"blocksProduced"`
}

type APIDPOPSDelegatesVotes struct {
	PublicAddress string `json:"publicAddress"`
	Amount        int64    `json:"amount"`
	ReserveProof  string `json:"reserveProof"`
}

type APIDPOPSVoteDetails struct {
	DelegateName  string `json:"delegateName"`
	Amount        int64    `json:"amount"`
}

type APIDPOPSRoundDetails struct {
	Delegates []string `json:"delegates"`
}

type APIDPOPSLastBlockProducer struct {
	LastBlockProducer string `json:"lastBlockProducer"`
}

// XCASH Namespace
type APINamespaceStats struct {
	TotalNamesRegisteredOrRenewed int `json:"totalNamesRegisteredOrRenewed"`
	TotalVolume                   int64 `json:"totalVolume"`
}

type APINamespaceRegisteredDelegates struct {
	DelegateName  string `json:"delegateName"`
	PublicAddress string `json:"publicAddress"`
	Amount        int64    `json:"amount"`
}

type APINamespaceDelegates struct {
	DelegateName                  string `json:"delegateName"`
	PublicAddress                 string `json:"publicAddress"`
	Amount                        int64    `json:"amount"`
	TotalNamesRegisteredOrRenewed int    `json:"totalNamesRegisteredOrRenewed"`
	TotalVolume                   int64    `json:"totalVolume"`
}

type APINamespaceName struct {
	Address        string `json:"address"`
	Saddress       string `json:"saddress"`
	Paddress       string `json:"paddress"`
	Expires        int    `json:"expires"`
	DelegateName   string `json:"delegateName"`
	DelegateAmount int64    `json:"delegateAmount"`
}

type APINamespaceNameStatus struct {
	Status bool `json:"status"`
}

type APINamespaceAddressStatus struct {
	Status string `json:"status"`
}

type APINamespaceNameToAddress struct {
	Address  string `json:"address"`
	Saddress string `json:"saddress"`
	Paddress string `json:"paddress"`
}

type APINamespaceAddressToName struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

// Xpayment Twitter 
type APIXpaymentTwitterStats struct {
	TotalUsers                     int `json:"totalUsers"`
	AvgTipAmount                   int `json:"avgTipAmount"`
	TotalDeposits                  int `json:"totalDeposits"`
	TotalWithdraws                 int `json:"totalWithdraws"`
	TotalTipsPublic                int `json:"totalTipsPublic"`
	TotalTipsPrivate               int `json:"totalTipsPrivate"`
	TotalVolumeSentPublic          int64 `json:"totalVolumeSentPublic"`
	TotalVolumeSentPrivate         int64 `json:"totalVolumeSentPrivate"`
	TotalTipsLastDayPublic         int `json:"totalTipsLastDayPublic"`
	TotalTipsLastDayPrivate        int `json:"totalTipsLastDayPrivate"`
	TotalVolumeSentLastDayPublic   int64 `json:"totalVolumeSentLastDayPublic"`
	TotalVolumeSentLastDayPrivate  int64 `json:"totalVolumeSentLastDayPrivate"`
	TotalTipsLastHourPublic        int `json:"totalTipsLastHourPublic"`
	TotalTipsLastHourPrivate       int `json:"totalTipsLastHourPrivate"`
	TotalVolumeSentLastHourPublic  int64 `json:"totalVolumeSentLastHourPublic"`
	TotalVolumeSentLastHourPrivate int64 `json:"totalVolumeSentLastHourPrivate"`
}

type APIXpaymentTwitterStatsPerDay struct {
	Time   int `json:"time"`
	Amount int `json:"amount"`
	Volume int64 `json:"volume"`
}

type TopTips struct {
		Username string `json:"username"`
		Tips     int    `json:"tips"`
	}
	
	type TopVolumes struct {
		Username string `json:"username"`
		Volume   int    `json:"volume"`
}

type APIXpaymentTwitterTopStats struct {
	TopTips []TopTips
	TopVolumes []TopVolumes
}

type APIXpaymentTwitterRecentTips struct {
	TweetID  string `json:"tweetId"`
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Amount   int64    `json:"amount"`
	Time     int    `json:"time"`
	Type     string `json:"type"`
}
