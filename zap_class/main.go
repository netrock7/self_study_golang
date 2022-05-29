// github.com/bigwhite/experiments/tree/master/uber-zap-advanced-usage/demo3/main.go
package main

import (
	"go.uber.org/zap"
)

// "github.com/bigwhite/zap-usage/pkg/log"

// func main() {
// 	file1, err := os.OpenFile("./access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	file2, err := os.OpenFile("./error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var tops = []log.TeeOption{
// 		{
// 			W: file1,
// 			Lef: func(lvl log.Level) bool {
// 				return lvl <= log.InfoLevel
// 			},
// 		},
// 		{
// 			W: file2,
// 			Lef: func(lvl log.Level) bool {
// 				return lvl > log.InfoLevel
// 			},
// 		},
// 	}

// 	logger := log.NewTee(tops)
// 	log.ResetDefault(logger)

// 	log.Info("demo3:", log.String("app", "start ok"),
// 		log.Int("major version", 3))
// 	log.Error("demo3:", log.String("app", "crash"),
// 		log.Int("reason", -1))

// }

func main() {
	zap.Info("call foo", zap.String("url", "https://tonybai.com"), zap.Int("attempt", 3))
}
