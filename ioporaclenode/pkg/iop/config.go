package iop

type Config struct {
	BindAddress string
	PrivateKey  string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	IOTA        IOTAConfig
}

type ContractsConfig struct {
	RegistryContractAddress string
	OracleContractAddress   string
	DistKeyContractAddress  string
}

type EthereumConfig struct {
	TargetAddress string
	SourceAddress string
	PrivateKey    string
}

type IOTAConfig struct {
	Address          string
	Zmq              string
	Seed             string
	BroadcastAddress string
}
