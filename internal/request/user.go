package request

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type user struct {
	ID            int64     `bson:"_id"`
	Name          string    `bson:"name"`
	Password      string    `bson:"password"`
	Created_at    time.Time `bson:"created_at"`
}

var StudentIDRegex = regexp.MustCompile("[8-9][0-9][0-9]{2}[0-9]{3}")

func (req user) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.ID, validation.Required, validation.Length(7, 7), is.Int, validation.Match(StudentIDRegex)),
	)
}
