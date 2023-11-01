package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
	Password  string    `json:"-" db:"password"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *User) ValidPassword(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)) == nil
}

func (u *User) IsProfileRecentlyUpdated(iat float64) (bool, error) {
	return u.UpdatedAt.Before(time.Unix(int64(iat), 0)), nil
}

func NewUser(name, email, password string) (*User, error) {

	encPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(encPwd),
	}, nil
}
