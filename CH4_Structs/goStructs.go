// 28 March 2026
// CH - 4 : Structs

package main

// ========================================== L1 : Structs in Go ==========================================

type messageToSend1 struct {
	phoneNumber int
	message    string
}

// ========================================== L2 : Nested Structs in Go ==========================================

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user2 struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	return mToSend.sender.number != 0 && mToSend.sender.name != "" && mToSend.recipient.number != 0 && mToSend.recipient.name != ""
}

// ========================================== L3 : Anonymous Structs in Go ==========================================

// It is only being used once

// ========================================== L4 : Anonymous Structs in Go ==========================================

// Anonymous structs prevent you from re-using a struct definition you never intended to re-use

// ========================================== L5 : Embedded Structs ==========================================

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

// ========================================== L6 : Struct Methods in Go ==========================================

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + a.username + ":" + a.password
}

// ========================================== L7 : Memory Layout ==========================================

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

// ========================================== L8 : Empty Struct ==========================================

// 'struct{}' is the type (empty struct) and '{}' is the value (empty struct literal)

// ========================================== L9 : Empty Struct ==========================================

// struct{}, bool, uint16, int64

// ========================================== L10 : Update Users ==========================================

type User1 struct {
	Membership
	Name string
}

type Membership1 struct {
	Type string
	MessageCharLimit int
}

func newUser1(name string, membershipType string) User {
	var MCL int = 100;
	if membershipType == "premium"{
		MCL = 1000
	}
	return  User {
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: MCL,
		},
	}
}

// ========================================== L11 : Send Message ==========================================

func (u User) SendMessage(message string, messageLenght int) (string, bool) {
	if messageLenght <= u.MessageCharLimit {
		return  message, true
	}
	return "",false
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
