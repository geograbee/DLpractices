package main

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

func main() {
	for {
		fmt.Println("Choose operation:")
		fmt.Println("1. Convert hex to little-endian")
		fmt.Println("2. Convert hex to big-endian")
		fmt.Println("3. Convert little-endian to hex")
		fmt.Println("4. Convert big-endian to hex")
		fmt.Println("0. Exit")
		fmt.Print("Enter the number of chosen operation: ")
		var input int
		fmt.Scan(&input)
		if input == 1 {
			hexToLittleEndian()
		}
		if input == 2 {
			hexToBigEndian()
		}
		if input == 3 {
			littleEndianToHex()
		}
		if input == 4 {
			bigEndianToHex()
		}
		if input == 0 {
			break
		}
	}
}

func formHex(hex string, length int) string {
	var hexArr []byte = []byte(hex)
	for i := 0; i < length-len(hex); i++ {
		hexArr = append(hexArr, '0')
	}
	var hexString string = "0x" + string(hexArr)
	return hexString
}

func enterHex() (string, int) {
	var enterHex string
	for {
		fmt.Print("Enter your hex value: ")
		fmt.Scan(&enterHex)
		matched, err := regexp.MatchString("^[a-f0-9]+$", enterHex)
		if matched {
			break
		} else {
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Please, enter hex value ([0-9] and [a-f])!")
		}
	}
	fmt.Print("Enter number of bytes: ")
	var enterLength int
	fmt.Scan(&enterLength)
	return enterHex, enterLength
}

func enterEndian() (string, int) {
	var enterHex string
	for {
		fmt.Print("Enter your value: ")
		fmt.Scan(&enterHex)
		matched, err := regexp.MatchString("^[0-9]+$", enterHex)
		if matched {
			break
		} else {
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Please, enter integer value!")
		}
	}
	fmt.Print("Enter number of bytes: ")
	var enterLength int
	fmt.Scan(&enterLength)
	return enterHex, enterLength
}

func hexToLittleEndian() {
	var enterHex, enterLength = enterHex()
	var hexString string = formHex(enterHex, enterLength)
	fmt.Println("Your hex: ", hexString)
	var hexStringArr []byte = []byte(hexString)
	var hexArrZero []byte = hexStringArr[2:]
	var hexArr []byte
	for i := 0; i < len(hexArrZero); i++ {
		if hexArrZero[i] == 0 {
			break
		} else {
			hexArr = append(hexArr, hexArrZero[i])
		}
	}
	var hexArrRev []byte
	for i := len(hexArr) - 1; i >= 0; i-- {
		hexArrRev = append(hexArrRev, hexArr[i])
	}
	var hex string = string(hexArrRev)
	value, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Println("Something went wrong!")
	} else {
		fmt.Println("Your little-endian value: ", value)
	}
}

func hexToBigEndian() {
	var enterHex, enterLength = enterHex()
	var hexString string = formHex(enterHex, enterLength)
	fmt.Println("Your hex: ", hexString)
	var value, ok = new(big.Int).SetString(hexString, 10)
	if !ok {
		fmt.Println("Something went wrong!")
	} else {
		fmt.Println("Your big-endian value: ", value)
	}
}

func littleEndianToHex() {
	var enterValue, enterLength = enterEndian()
	intValue, err := strconv.Atoi(enterValue)
	if err != nil {
		fmt.Println("Something went wrong!")
	} else {
		var hex string = strconv.FormatInt(int64(intValue), 16)
		var hexString string = formHex(hex, enterLength)
		fmt.Println("Your hex value: ", hexString)
	}

}

func bigEndianToHex() {
	var enterValue, enterLength = enterEndian()
	var value, ok = new(big.Int).SetString(enterValue, 10)
	if !ok {
		fmt.Println("Something went wrong!")
	} else {
		var hex string = formHex(value.Text(16), enterLength)
		fmt.Println("Your hex value: ", hex)
	}
}
