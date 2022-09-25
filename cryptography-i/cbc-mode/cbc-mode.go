package main

import (
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
)

/*
 Obviously insecure CBC implementation - For education purposes only
*/
func main() {
	decryptFlag := flag.Bool("decrypt", false, "otherwise, encrypt")
	inputFile := flag.String("input", "", "either message or ciphertext")

	flag.Parse()

	data, _ := ioutil.ReadFile(*inputFile)
	key, _ := ioutil.ReadFile("./key.txt")

	if !(*decryptFlag) {
		fmt.Println(encrypt(string(data), string(key)))
	} else {
		fmt.Println(decrypt(string(data), string(key)))
	}
}

func encrypt(message string, key string) string {
	initializationVector, _ := hex.DecodeString("12345678123456781234567812345678")
	messageBytes, _ := hex.DecodeString(message)
	keyBytes, _ := hex.DecodeString(key)
	blockCipher, _ := aes.NewCipher(keyBytes[:16])

	var mi []byte
	i := 0
	j := len(messageBytes)
	ci := initializationVector
	ret := make([]byte, 0)
	ret = append(ret, ci...)
	for {
		if i > j {
			break
		}
		if j < 16 {
			padLength := 16 - j
			mi = messageBytes[i:j]
			for k := 0; k < padLength; k++ {
				mi = append(mi, byte(padLength))
			}
			i = j + 1 //make sure we exit loop
		} else if j-i >= 16 {
			mi = messageBytes[i : i+16]
			i += 16
		} else {
			padLength := 16 - (j - i)
			mi = messageBytes[i:j]
			for k := 0; k < padLength; k++ {
				mi = append(mi, byte(padLength))
			}
			i = j + 1 //make sure we exit loop
		}
		dst := make([]byte, 16)
		blockCipher.Encrypt(dst, byteArrayXor(mi, ci))
		ci = dst
		ret = append(ret, ci...)
	}
	return hex.EncodeToString(ret)
}

func decrypt(ciphertext string, key string) string {
	ciphertextBytes, _ := hex.DecodeString(ciphertext)
	initializationVector := ciphertextBytes[:16]
	keyBytes, _ := hex.DecodeString(key)
	blockCipher, _ := aes.NewCipher(keyBytes[:16])

	var ci, mi []byte
	i := 16
	j := len(ciphertextBytes)
	ret := make([]byte, 0)
	ret = append(ret, ci...)
	for {
		if i == j {
			break
		}
		ci = ciphertextBytes[i : i+16]
		i = i + 16

		dst := make([]byte, 16)
		blockCipher.Decrypt(dst, ci)
		mi = byteArrayXor(dst, initializationVector)
		initializationVector = ci
		ret = append(ret, mi...)
	}
	finalByte := int(ret[len(ret)-1])
	return hex.EncodeToString(ret[:len(ret)-finalByte])
}

func byteArrayXor(arr1 []byte, arr2 []byte) []byte {
	maxBytes := Min(len(arr1), len(arr2))
	ret := make([]byte, maxBytes)
	for i := 0; i < maxBytes; i++ {
		ret[i] = arr1[i] ^ arr2[i]
	}
	return ret
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
