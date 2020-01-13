package linkaja

import (
	"bytes"
	"fmt"
	"strings"
)

func GenerateItems(items []PublicTokenItemRequest) string {
	var is string
	for i, v := range items {
		if i > 0 {
			is = is + ","
		}
		is = is + fmt.Sprintf("[\"%v\", \"%v\", \"%v\"]", v.Name, v.Price, v.Quantity)
	}

	return fmt.Sprintf("[%v]", is)
}

func PadRight(str string, item string, length int) string {
	strLength := len(strings.TrimSpace(str))
	count := length - strLength
	return str + strings.Repeat(item, count)
}

func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize < 0 || blockSize > 256 {
		return nil, fmt.Errorf("pkcs7: Invalid block size %d", blockSize)
	}

	padLen := 16 - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...), nil
}
