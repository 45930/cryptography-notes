package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	TARGET_CIPHERTEXT := "32510ba9babebbbefd001547a810e67149caee11d945cd7fc81a05e9f85aac650e9052ba6a8cd8257bf14d13e6f0a803b54fde9e77472dbff89d71b57bddef121336cb85ccb8f3315f4b52e301d16e9f52f904"

	targer_ciphertext_hex, _ := hex.DecodeString(TARGET_CIPHERTEXT)
	fmt.Printf("Target cipher length: %d \n", len(targer_ciphertext_hex))

	file, err := os.Open("./cyphertexts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ciphertexts := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ciphertext, err := hex.DecodeString(scanner.Text())
		// fmt.Println(string(ciphertext))
		if err != nil {
			log.Fatal(err)
		}
		ciphertexts = append(ciphertexts, ciphertext)
	}

	file, err = os.Create("./m1xors.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	m1 := ciphertexts[0]
	for i := 1; i < len(ciphertexts); i++ {
		msg_xors := byteArrayXor(m1, ciphertexts[i])
		fmt.Printf("M1 XOR M%d:\n", i+1)
		fmt.Println(strings.Join(strings.Fields(fmt.Sprint(msg_xors)), ", "))
		fmt.Printf("\n\n\n")
		_, _ = writer.WriteString(strings.Join(strings.Fields(fmt.Sprint(msg_xors)), ", ") + "\n")
	}
	writer.Flush()
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
