package main

import (
	"fmt"

	"github.com/meilleur/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port= %d  user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	// us.DestructiveReset()
	us.AutoMigrate()

	user := models.User{
		Name:     "Jon Calhoun",
		Email:    "jon@calhoun.io",
		Password: "jon",
		Remember: "abc123",
	}
	err = us.Create(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)

	// toHash := []byte("this is my string to hash")
	// h := hmac.New(sha256.New, []byte("my-secret-key"))
	// h.Write(toHash)
	// b := h.Sum(nil)
	// fmt.Println(base64.URLEncoding.EncodeToString(b))

	// hmac := hash.NewHMAC("my-secret-key")
	// fmt.Println(hmac.Hash("this is my string to hash"))

}
