package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (appUser *User) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func (appUser *User) ClearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
