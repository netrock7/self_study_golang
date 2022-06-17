package main

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// type DBConfig struct {
// 	DBUser     string `json:"dbuser"`
// 	DBPassword string `json:"dbpassword"`
// 	DBHost     string `json:"DBHost"`
// 	DBPort     int    `json:"DBPort"`
// }

// type RedisConfig struct {
// 	RedisHost     string `json:"redishost"`
// 	RedisPassword string `json:"redispassword"`
// 	RedisPort     int    `json:"redisport"`
// 	RedisDB       int    `json:"redisdb"`
// 	RedisExpire   int    `json:"redisexpire"`
// }

// type JWT struct {
// 	JWTKey string `json:"jwtkey"`
// }

// type Server struct {
// 	MySQLConfig DBConfig    `json:"MySQL"`
// 	RedisUser   RedisConfig `json:"redisUser"`
// 	RedisGate   RedisConfig `json:"redisGate"`
// 	JWT         JWT         `json:"jwt"`
// }

type Server struct {
	MySQL     MySQL     `json:"MySQL"`
	RedisUser RedisUser `json:"redisUser"`
	RedisGate RedisGate `json:"redisGate"`
	Jwt       Jwt       `json:"jwt"`
}
type MySQL struct {
	DBHost     string `json:"DBHost"`
	Dbuser     string `json:"dbuser"`
	Dbpassword string `json:"dbpassword"`
	DBPort     int    `json:"DBPort"`
}
type RedisUser struct {
	Redishost     string `json:"redishost"`
	Redispassword string `json:"redispassword"`
	Redisport     int    `json:"redisport"`
	Redisdb       int    `json:"redisdb"`
	Redisexpire   int    `json:"redisexpire"`
}
type RedisGate struct {
	Redishost     string `json:"redishost"`
	Redispassword string `json:"redispassword"`
	Redisport     int    `json:"redisport"`
	Redisdb       int    `json:"redisdb"`
	Redisexpire   int    `json:"redisexpire"`
}
type Jwt struct {
	Jwtkey string `json:"jwtkey"`
}

func main() {
	// 配置中心信息
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "172.20.27.162",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 客户端信息
	clientConfig := constant.ClientConfig{
		NamespaceId:         "26ad453c-c8b0-4edf-aa52-a317c40032b8", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	//// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "geo",
		Group:  "base"})

	if err != nil {
		panic(err)
	}

	var ServConf Server
	err = json.Unmarshal([]byte(content), &ServConf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("ServConf: %v\n", ServConf)

	// fmt.Printf("conf: %v\n", string(conf))

	//监听配置中心变化
	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "geo",
	// 	Group:  "base",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// time.Sleep(time.Hour)
}
