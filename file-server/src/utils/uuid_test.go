package utils

import (
	"testing"
)

func TestBase64(t *testing.T) {
	data := "dir1/dir2/dir3" // Encoded string: ZGlyMS9kaXIyL2RpcjM=
	encodedStr := StrToBase64(data)
	decodedStr := Base64ToStr(encodedStr)

	if decodedStr != data {
		t.Logf(`Decoding result is different to original data. \n
			Original: %s
			Decoded data: %s`, data, decodedStr)
	}
}
