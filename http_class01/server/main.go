package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了给你返回GET"))
		break
	case "POST":
		b, _ := io.ReadAll(req.Body)
		// bb, _ := json.Marshal(b)
		fmt.Println(req.Header["Test"])
		header := res.Header()
		header["Authenrization"] = []string{"sdjfo3uw9ru0jf0q9urt0u32r034"}
		res.WriteHeader(http.StatusOK)
		res.Write(b)
		// res.Write(([]byte("我收到了给你返回POST")))
		break
	}
	// res.Write([]byte("我收到了给你返回"))
	fmt.Println("成功收到了数据ß")
}

// func main() {
// 	http.HandleFunc("/test", handler)
// 	http.ListenAndServe(":8080", nil)
// 	// http.Handle()
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}
