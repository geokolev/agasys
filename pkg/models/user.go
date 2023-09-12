package models

import (
	"time"

	"github.com/Kamva/mgm"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jonsch318/royalafg/pkg/user"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Email            string `json:"email" bson:"email"`
	Hash             string `json:"-" bson:"hash"`
	FullName         string `json:"fullName" bson:"fullName"`
	Birthdate        int64  `json:"birthdate" bson:"bithdate"`
	Verified         byte   `json:"verified" bson:"verified"`
	Banned           byte   `json:"banned" bson:"banned"`
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&user.Email, is.Email),
		validation.Field(&user.Hash, validation.Required),
		validation.Field(&user.FullName, validation.Required, validation.Length(1, 100)),
		validation.Field(&user.Birthdate, validation.Required,
			validation.Max(time.Now().AddDate(-16, 0, 0).Unix()),
			validation.Min(time.Now().AddDate(-150, 0, 0).Unix()),
		),
		validation.Field(&user.Verified, validation.Required),
		validation.Field(&user.Banned, validation.Required),
	)
}

// NewUser creates a new user with the given details.
// IMPORTANT THIS DOES NOT SAVE OR HASH THE PASSWORD. This has to be done seperatly
func NewUser(username, email, fullName string, birthdate int64) *User {
	return &User{
		Username:  username,
		Email:     email,
		FullName:  fullName,
		Birthdate: birthdate,
		Banned:    user.Valid,
		Verified:  0,
	}
}

func (user *User) GetBirthdate() time.Time {
	return time.Unix(user.Birthdate, 0)
}
