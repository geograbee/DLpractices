package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	for {
		fmt.Println("Choose operation:")
		fmt.Println("1. Show number of posible unique keys variants")
		fmt.Println("2. Show examples of posible unique keys variants")
		fmt.Println("3. Brute force")
		fmt.Println("0. Exit")
		fmt.Print("Enter the number of chosen operation: ")
		var input int
		fmt.Scan(&input)
		if input == 1 {
			ShowKeysNumber()
		}
		if input == 2 {
			ShowRandomKeys()
		}
		if input == 3 {
			BruteForce()
		}
		if input == 0 {
			break
		}
	}
}

func ShowKeysNumber() {
	numOfVariants8 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(8), nil)
	fmt.Println("Number of posible unique 8-bit keys variants is", numOfVariants8)
	numOfVariants16 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(16), nil)
	fmt.Println("Number of posible unique 16-bit keys variants is", numOfVariants16)
	numOfVariants32 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(32), nil)
	fmt.Println("Number of posible unique 32-bit keys variants is", numOfVariants32)
	numOfVariants64 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(64), nil)
	fmt.Println("Number of posible unique 64-bit keys variants is", numOfVariants64)
	numOfVariants128 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(128), nil)
	fmt.Println("Number of posible unique 128-bit keys variants is", numOfVariants128)
	numOfVariants256 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(256), nil)
	fmt.Println("Number of posible unique 256-bit keys variants is", numOfVariants256)
	numOfVariants512 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(512), nil)
	fmt.Println("Number of posible unique 512-bit keys variants is", numOfVariants512)
	numOfVariants1024 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(1024), nil)
	fmt.Println("Number of posible unique 1024-bit keys variants is", numOfVariants1024)
	numOfVariants2048 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(2048), nil)
	fmt.Println("Number of posible unique 2048-bit keys variants is", numOfVariants2048)
	numOfVariants4096 := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(4096), nil)
	fmt.Println("Number of posible unique 4096-bit keys variants is", numOfVariants4096)
}

func ShowRandomKeys() {
	randKey8, _ := randomHex(8)
	fmt.Println("Random 8-bit key - 0x" + randKey8)
	randKey16, _ := randomHex(16)
	fmt.Println("Random 16-bit key - 0x" + randKey16)
	randKey32, _ := randomHex(32)
	fmt.Println("Random 32-bit key - 0x" + randKey32)
	randKey64, _ := randomHex(64)
	fmt.Println("Random 64-bit key - 0x" + randKey64)
	randKey128, _ := randomHex(128)
	fmt.Println("Random 128-bit key - 0x" + randKey128)
	randKey256, _ := randomHex(256)
	fmt.Println("Random 256-it key - 0x" + randKey256)
	randKey512, _ := randomHex(512)
	fmt.Println("Random 512-bit key - 0x" + randKey512)
	randKey1024, _ := randomHex(1024)
	fmt.Println("Random 1024-bit key - 0x" + randKey1024)
	randKey2048, _ := randomHex(2048)
	fmt.Println("Random 2048-bit key - 0x" + randKey2048)
	randKey4096, _ := randomHex(4096)
	fmt.Println("Random 4096-bit key - 0x" + randKey4096)
}

func randomHex(n int) (string, error) {
	var bytes []byte = make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func BruteForce() {
	/*var lenghts []byte = []byte{8, 16, 32, 64, 128}
	randomLength := math.rand rand.Read(lenghts)
	pick := lenghts[randomLength]
	Key, _ := randomHex(pick)
	for i := make([]byte, pick); ; {

	}
	i := big.Int(0)
	for i.Cmp() == 1{*/
}
