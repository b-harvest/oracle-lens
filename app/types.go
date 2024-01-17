package app

import "bharvest.io/oracle-lens/wallet"

type Config struct {
	General struct {
		Network    string `toml:"network"`
		EthAPI     string `toml:"eth_api"`
		ListenPort int    `toml:"listen_port"`
		Period     int    `toml:"period"`
	} `toml:"general"`
	Injective struct {
		Enable       bool   `toml:"enable"`
		GRPC         string `toml:"grpc"`
		ValidatorAcc string `toml:"validator_acc"`
		NonceDiff    uint64 `toml:"nonce_diff"`
		Wallet       *wallet.Wallet
	} `toml:"injective"`
	UmeeOracle struct {
		Enable       bool    `toml:"enable"`
		GRPC         string  `toml:"grpc"`
		ValidatorAcc string  `toml:"validator_acc"`
		MinUptime    float64 `toml:"min_uptime"`
		Wallet       *wallet.Wallet
	} `toml:"umee-oracle"`
	Tg struct {
		Enable bool   `toml:"enable"`
		Token  string `toml:"token"`
		ChatID string `toml:"chat_id"`
	} `toml:"tg"`
}
