// 29 March 2026
// CH - 5 : Interfaces

package main

// ========================================== L1 : Interfaces in Go ==========================================

import (
	"fmt"
	"time"
)

// func sendMessage(msg message) (string, int) {
// 	return msg.getMessage() , len(msg.getMessage())*3
// }

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// ========================================== L2 : Interface Implementation ==========================================

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return  c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}


// ========================================== L3 : Interfaces Are Implemented Implicitly ==========================================

// A type has all the required interface's methods defined on it

// ========================================== L4 : Interfaces Are Implemented Implicitly ==========================================

// yes, why not

// ========================================== L5 : Interfaces Quiz ==========================================

// there is no keyword in Go

// ========================================== L6 : Interfaces Quiz ==========================================

// circle, shape

// ========================================== L7 : Multiple Interfaces ==========================================

func (e email) cost0() int {
	if !e.isSubscribed {
		return len(e.body)*5
	} else { return len(e.body)*2}

}

func (e email) format() string {
	if e.isSubscribed {
		return "'" + e.body + "'" + " | " + "Subscribed"
	} else { 
		return  "'" + e.body + "'" + " | " + "Not Subscribed"
	}
}

type expense0 interface {
	cost0() int
}

// type formatter interface {
// 	format() string
// }

type email0 struct {
	isSubscribed bool
	body         string
}


// ========================================== L8 : Name Your Interface Parameters ==========================================

// No

// ========================================== L9 : Name Your Interface Parameters ==========================================

// Readability and clarity

// ========================================== L10 : Type Assertions in Go ==========================================


// func getExpenseReport(e expense) (string, float64) {
// 	m, ok := e.(email)
// 	if ok {
// 		toAddress := m.toAddress
// 		return toAddress, m.cost()
// 	}

// 	s, ok := e.(sms)
// 	if ok {
// 		return  s.toPhoneNumber, s.cost()
// 	}

// 	return  "", 0.0
// }

// // don't touch below this line

// type expense interface {
// 	cost() float64
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }

// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }

// type invalid struct{}

// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .03
// }

// func (i invalid) cost() float64 {
// 	return 0.0
// }

// ========================================== L11 : Type Switches ==========================================

func getExpenseReport(e expense) (string, float64) {
	switch t := e.(type){
	case email:
		return t.toAddress, t.cost()
	case sms :
		return t.toPhoneNumber, t.cost()
	default :
		return  "", 0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

// ========================================== L12 : Clean Interfaces ==========================================

// Few

// ========================================== L13 : Clean Interfaces ==========================================

// False

// ========================================== L14 : Clean Interfaces ==========================================

// To only define the essential behaviors

// ========================================== L15 : Clean Interfaces ==========================================

// True wrong (False)

// ========================================== L16 : Message Formatter ==========================================


type formatter interface {
	format() string
}

type plainText struct{
	message string
}

type bold struct{
	message string
}

type code struct{
	message string
}

func (pt plainText) format() string {
	return  pt.message
}

func (b bold) format() string {
	return  "**" + b.message + "**"
}

func (c code) format() string {
	return  "`" + c.message + "`"
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}


// ========================================== L17 : Process Notification ==========================================

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return  50
	} else {
		return dm.priorityLevel
	}
}

func (gm groupMessage) importance() int {
	return  gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch pn := n.(type) {
	case directMessage:
		return pn.senderUsername, pn.importance()
	case groupMessage :
		return pn.groupName, pn.importance()
	case systemAlert :
		return  pn.alertCode, pn.importance()
	default :
		return  "", 0
	}
}
