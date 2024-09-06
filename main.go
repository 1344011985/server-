package main

import "server-/api"

import (
	"log"
	"net/http"
)

func main() {
	// 创建一个新的路由器
	router := api.NewRouter()

	// 启动服务器
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
