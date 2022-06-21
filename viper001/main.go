package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var NacosConfig *Config

func ReadYaml() {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("v: %v\n", v)

	// NacosConfig := Config{}
	err = v.Unmarshal(&NacosConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(NacosConfig.Server.IpAddr)

	// NacosConfig.Server.IpAddr = v.GetString("server.IpAddr")
	// NacosConfig.Server.ContextPath = v.GetString("server.ContextPath")
	// NacosConfig.Server.Port = v.GetInt64("server.Port")
	// NacosConfig.Server.Scheme = v.GetString("server.Scheme")
	// NacosConfig.Client.NamespaceId = v.GetString("client.NamespaceId")
	// NacosConfig.Client.TimeoutMs = v.GetInt64("client.TimeoutMs")
	// NacosConfig.Client.NotLoadCacheAtStart = v.GetBool("client.NotLoadCacheAtStart")
	// NacosConfig.Client.LogDir = v.GetString("client.LogDir")
	// NacosConfig.Client.CacheDir = v.GetString("client.CacheDir")
	// NacosConfig.Client.LogLevel = v.GetString("client.LogLevel")
	// fmt.Printf("s: %v\n", Config.Server.IpAddr)

	// return NacosConfig
}

func main() {
	// NacosConfig := Config{}
	ReadYaml()
	// fmt.Printf("NacosConfig: %v\n", NacosConfig)

}

type Config struct {
	Server Server `json:"server"`
	Client Client `json:"client"`
}

type Server struct {
	IpAddr      string `json:"IpAddr"`
	ContextPath string `json:"ContextPath"`
	Port        int64  `json:"Port"`
	Scheme      string `json:"Scheme"`
}

type Client struct {
	NamespaceId         string `json:"NamespaceId"`
	TimeoutMs           int64  `json:"TimeoutMs"`
	NotLoadCacheAtStart bool   `json:"NotLoadCacheAtStart"`
	LogDir              string `json:"LogDir"`
	CacheDir            string `json:"CacheDir"`
	LogLevel            string `json:"LogLevel"`
}
