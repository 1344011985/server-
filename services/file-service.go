package services

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadFile 处理文件的下载操作
func DownloadFile(w http.ResponseWriter, filePath string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.Header().Set("Content-Type", "application/octet-stream")

	// 将文件内容复制到响应体中
	_, err = io.Copy(w, file)
	return err
}

// UploadFile 处理文件的上传操作
func UploadFile(w http.ResponseWriter, r *http.Request) error {
	file, handler, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个目标文件
	dst, err := os.Create(handler.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件中
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}
