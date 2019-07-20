package template

type PeerTemplate struct {
	Name              string
	Labels            map[string]string
	Image             string
	LogLevel          string
	MSPID             string
	InternalIP        string
	PublishMainPort   uint32
	PublishCCPort     uint32
	PublishEventPort  uint32
	PublishBlockPort  uint32
	ExternalEndpoint  string
	BootstrapEndpoint string

	CoreChaincodeBuilder       string
	CoreChaincodeGolangRuntime string

	Files []*FileConfig
	GM    bool
}

type FileConfig struct {
	ID   string
	Name string
	Path string
}