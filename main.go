package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
)

func generatePassword(desiredLength int64) string {
	min := 33
	max := 126

	var currentLength int64 = 1
	var newPassword string

	for {
		rand.Seed(time.Now().UTC().UnixNano())
		myRand := rand.Intn(max - min) + min
		newPassword += string(myRand)
        if currentLength == desiredLength {
                break
        }
		currentLength++
}
	return newPassword
}

// TODO: there is probably a better way to build this string for the http writer
func getPasswd(w http.ResponseWriter, r *http.Request) {
	passwords := 
		"8  char: " + generatePassword(8) + "\n" + 
		"9  char: " + generatePassword(9) + "\n" + 
		"10 char: " + generatePassword(10) + "\n" + 
		"11 char: " + generatePassword(11) + "\n" + 
		"12 char: " + generatePassword(12) + "\n" + 
		"13 char: " + generatePassword(13) + "\n" + 
		"14 char: " + generatePassword(14) + "\n" + 
		"15 char: " + generatePassword(15) + "\n" + 
		"16 char: " + generatePassword(16) + "\n" + 
		"17 char: " + generatePassword(17) + "\n" + 
		"18 char: " + generatePassword(18) + "\n" + 
		"19 char: " + generatePassword(19) + "\n" + 
		"20 char: " + generatePassword(20) + "\n" + 
		"21 char: " + generatePassword(21) + "\n" + 
		"22 char: " + generatePassword(22) + "\n" + 
		"23 char: " + generatePassword(23) + "\n" + 
		"24 char: " + generatePassword(24) + "\n" + 
		"25 char: " + generatePassword(25)
	w.Write([]byte(passwords))
}

func main () {
	fmt.Println("I am awake! Finally, github is isolating")
	fmt.Println("If you want me to generate passwords go to the below webpage")
	fmt.Println("http://<my_ip>:8000/freePasswords")

	router := mux.NewRouter()
	router.HandleFunc("/freePasswords", getPasswd).Methods("GET")
  
  // Switch to serverTLS dont forget cert generation
	http.ListenAndServe(":8000", router)
}
