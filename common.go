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

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
