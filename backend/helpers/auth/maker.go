package auth

/**
 * An interface for managing tokens (for flexibiltiy of switching between different types of token)
 */
type Maker interface {
	// CreateToken creates a new token
	CreateToken(userID string, duration int64) (string, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(accessToken string) (*Claim, error)
}
