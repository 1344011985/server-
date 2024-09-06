package conf

import "os"

// GetConfig 获取配置，例如端口号
func GetConfig(key string) string {
	return os.Getenv(key)
}
