package main

import (
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

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
	nonce, _ := hex.DecodeString("1234567812345678")
	messageBytes, _ := hex.DecodeString(message)
	keyBytes, _ := hex.DecodeString(key)
	blockCipher, _ := aes.NewCipher(keyBytes[:16]) // 8 byte nonce

	var mi []byte
	i := 0
	j := len(messageBytes)
	ret := make([]byte, 0)
	ret = append(ret, nonce...)
	for {
		ci := append(nonce, eightByteCounter(int64(i))...)
		if (16 * i) >= j {
			break
		}
		if j < 16 {
			padLength := 16 - j
			mi = messageBytes[i:j]
			for k := 0; k < padLength; k++ {
				mi = append(mi, byte(padLength))
			}
			i = j + 1 //make sure we exit loop
		} else if j-(16*i) >= 16 {
			mi = messageBytes[(16 * i) : (16*i)+16]
			i += 1
		} else {
			padLength := 16 - (j - (16 * i))
			mi = messageBytes[(16 * i):j]
			for k := 0; k < padLength; k++ {
				mi = append(mi, byte(padLength))
			}
			i = j + 1 //make sure we exit loop
		}
		dst := make([]byte, 16)
		blockCipher.Encrypt(dst, ci)
		ret = append(ret, byteArrayXor(mi, dst)...)
	}
	return hex.EncodeToString(ret)
}

func decrypt(message string, key string) string {
	ciphertextBytes, _ := hex.DecodeString(message)
	nonce := ciphertextBytes[:16]
	keyBytes, _ := hex.DecodeString(key)
	blockCipher, _ := aes.NewCipher(keyBytes[:16]) // 8 byte nonce

	var ci, mi []byte
	i := 16
	j := len(ciphertextBytes)
	ret := make([]byte, 0)
	ret = append(ret, ci...)
	for {
		if i >= j {
			break
		}
		ci = ciphertextBytes[i : i+16]
		i = i + 16

		dst := make([]byte, 16)
		blockCipher.Encrypt(dst, nonce)
		mi = byteArrayXor(dst, ci)
		ret = append(ret, mi...)
		nonce = incrementNonce(nonce)
	}
	return hex.EncodeToString(ret)
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

func Int64ToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func eightByteCounter(num int64) []byte {
	ret := make([]byte, 8)
	numBytes := []byte(fmt.Sprint(num))
	ret = append(ret[0:8-len(numBytes)], numBytes...)
	return ret
}

func incrementNonce(nonce []byte) []byte {
	ret := nonce
	if ret[15] < 255 {
		ret[15] = ret[15] + 1
	} else {
		ret[15] = 0
		ret[14] = ret[14] + 1
	}
	return ret
}
