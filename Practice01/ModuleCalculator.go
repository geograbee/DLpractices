package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	for {
		var end bool = false
		fmt.Print("Enter the module (m): ")
		var m int
		fmt.Scan(&m)
		for {
			fmt.Println("Choose operation:")
			fmt.Println("1. a mod m = x")
			fmt.Println("2. a^b mod m = x")
			fmt.Println("3. a*x ≡ b mod m")
			fmt.Println("4. Get simple number in range A - B")
			fmt.Println("5. Enter new m")
			fmt.Println("0. Exit")
			fmt.Print("Enter the number of chosen operation: ")
			var input int
			fmt.Scan(&input)
			if input == 1 {
				func1(m)
			}
			if input == 2 {
				func2(m)
			}
			if input == 3 {
				func3(m)
			}
			if input == 4 {
				func4()
			}
			if input == 5 {
				break
			}
			if input == 0 {
				end = true
				break
			}
		}
		if end == true {
			break
		}
	}
}

func func1(m int) {
	fmt.Print("Enter your number: ")
	var a int
	fmt.Scan(&a)
	var x int = a % m
	fmt.Println("Remainder: ", x)
}

func func2(m int) {
	for {
		fmt.Print("Enter your number and power of number: ")
		var a, b int
		fmt.Scan(&a, &b)
		if a <= 0 || b <= 0 {
			fmt.Println("The numbers must be positive!")
			continue
		} else {
			var bBin []int
			var o int = b
			for o != 0 {
				bBin = append(bBin, o%2)
				o = o / 2
			}
			var d int = 1
			var t int = a
			for i := 0; i < len(bBin); i++ {
				if bBin[i] == 1 {
					d = (d * t) % m
					t = (t * t) % m
				} else {
					t = (t * t) % m
				}
			}
			fmt.Println("Remainder: ", d)
			break
		}
	}

}

func func3(m int) {
	for {
		fmt.Print("Enter your two numbers: ")
		var a, b int
		fmt.Scan(&a, &b)
		if a <= 0 || b <= 0 {
			fmt.Println("The numbers must be positive!")
			continue
		} else {
			if GCD(a, m) == 1 {
				var af float64 = float64(a)
				var phi float64 = float64(phi(m))
				var phi1 float64 = phi - 1
				aphi := int(math.Pow(af, phi1))
				var xaphi int = aphi % m
				var x int = (b * xaphi) % m
				fmt.Println("X: ", x)
				break
			}
		}
	}
}
func func4() {
	for {
		fmt.Print("Enter your two numbers: ")
		var a, b int
		fmt.Scan(&a, &b)
		if a <= 0 || b <= 0 {
			fmt.Println("The numbers must be positive!")
			continue
		} else {
			var primes []int
			for i := a; i <= b; i++ {
				if i == 1 {
					continue
				}
				if i == 2 || i == 3 || (i*i)%24 == 1 {
					primes = append(primes, i)
				}
			}
			var rand int = primes[rand.Intn(len(primes))]
			fmt.Println("Your random prime number is: ", rand)
			break
		}
	}
}

func phi(n int) int {
	var res int = 1
	for i := 2; i < n; i++ {
		if GCD(i, n) == 1 {
			res++
		}
	}
	return res
}

func GCD(a int, b int) int {
	if a == 0 {
		return b
	}
	i := 1
	max := 1000
	for i < max {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
		if b == 0 {
			break
		}
		i++
	}
	return a
}
