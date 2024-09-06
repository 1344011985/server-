package api

import (
	"net/http"
	controllers "server-/controller"
)

// NewRouter 创建一个新的 HTTP 路由器
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	// 文件下载和上传路由
	router.HandleFunc("/download", controllers.DownloadFileHandler) // GET 请求下载文件
	router.HandleFunc("/upload", controllers.UploadFileHandler)     // POST 请求上传文件

	return router
}
