package database

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *DB) CreateUser(email string, password string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Email:    email,
		Password: string(hashedPassword),
	}
	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUser(id int) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := dbStructure.Users[id]
	if !ok {
		return User{}, ErrNotExist
	}

	return user, nil
}

func (db *DB) UserAuth(user User) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, u := range dbStructure.Users {
		if u.Email == user.Email {
			// Compare the hashed password with the provided password
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
			if err != nil {
				return User{}, bcrypt.ErrMismatchedHashAndPassword
			}
			return u, nil
		}
	}

	return User{}, ErrNotExist
}
