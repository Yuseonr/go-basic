// 21 March 2026
// CH - 3 : Functions

package main

// ========================================== L1 : Functions ==========================================

import (
		"fmt"
		"errors")

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

// func main1() {
// 	test("Lane,", " happy birthday!")
// 	test("Zuck,", " hope that Metaverse thing works out")
// 	test("Go", " is fantastic")
// }

// func test(s1 string, s2 string) {
// 	fmt.Println(concat(s1, s2))
// }


// ========================================== L2 : Multiple Parameters ==========================================

// func createUser(firstName, lastName string, age int)

// ========================================== L3 : Unit Test Lessons ==========================================

func getMonthlyPrice(tier string) int {
	switch tier{
	case "basic" :
		return 10000 
	case "premium" :
		return 15000
	case "enterprise":
		return  50000
	default :
		return  0
	}
}


// ========================================== L4 : Declaration Syntax ==========================================

// The style of language used to create new variables, types, functions, etc...

// ========================================== L5 : Declaration Syntax ==========================================

// Go

// ========================================== L6 : Declaration Syntax ==========================================

// A function named 'f' that takes a function and an int as arguments and returns an int

// ========================================== L7 : Passing Variables by Value ==========================================

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	thisMonthBill = getBillForMonth(costPerSend, numLastMonth)
	lastMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return lastMonthBill - thisMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}


// ========================================== L8 : Ignoring Return Values ==========================================


func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

// don't touch below this line

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}

// ========================================== L9 : Named Return Values ==========================================

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	
	// don't touch below this line

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}


// ========================================== L10 : The Benefits of Named Returns ==========================================

// For small functions

// ========================================== L11 : The Benefits of Named Returns ==========================================

// When there are many values being returned

// ========================================== L12 : Explicit Returns ==========================================

// func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking = 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilCarRental = 25 - age
// 	if yearsUntilCarRental < 0 {
// 		yearsUntilCarRental = 0
// 	}
// 	return 
// }


// ========================================== L13 : Early Returns ==========================================

// Guard clauses provide a linear approach to logic trees

// ========================================== L14 : Early Returns ==========================================

// An early return from a function when a given condition is met

// ========================================== L15 : Functions As Values ==========================================

func reformat(message string, formatter func(string) string) string {
	stuff := formatter(formatter(formatter(message)))
	return "TEXTIO: " + stuff
}

// ========================================== L16 : Anonymous Functions ==========================================

func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return len(s)*2
	}, intro)
	printCostReport(func(s string) int {
		return len(s)*3
	}, body)
	printCostReport(func(s string) int {
		return len(s)*4
	}, outro)
}

// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

// ========================================== L17 : Defer ==========================================


func bootup() {

	defer fmt.Println("TEXTIO BOOTUP DONE")
	
	ok := connectToDB()
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

// don't touch below this line

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func main2() {
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}

// ========================================== L18 : Block Scope ==========================================

func splitEmail(email string) (string, string) {
	
	username, domain := "", ""
	
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

// ========================================== L19 : Processing Orders ==========================================


func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	if !(amountInStock(productID) >= quantity) {
		return false, accountBalance
	}
	if !(calcPrice(productID, quantity) <= accountBalance){
		return false, accountBalance
	}

	return true, accountBalance - calcPrice(productID,quantity)
	
	
}

// Don't touch below this line

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}


// ========================================== L20 : Closures ==========================================


func adder() func(int) int {
	sum := 0
	return func(a int)int{
		sum = sum + a
		return sum		
	}
}

// ========================================== L21 : Currying ==========================================

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(first string, second string) {
		result := formatter(first, second)
		fmt.Println(result)
	}
}

// don't touch below this line

func test3(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main5() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test3("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test3("Error on mail server", mailErrors, commaDelimit)
}
