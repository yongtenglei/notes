package setting

type AccountServerConf struct {
	AccountWebServerConfig *AccountWebServerConfig `mapstructure:"web_server"`
	AccountWebClientConfig *AccountWebClientConfig `mapstructure:"web_client"`
}

type AccountWebServerConfig struct {
	Name string   `mapstructure:"name"`
	Host string   `mapstructure:"host"`
	Port int32    `mapstructure:"port"`
	Tags []string `mapstructure:"tags"`
}

type AccountWebClientConfig struct {
	ID   string   `mapstructure:"id"`
	Name string   `mapstructure:"name"`
	Host string   `mapstructure:"host"`
	Port int32    `mapstructure:"port"`
	Tags []string `mapstructure:"tags"`
}
