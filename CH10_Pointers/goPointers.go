package main

// 16 apr 2026
// CH - 10 : Pointers

// ========================================== L1 : Introduction to Pointers ==========================================

import (
	"fmt"
	"strings"
	"errors"
)

type Message1 struct {
	Recipient string
	Text      string
}

func getMessageText(m Message1) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}


// ========================================== L2 : References ==========================================

// func removeProfanity(message *string) {
// 	val := *message

// 	val = strings.ReplaceAll(val, "fubb", "****")
// 	val = strings.ReplaceAll(val, "shiz", "****")
// 	val = strings.ReplaceAll(val, "witch", "*****")

// 	*message = val
	
// }

// ========================================== L3 : Pass by Reference ==========================================

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage (al *Analytics, msg Message){
	if msg.Success {
		(*al).MessagesSucceeded++
	} else {
		(*al).MessagesFailed++
	}
	(*al).MessagesTotal++
}

// ========================================== L4 : Pointers Quiz ==========================================

// 100

// ========================================== L5 : Pointers Quiz ==========================================

// 100

// ========================================== L6 : Nil Pointers ==========================================

func removeProfanity(message *string) {
	if message == nil { return }
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

// ========================================== L7 : Pointer Receivers ==========================================

// Pointer receivers

// ========================================== L8 : Pointer Receiver Code ==========================================

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

// ========================================== L9 : Pointer Performance ==========================================

// Heap

// ========================================== L10 : Pointer Performance ==========================================

// Value

// ========================================== L11 : Update Balance ==========================================

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance (cust *customer, txn transaction) error{
	if txn.transactionType == transactionDeposit {
		(*cust).balance += txn.amount
		return  nil
	} else if txn.transactionType == transactionWithdrawal {
		temp := (*cust).balance
		if temp - txn.amount < 0 { return errors.New("insufficient funds")}
		(*cust).balance -= txn.amount
		return  nil
	} else {
		return  errors.New("unknown transaction type")
	}
}

func updateBalance0(cust *customer, txn transaction) error {
    switch txn.transactionType {
    case transactionDeposit:
        cust.balance += txn.amount
        return nil

    case transactionWithdrawal:
        if cust.balance < txn.amount {
            return errors.New("insufficient funds")
        }
        cust.balance -= txn.amount
        return nil

    default:
        return errors.New("unknown transaction type")
    }
}
