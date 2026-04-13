package main

// 13 Apr 2026
// CH - 6 : Errors

// ========================================== L1 : The Error Interface ==========================================

import (
	"errors"
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost1, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	cost2, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return cost1+cost2, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// ========================================== L2 : Formatting Strings Review ==========================================


func getSMSErrorString(cost float64, recipient string) string {
	s := fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
	return s
}


// ========================================== L3 : The Error Interface ==========================================

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

func divide1(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}


// ========================================== L4 : Errors Quiz ==========================================

// Interface

// ========================================== L5 : Errors Quiz ==========================================

// Yes

// ========================================== L6 : The Errors Package ==========================================

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}


// ========================================== L7 : Panic ==========================================

// basically never


// ========================================== L8 : Panic ==========================================

// log.Fatal()

// ========================================== L9 : User Input ==========================================

type statusError struct {
	statusEmpty string
	statusExceed string
}

func (se statusError) Error() string {return "status cannot be empty"}

type statusEmptyError struct {
	status string 
}

func (seme statusEmptyError) Error() string {return "status cannot be empty"}

type statusExceedError struct {
	status string
	chara int
}

func (sexe statusExceedError) Error() string {return fmt.Sprintf("status exceeds %v characters", sexe.chara)}

func validateStatus(status string) error {
	if status == "" {return statusEmptyError{status:  status}}
	if len(status) > 140 {return statusExceedError{status: status, chara: 140}}
	return nil
}


func validateStatus1(status string) error {
	if status == "" {return errors.New("status cannot be empty")}
	if len(status) > 140 {return errors.New("status exceeds 140 characters")}
	return nil
}
