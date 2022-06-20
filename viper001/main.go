package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadYaml(NacosConfig Config) (NConfig Config) {
	v := viper.New()
	v.AddConfigPath("./viper001/config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("v: %v\n", v)

	// NacosConfig := Config{}

	NacosConfig.Server.IpAddr = v.GetString("server.IpAddr")
	NacosConfig.Server.ContextPath = v.GetString("server.ContextPath")
	NacosConfig.Server.Port = v.GetString("server.Port")
	NacosConfig.Server.Scheme = v.GetString("server.Scheme")
	NacosConfig.Client.NamespaceId = v.GetString("client.NamespaceId")
	NacosConfig.Client.TimeoutMs = v.GetString("client.TimeoutMs")
	NacosConfig.Client.NotLoadCacheAtStart = v.GetString("client.NotLoadCacheAtStart")
	NacosConfig.Client.LogDir = v.GetString("client.LogDir")
	NacosConfig.Client.CacheDir = v.GetString("client.CacheDir")
	NacosConfig.Client.LogLevel = v.GetString("client.LogLevel")
	// fmt.Printf("s: %v\n", Config.Server.IpAddr)

	return NacosConfig
}

func main() {
	NacosConfig := Config{}
	ReadYaml(NacosConfig)
	// fmt.Printf("NacosConfig: %v\n", NacosConfig)

}

type Config struct {
	Server Server `json:"server"`
	Client Client `json:"client"`
}

type Server struct {
	IpAddr      string `json:"IpAddr"`
	ContextPath string `json:"ContextPath"`
	Port        string `json:"Port"`
	Scheme      string `json:"Scheme"`
}

type Client struct {
	NamespaceId         string `json:"NamespaceId"`
	TimeoutMs           string `json:"TimeoutMs"`
	NotLoadCacheAtStart string `json:"NotLoadCacheAtStart"`
	LogDir              string `json:"LogDir"`
	CacheDir            string `json:"CacheDir"`
	LogLevel            string `json:"LogLevel"`
}
