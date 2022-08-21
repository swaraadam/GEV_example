package auth

import (
	"errors"
	"fmt"
	"gev_example/helpers"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JWTMaker struct{}

// CreateToken implements Maker
func (*JWTMaker) CreateToken(userID string, duration int64) (string, error) {
	claim, err := NewClaim(userID, time.Duration(duration))
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(helpers.Env.TokenSecret))
}

// VerifyToken implements Maker
func (*JWTMaker) VerifyToken(accessToken string) (*Claim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("AccessToken Invalid!")
		}
		return []byte(helpers.Env.TokenSecret), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &Claim{}, keyFunc)
	if err != nil {
		validator, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validator.Inner, errors.New("AccessToken Expired!")) {
			return nil, errors.New("AccessToken Expired!")

		}
		return nil, errors.New("AccessToken Invalid!")
	}
	claim, ok := token.Claims.(*Claim)
	if !ok {
		return nil, errors.New("AccessToken Invalid!")
	}
	return claim, nil
}

func NewJWTMaker() (Maker, error) {
	if len(helpers.Env.TokenSecret) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size, must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{}, nil
}
