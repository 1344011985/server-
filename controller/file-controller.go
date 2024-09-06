package controllers

import (
	"net/http"
	"server-/services"
	"server-/utils"
)

// DownloadFileHandler 处理文件下载请求
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "File parameter is missing")
		return
	}

	// 调用服务层的文件下载逻辑
	err := services.DownloadFile(w, filePath)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// UploadFileHandler 处理文件上传请求
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB 限制
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Could not parse multipart form")
		return
	}

	// 调用服务层的文件上传逻辑
	err = services.UploadFile(w, r)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "File uploaded successfully"})
}
