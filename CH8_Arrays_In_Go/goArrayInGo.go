// 14 April 2026
// CH - 8 : Array in Go

package main


import (
	"errors"
	"strings"
)

// ========================================== L1 : Arrays in Go ==========================================

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{primary, secondary, tertiary}, [3]int{len(primary), len(primary)+len(secondary), len(primary)+len(secondary)+len(tertiary)}
}

// ========================================== L2 : Slices in Go ==========================================

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro { return messages[:],nil}
	if plan == planFree { return  messages[0:2], nil}
	return nil, errors.New("unsupported plan")
}

// ========================================== L3 : Slices Review ==========================================

// Slices reference arrays

// ========================================== L4 : Slices Review ==========================================

// True

// ========================================== L5 : Slices Review ==========================================

// True

// ========================================== L6 : Make ==========================================


func getMessageCosts(messages []string) []float64 {
	lenght := len(messages)
	cost := make([]float64, lenght)
	for i:=0;i<lenght;i++{
		cost[i] = float64(len(messages[i])) * 0.01
	}
	return cost
}

// ========================================== L7 : Len and Cap Review ==========================================

// The maximum length of the slice before reallocation of the array is necessary

// ========================================== L8 : Len and Cap Review ==========================================

// The current length of the slice

// ========================================== L9 : Len and Cap Review ==========================================

// Return 0

// ========================================== L10 : Variadic ==========================================

func sum(nums ...int) int {
	sum := 0
	for i:=0; i<len(nums); i++ { sum += nums[i]}
	return sum
}

// ========================================== L11 : Append ==========================================

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	costDay := []float64{}
	for i:=0;i<len(costs);i++ {
		if costs[i].day == day { costDay = append(costDay, costs[i].value) }
	}
	return costDay
}

// ========================================== L12 : Range ==========================================

func indexOfFirstBadWord(msg []string, badWords []string) int {
	
	for i, word := range msg {
		for _, bword := range badWords {
			if word == bword {return i}
		}
	}
	return -1
}

// ========================================== L13 : Slice of Slices ==========================================

func createMatrix0(rows, cols int) [][]int {
	result := [][]int{}
	for i:=0;i<rows;i++{
		temp := []int{}
		for j:=0;j<cols;j++{
			temp = append(temp, i*j)
		}
		result = append(result, temp)
	}
	return  result
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = i * j
		}
	}

	return matrix
}

// ========================================== L14 : Tricky Slices ==========================================

// j and g point to the same underlying array so g's append overwrote j

// ========================================== L15 : Tricky Slices ==========================================

// The array's cap() is exceeded so a new underlying array is allocated

// ========================================== L16 : Tricky Slices ==========================================

// Always assign the result of the append() function back to the same slice

// ========================================== L17 : Message Filter ==========================================

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	result := make([]Message, 0)
	for _, message := range messages {
		if message.Type() == filterType {
			result = append(result, message)
		}
	}
	return  result
}

// ========================================== L18 : Password Strength ==========================================

func isValidPassword(password string) bool {
	capitalOk := false
	numberOk := false
	if len(password) < 5 || len(password) > 12 {return false}
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {capitalOk = true}
		if char >= '0' && char <= '9' {numberOk = true}
		if capitalOk && numberOk { return true}
	}
	return false
}

// ========================================== L19 : Message Tagger ==========================================

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages0(messages []sms, tagger func(sms) []string) []sms {
	result := make([]sms,0,len(messages))
	for _, message := range messages {
		newMessage := message
		newMessage.tags = tagger(newMessage)
		result = append(result, newMessage)
	}
	return  result
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i:=0; i<len(messages); i++{
		messages[i].tags = tagger(messages[i])
	}
	return  messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {tags = append(tags, "Urgent")}
	if strings.Contains(strings.ToLower(msg.content), "sale") {tags = append(tags, "Promo")}
	return tags
}
