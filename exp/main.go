package main

import (
	"fmt"

	"github.com/meilleur/models"

	_ "github.com/lib/pq"
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
	//db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	//us.DestructiveReset()

	user := models.User{
		Name:  "Sete Gibernau",
		Email: "sete@worldf1.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	// user, err := us.ByID(2)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(user)

}
