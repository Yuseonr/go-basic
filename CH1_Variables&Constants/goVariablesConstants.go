// 20 March 2026
// CH - 1 : Variables

// ========================================== L1 : Learn Go for Developer ==========================================

package main

import "fmt"

func main1() {
	fmt.Println("Starting Textio server...")
}

// ========================================== L2 : Declaring variable and Basic Variables  ==========================================

func main2() {
	// initialize variables here

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

func main2_1() {
	var username string
	username = "eddie_cabot"

	var isAdmin bool
	isAdmin = true

	var permissions int
	permissions = 0x1F

	var costPerSMS float64
	costPerSMS = 0.05

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("permissions:", permissions)
	fmt.Println("costPerSMS:", costPerSMS)
}


// ========================================== L3 : Short Variable declaration  ==========================================

func main3() {

	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)
}


// ========================================== L4 : Multiple Variable declaration  ==========================================

func main4() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}

// ========================================== L5 : Comments  ==========================================

func main5() {
	/*
		We are increasing the maximum message length from 140 to 280 characters.
		Very reluctantly, I might add.
		Users actually want to write more than 140 characters?!? Madness.
	*/
	maxMessageLength := 140
	newMaxMessageLength := 280
	fmt.Println("Textio is increasing the maximum message length from", maxMessageLength, "to", newMaxMessageLength, "characters.")
}

// ========================================== L6 : The Compilation Process  ==========================================

func main6() {
	fmt.Println("The compiled textio server is starting")
}

// ========================================== L7 : Fast and Compiled ==========================================

// Faster - Faster

// ========================================= L8 : Type size  ==========================================

func main8() {
	accountAgeFloat := 2.6
	accountAgeInt := int32(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

// ========================================= L9 : Which type should I use?  ==========================================

// When performance and memory are the primary concerns

// ========================================= L10 : Which type should I use?  ==========================================

// Bits

// ========================================= L11 : Which type should I use?  ==========================================

func main11() {
	var username string = "presidentSkroob"
	var password string = "12345"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}

// ========================================== L12 : Compiled vs. Interpreted =========================================

// The compiled executable

// ========================================== L13 : Compiled vs. Interpreted =========================================

// Python

// ========================================== L14 : Same Line Declarations =========================================

func main14() {

	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)
}

// ========================================== L15 : Small Memory Footprint =========================================

// Java

// ========================================== L16 : Small Memory Footprint =========================================

// To cleanup unused memory

// ========================================== L17 : Constants =========================================

func main17() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}

// ========================================== L18 : Computed Constants =========================================

func main18() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}

// ========================================== L19 : Comparing Go's Speed =========================================

// No

// ========================================== L20 : Formatting Strings in Go =========================================

func main20() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// don't edit above this line

	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent \n", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
}

// ========================================== L21 : Runes and String Encoding =========================================

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main21() {
// 	const name = "🐻"
// 	fmt.Printf("constant 'name' byte length: %d\n", len(name))
// 	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
// 	fmt.Println("=====================================")
// 	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
// }


// ========================================== L22 : Fix Bugs =========================================

func main22() {
	var startup string = "Textio SMS service booting up..."
	var message string = "Sending test message"
	var confirmation string = "Message sent!"

	// don't touch below this line

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirmation)
}


// ========================================== L23 : Fix Types =========================================

func main23() {
	var senderName string = "Syl"
	recipient := "Kaladin"
	message := "The Words, Kaladin. You have to speak the Words!"

	fmt.Printf("%s to %s: %s\n", senderName, recipient, message)
}

// ========================================== L24 : Type Inference =========================================

func main24() {
	penniesPerText := 2.0

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}

// ========================================== L25 : Format Practice ==========================================

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	// Don't touch below this line

	fmt.Println(userLog)
}
