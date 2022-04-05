package utils

import (
	b64 "encoding/base64"
)

/*
	StrToBase64 encoded directory path to base64 string.
	For example, if directory path is "dir1/dir2/dir3/",
	encoded directory path is "ZGlyMS9kaXIyL2RpcjM="
*/
func StrToBase64(directoryPath string) (base64String string) {
	encodedStr := b64.StdEncoding.EncodeToString([]byte(directoryPath))
	return encodedStr
}

/*
	Base64ToStr decode a base64 string to directory path string.
	For example, if base64 string is "ZGlyMS9kaXIyL2RpcjM=",
	decoded base64 string is "dir1/dir2/dir3/"
*/
func Base64ToStr(encodedStr string) (dirPath string) {
	decodedStr, _ := b64.StdEncoding.DecodeString(encodedStr)
	return string(decodedStr)
}
