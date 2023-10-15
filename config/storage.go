package config

import "fmt"

func GetStorageProxyUrl () string {
	url:= fmt.Sprintf("%s/api/file", Config("STORAGE_URL"))
	return url
}