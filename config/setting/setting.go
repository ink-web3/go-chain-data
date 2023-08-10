package setting

type DbConfig struct {
	DbType    string
	DbName    string
	Host      string
	Username  string
	Pwd       string
	Charset   string
	ParseTime bool
}
type BlockChainConfig struct {
	RpcUrl string
}
