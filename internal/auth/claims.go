package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID           string
	RegisteredClaims jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	// TODO
	return nil
}

func (c Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.GetExpirationTime()
}

func (c Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.GetIssuedAt()
}

func (c Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.GetNotBefore()
}

func (c Claims) GetIssuer() (string, error) {
	return c.RegisteredClaims.GetIssuer()
}

func (c Claims) GetSubject() (string, error) {
	return c.RegisteredClaims.GetSubject()
}

func (c Claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.RegisteredClaims.GetAudience()
}
