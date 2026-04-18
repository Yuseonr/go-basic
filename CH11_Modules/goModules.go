// 17 April 2026
// CH - 11 : Variables

package main

// ========================================== L1 : Packages ==========================================

// package main

import (
	"fmt"
)

func test(text string) {
	fmt.Println(text)
}

func main() {
	test("starting Textio server")
	test("stopping Textio server")
}

// ========================================== L2 : Package Naming ==========================================

// parser

// ========================================== L3 : Package Naming ==========================================

// any of these

// ========================================== L4 : Install ==========================================

// done install

// ========================================== L5 : Modules ==========================================

// A collection of packages that are released together

// ========================================== L6 : Modules ==========================================

// No

// ========================================== L7 : Modules ==========================================

// A module path + package subdirectory

// ========================================== L8 : Go Environment ==========================================

// No, in fact you shouldn't

// ========================================== L9 : First Local Program ==========================================

// go mod init github.com/Yuseonr/hellogo
// cat go.mod

// ========================================== L10 : Go Run ==========================================

// go run main.go

// ========================================== L11 : Go Build ==========================================

// go build
// ./goLocal

// ========================================== L12 : Go Install ==========================================

// go install
// Installed globaly in machine

// ========================================== L13 : Custom Package ==========================================

// goLocal-Write

// ========================================== L14 : Custom Package Continued ==========================================

// module example.com/username/hellogo

// go 1.26.0

// replace example.com/username/mystrings v0.0.0 => ../mystrings

// require example.com/username/mystrings v0.0.0

// ========================================== L15 : Remote Packages ==========================================

// go get github.com/wagslane/go-tinytime

// ========================================== L16 : Clean Packages ==========================================

// Nope

// ========================================== L17 : Clean Packages ==========================================

// When the end-user doesn't need to know about it

// ========================================== L18 : Clean Packages ==========================================

// No, try to keep changes to internal functionality