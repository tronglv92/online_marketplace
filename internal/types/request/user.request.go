package request

import (
	"context"

	validation "github.com/itgelo/ozzo-validation"
	"github.com/itgelo/ozzo-validation/is"
)

type RegisterReq struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (req RegisterReq) Validate(ctx context.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 10).Error("Password require length between 6 and 10")),
	)
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginReq) Validate(ctx context.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 10).Error("Password require length between 6 and 10")),
	)
}

type ProfileReq struct {
	Uid string `json:"uid"`
}

func (req ProfileReq) Validate(ctx context.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Uid, validation.Required, is.UUID),
	)
}
