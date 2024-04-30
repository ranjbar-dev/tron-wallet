package tronwallet

// ========== address structs ========== //

type Address struct {
	Hex    string
	Base58 string
}

// ========== transaction structs ========== //

type Transaction struct {
	RawData    RawData `json:"raw_data"`
	TxId       string  `json:"txID"`
	RawDataHex string  `json:"raw_data_hex"`
}

type Contract struct {
	Type      string            `json:"type"`
	Parameter ContractParameter `json:"parameter"`
}

type ContractParameter struct {
	TypeUrl string        `json:"type_url"`
	Value   ContractValue `json:"value"`
}

type ContractValue struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	Amount       int64  `json:"amount"`
}

type RawData struct {
	Contract      []Contract
	RefBlockBytes string `json:"ref_block_bytes"`
	RefBlockHash  string `json:"ref_block_hash"`
	Expiration    int64  `json:"expiration"`
	Timestamp     int64  `json:"timestamp"`
}
