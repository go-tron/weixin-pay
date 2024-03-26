package weixinPay

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

func ReplaceUrl(url string, data map[string]interface{}) string {
	if !strings.Contains(url, "{") {
		return url
	}
	path := "{(.*?)}"
	exp := regexp.MustCompile(path)
	result := exp.FindAllStringSubmatch(url, -1)
	if len(result) > 0 {
		for _, v := range result {
			if r := data[v[1]]; r != nil {
				url = strings.Replace(url, v[0], fmt.Sprintf("%v", r), 1)
				delete(data, v[1])
			}
		}
	}
	return url
}

func AESGCMDecrypter(aesKey, nonce, ciphertext, additiontext string) ([]byte, error) {
	textByte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := aesgcm.Open(nil, []byte(nonce), textByte, []byte(additiontext))
	if err != nil {
		return nil, err
	}
	return plaintext, err
}
