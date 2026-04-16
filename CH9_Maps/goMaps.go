// 16 April 2026
// CH - 9 : Maps

package main

// ========================================== L1 : Maps ==========================================

import ("errors"
"strings"
)
// func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
// 	if len(names) != len(phoneNumbers) {return nil, errors.New("invalid sizes")}
// 	result := make(map[string]user)
// 	for i:=0; i<len(names); i++ {
// 		result[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]} 
// 	}
// 	return  result, nil
// }

// type user struct {
// 	name        string
// 	phoneNumber int
// }

// ========================================== L2 : Mutations ==========================================

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	userRecord, ok := users[name]

	if !ok {return false, errors.New("not found")}

	if !userRecord.scheduledForDeletion {return false, nil}

	delete(users, name)

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

// ========================================== L3 : Key Types ==========================================

// The type is comparable

// ========================================== L4 : Key Types ==========================================

// To use a struct directly as a key

// ========================================== L5 : Count Instances ==========================================

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, messagedUser := range messagedUsers {
		if _, ok := validUsers[messagedUser]; ok {
			validUsers[messagedUser]++
		}
	}
}

// ========================================== L6 : Effective Go ==========================================

// 1

// ========================================== L7 : Effective Go ==========================================

// Returns the zero value

// ========================================== L8 : Effective Go ==========================================

// affect

// ========================================== L9 : Effective Go ==========================================

// A boolean that indicates whether the key exists

// ========================================== L10 : Nested ==========================================

func getNameCounts(names []string) map[rune]map[string]int {
	result := map[rune]map[string]int{}

	for _, name := range names{
		r := []rune(name)[0]
		if mappedRune, ok := result[r]; ok{
			mappedRune[name]++
		} else {
			temp := map[string]int{name:1}
			result[r] = temp
		}
	}
	return result
	
}

// ========================================== L11 : Distinct Words ==========================================

func countDistinctWords(messages []string) int {
	seenMessage := make(map[string]struct{},100)
	
	for _, message := range messages {
		messageSplit := strings.Fields(strings.ToLower(message))

		for _, ms := range messageSplit {
			seenMessage[ms] = struct{}{}
		}
	}
	return  len(seenMessage)
}
