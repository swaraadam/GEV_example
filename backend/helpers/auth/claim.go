package auth

import (
	"errors"
	"gev_example/helpers"
	"time"

	"github.com/google/uuid"
)

type Claim struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"userer_id"`
	Issuer    string    `json:"issuer"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewClaim(gamerID string, duration time.Duration) (*Claim, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	claim := &Claim{
		ID:       tokenID,
		UserId:   gamerID,
		Issuer:   helpers.Env.TokenIssuer,
		IssuedAt: time.Now(),
	}
	if duration != 0 {
		// if duration is specified, use specified duration
		claim.ExpiredAt = time.Now().Add(duration * time.Second)
	} else {
		// else use default duration of 1 month
		claim.ExpiredAt = time.Now().Add(helpers.Env.TokenDuration)
	}
	return claim, nil
}

/**
 * Valid checks if the token payload is valid or not
 */
func (claim *Claim) Valid() error {
	if time.Now().After(claim.ExpiredAt) {
		return errors.New("Token Expired!")
	}
	return nil
}
