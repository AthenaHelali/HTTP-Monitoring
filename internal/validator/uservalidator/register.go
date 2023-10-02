package uservalidator

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) error {

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),

		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]{10,}$"))),

		validation.Field(&req.ID, validation.Required,
			validation.Match(regexp.MustCompile(IDRegex)),
			validation.By(v.checkIDUniqueness)),
	); err != nil {
		return errors.New("invalid input")
	}
	return nil
}

func (v Validator) checkIDUniqueness(value interface{}) error {
	id := value.(string)

	uniqueness, err := v.repo.IsIdUnique(id)
	if err != nil {
		return errors.New("StatusNotFound")
	}
	if !uniqueness {
		return errors.New("UserID is not unique")
	}
	return nil
}
