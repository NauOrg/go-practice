package models

import (
	"errors"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// Mock database
var users = []User{
	{ID: 1, Email: "user@example.com", Password: "$2a$14$9.3P.yjrKNb8/TUItwl6ueOW/5H20.LecDPZVbB6fIo5fP7f1eLZi"}, // bcrypt hash for "password"
}

var nextID = 2

func FindUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

var mu sync.Mutex

func CreateUser(user *User) error {
	mu.Lock()
	defer mu.Unlock()
	user.ID = nextID
	nextID++
	users = append(users, *user)
	return nil
}
