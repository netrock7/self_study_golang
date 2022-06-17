/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	// sc := []constant.ServerConfig{
	// 	{
	// 		IpAddr: "172.20.27.162",
	// 		Port:   8848,
	// 	},
	// }
	//or a more graceful way to create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"172.20.27.162",
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos")),
	}

	// cc := constant.ClientConfig{
	// 	NamespaceId:         "26ad453c-c8b0-4edf-aa52-a317c40032b8", //namespace id
	// 	TimeoutMs:           5000,
	// 	NotLoadCacheAtStart: true,
	// 	LogDir:              "/tmp/nacos/log",
	// 	CacheDir:            "/tmp/nacos/cache",
	// 	LogLevel:            "debug",
	// }
	//or a more graceful way to create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("26ad453c-c8b0-4edf-aa52-a317c40032b8"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//publish config
	//config key=dataId+group+namespaceId
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "hello world!",
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "hello world!",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}

	//get config
	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	fmt.Println("GetConfig,config :" + content)

	//Listen config change,key=dataId+group+namespaceId.
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data-2",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "test-listen",
	})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "test-listen",
	})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)

	//cancel config change
	err = client.CancelListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)

	searchPage, _ := client.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Printf("Search config:%+v \n", searchPage)
}
