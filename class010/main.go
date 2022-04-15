package main

// import "fmt"

// // var _ Study = (*study)(nil)

// type study struct {
// 	Name string
// }

// type Study interface {
// 	Listen(message string) string
// }

// func main() {
// 	fmt.Println("hello world")
// }

import (
	"fmt"
	"os"
	"runtime"
)

var (
	GoVersion  = runtime.Version()
	CommitId   string
	BranchName string
	BuildTime  string
	AppVersion string
)

func main() {
	fmt.Printf("go version: %s\r\n", GoVersion)
	fmt.Printf("git commit ID: %s\r\n", CommitId)
	fmt.Printf("git branch name: %s\r\n", BranchName)
	fmt.Printf("app build time: %s\r\n", BuildTime)
	fmt.Printf("app version: %s\r\n", AppVersion)
	// 打印完退出
	os.Exit(0)
}
