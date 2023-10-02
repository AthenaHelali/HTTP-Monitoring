package uservalidator

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) error {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required,
			validation.Match(regexp.MustCompile(IDRegex)),
			validation.By(v.doesIdExist)),
		validation.Field(&req.Password, validation.Required),
	); err != nil {
		return errors.New("invalid input")
	}
	return nil
}

func (v Validator) doesIdExist(value interface{}) error {
	id := value.(string)

	_, err := v.repo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
