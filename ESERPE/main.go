package main

import (
	"errors"
	"fmt"
)

// Lemari struct merepensetasikan sebuah Lemari yang sergin kita temukan di rumah
type Lemari struct {
	Nama      string
	IsiLemari interface{}
	Rak       int
}

func LemariBaru() *Lemari {
	return &Lemari{}
}

func (l *Lemari) MasukkanBaju(baju interface{}, nama string, rak int) {
	l.Nama = nama
	l.IsiLemari = baju
	l.Rak = rak
}

// Contoh Menggunakan Single Responsibility Principle pada authentication User
type User struct {
	Username string
	Password string
}

type UserRepository struct {
	User []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		User: make([]User, 0),
	}
}

func (users *UserRepository) NewUser(user *User) {
	users.User = append(users.User, *user)
}

type AuthenticationUser struct {
	userRepository *UserRepository
}

func NewAuthenticationUser(userRepository *UserRepository) *AuthenticationUser {
	return &AuthenticationUser{
		userRepository: userRepository,
	}
}

func (auth *AuthenticationUser) Authenticate(username string, password string) (bool, error) {
	users := auth.userRepository.User

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true, nil
		}
	}

	return false, errors.New("error not match of any user")
}

func main() {

	adolo := User{
		Username: "adolo",
		Password: "self_auth",
	}

	UserRepository := NewUserRepository()

	UserRepository.NewUser(&adolo)

	UserAuth := NewAuthenticationUser(UserRepository)

	Isauthenticate, err := UserAuth.Authenticate("adolo", "self_auth")
	if err != nil {
		fmt.Println(err)
	}

	if Isauthenticate {
		fmt.Println("User authen")
	}
	// lemariAnton := LemariBaru()
	// //Lemari anton hanya boleh memasukkan baju ke dalam sebuah lemari
	// lemariAnton.MasukkanBaju("louis vuitton", "lemari anton", 1)

	// fmt.Println(lemariAnton.IsiLemari)
}
