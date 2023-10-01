package request

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID        string    `bson:"id"`
	Name      string    `bson:"name"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}

var StudentIDRegex = regexp.MustCompile("[8-9][0-9][0-9]{2}[0-9]{3}")

func (req User) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.ID, validation.Required, validation.Length(8, 8), is.Int, validation.Match(StudentIDRegex)),
	)
}
