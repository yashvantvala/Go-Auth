package models

import (
	"context"
	"errors"

	"github.com/yashvantvala/Go-Auth/db"
	"github.com/yashvantvala/Go-Auth/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	FirstName string `json:"firstname", bson:"firstname"`
	LastName  string `json:"lastname", bson:"lastname"`
	Email     string `json:"email", bson:"email" binding: "required"`
	Password  string `json:"password", bson:"password" binding: "required"`
}

func (u *User) CreateUser() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	if err != nil {
		return err
	}
	_, err = db.Collections.InsertOne(context.Background(), u)
	if err != nil {
		return err
	}

	return err
}

func (u *User) FindUser() error {
	filter := bson.M{"email": u.Email}
	var user User
	err := db.Collections.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return err
	}
	isValidPassword := utils.ComparePasswordHash(u.Password, user.Password)
	if !isValidPassword {
		return errors.New("Invalid credentials!")
	}
	return nil
}
