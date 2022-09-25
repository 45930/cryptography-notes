package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	VIDEO_FILE := "target-video.mp4"

	data, _ := ioutil.ReadFile(VIDEO_FILE)

	blockSize := 1024
	dataLength := len(data)

	i := 0
	blocks := make([][]byte, 0)
	for {
		if dataLength-i <= blockSize {
			block := data[i:dataLength]
			blocks = append(blocks, block)
			break
		} else {
			block := data[i : i+blockSize]
			blocks = append(blocks, block)
			i += blockSize
		}
	}

	prevHash := make([]byte, 0)
	for j := len(blocks) - 1; j >= 0; j-- {
		shaHash := sha256.New()
		block := blocks[j]
		if j == len(blocks)-1 {
			shaHash.Write(block)
			prevHash = shaHash.Sum(nil)
		} else {
			shaHash.Write(append(block, prevHash...))
			prevHash = shaHash.Sum(nil)
		}
	}

	fmt.Println(hex.EncodeToString(prevHash))

}
