package main

import (
	"fmt"
)

// 13 Apr 2026
// CH - 7 : Loop in Go

// ========================================== L1 : Loops in Go ==========================================

func bulkSend(numMessages int) float64 {
	var total float64 = 0.0
	for i:= 0; i<numMessages; i++ {
		total += 1.00 + (0.01 * float64(i))
	}
	return  total
}

// ========================================== L2 : Omitting Conditions from a for Loop in Go ==========================================

func maxMessages(thresh int) int {
	total := 0
	for i:=0; ;i++ {
		total += 100 + i
		if total > thresh { return i }
	}
}

// ========================================== L3 : There Is No While Loop in Go ==========================================

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance >= 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

// ========================================== L4 : Fizzbuzz ==========================================

func fizzbuzz() {
	for i:=1; i<=100; i++{

		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			println(i)
		}
	}
}

// don't touch below this line

// func main() {
// 	fizzbuzz()
// }


// ========================================== L5 : Continue & Break ==========================================

// package main

// import (
// 	"fmt"
// )

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		isPrime := true
		for n := 2; n*n <= i; n++ {
			if i%n == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}


// ========================================== L6 : Connections ==========================================

func countConnections(groupSize int) int {
	total := groupSize
	count := 0
	for total > 0 {
		count += total - 1
		total -=1
	}
	return  count
}
