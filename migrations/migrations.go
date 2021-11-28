package migrations

import (
	"github.com/LuisFrank/backendshopping/helpers"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=shoppingcart password=postgres sslmode=disabled")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "LuisFrank", Email: "lflorentinomedrano@gmail.com"},
		{Username: "Juan", Email: "esteemailesdeprueba@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

	}

	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{})
	defer db.Close()

	createAccounts()
}
