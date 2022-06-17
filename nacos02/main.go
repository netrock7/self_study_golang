package main

import (
"github.com/spf13/viper"
remote "github.com/yoyofxteam/nacos-viper-remote"
)

runtime_viper := viper.New()
// 配置 Viper for Nacos 的远程仓库参数
remote.SetOptions(&remote.Option{
   Url:         "172.20.27.162",            // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
   Port:        8848,                     // nacos server端口号
   NamespaceId: "26ad453c-c8b0-4edf-aa52-a317c40032b8",               // nacos namespace
   GroupName:   "DEFAULT_GROUP",        // nacos group
   Config:     remote.Config{ DataId: "20220614" }, // nacos DataID
   Auth:        nil,                    // 如果需要验证登录,需要此参数
})

err := remote_viper.AddRemoteProvider("nacos", "172.20.27.162", "")
remote_viper.SetConfigType("yaml")

_ = remote_viper.ReadRemoteConfig()             //sync get remote configs to remote_viper instance memory . for example , remote_viper.GetString(key)
_ = remote_viper.WatchRemoteConfigOnChannel()   //异步监听Nacos中的配置变化，如发生配置更改，会直接同步到 viper实例中。

appName := remote_viper.GetString("key")   // sync get config by key

fmt.Println(appName)

// 这里不是必须的，只为了监控Demo中的配置变化，并打印出来
go func() {
    for {
        time.Sleep(time.Second * 30) // 每30秒检查配置是否发生变化 
        appName = config_viper.GetString("yoyogo.application.name")
        fmt.Println(appName)
    }
}