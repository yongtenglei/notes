package setting

type AccountServerConf struct {
	AccountWebServerConfig *AccountWebServerConfig `mapstructure:"web_server" json:"web_server"`
	AccountWebClientConfig *AccountWebClientConfig `mapstructure:"web_client" json:"web_client"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"data_id"`
	Group     string `mapstructure:"group"`
	Port      int32  `mapstructure:"port"`
}
type AccountWebServerConfig struct {
	Name string   `mapstructure:"name" json:"name"`
	Host string   `mapstructure:"host" json:"host"`
	Port int32    `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}

type AccountWebClientConfig struct {
	ID   string   `mapstructure:"id" json:"id"`
	Name string   `mapstructure:"name" json:"name"`
	Host string   `mapstructure:"host" json:"host"`
	Port int32    `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}
