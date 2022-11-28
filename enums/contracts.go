package enums

import "github.com/ranjbar-dev/tron-wallet/util"

type ContractAddress string

func (ca ContractAddress) Base58() string {
	return string(ca)
}

func (ca ContractAddress) Hex() string {
	addr, _ := util.Base58ToAddress(string(ca))
	return addr.Hex()
}

func (ca ContractAddress) Bytes() []byte {
	addr, _ := util.Base58ToAddress(string(ca))
	return addr.Bytes()
}

func CreateContractAddress(contractAddressBase58 string) ContractAddress {
	return ContractAddress(contractAddressBase58)
}

const (
	SHASTA_Tether_USDT ContractAddress = "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs"
)

const (
	MAIN_Tether_USDT            ContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	MAIN_Decentralized_USD_USDD ContractAddress = "TPYmHEhy5n8TCEfYGqW2rPxsghSfzghPDn"
	MAIN_USD_COIN_USDC          ContractAddress = "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8"
	MAIN_TRUE_USD_TSDD          ContractAddress = "TUpMhErZL2fhh4sVNULAbNKLokS4GjC1F4"
	MAIN_JUST_USDJ              ContractAddress = "TMwFHYXLJaRUPeW6421aqXL4ZEzPRFGkGT"
	MAIN_JUST_JST               ContractAddress = "TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9"
	MAIN_BitTorrent_BTT         ContractAddress = "TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4"
)
