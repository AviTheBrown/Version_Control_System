package datatypes

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

const userDataFile = "userdata.gob"

func SaveUser(user *User) {
	file, err := os.Create(userDataFile)
	if err != nil {
		log.Fatal("Error creating file", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(user)
	if err != nil {
		fmt.Println("Error encoding user:", err)
	}
}

func LoadUser() *User {
	file, err := os.Open(userDataFile)
	if err != nil {
		fmt.Println("No exisiting user data found", err)
		return CreateUser()
	}
	defer file.Close()
	var user User
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error decoding user data:", err)
		return CreateUser()
	}
	return &user
}
